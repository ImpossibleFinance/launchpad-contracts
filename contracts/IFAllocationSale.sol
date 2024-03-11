// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import '@openzeppelin/contracts/utils/math/Math.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';
import './interfaces/IIFRetrievableStakeWeight.sol';
import './IFSale.sol';


/**
  @title Sale contract with user's allocation determined by staked weight
 */
contract IFAllocationSale is IFSale {
    IIFRetrievableStakeWeight public allocationMaster;
    uint80 public allocSnapshotTimestamp;
    uint24 public trackId;

    // --- CONSTRUCTOR

    constructor(
        uint256 _salePrice,
        address _funder,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        IIFRetrievableStakeWeight _allocationMaster,
        uint24 _trackId,
        uint80 _allocSnapshotTimestamp,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    ) 
        IFSale(
            _funder,
            _salePrice,
            _paymentToken,
            _saleToken,
            _startTime,
            _endTime,
            _maxTotalPayment
        )
    {
        require(
            _allocSnapshotTimestamp > block.timestamp ||
                (_allocSnapshotTimestamp <= block.timestamp &&
                    _allocationMaster.getTotalStakeWeight(
                        _trackId,
                        _allocSnapshotTimestamp
                    ) >
                    0),
            'total weight is 0 on while using older timestamp'
        );

        trackId = _trackId;
        allocationMaster = _allocationMaster; // can be 0 (with allocation override)
        allocSnapshotTimestamp = _allocSnapshotTimestamp; // can be 0 (with allocation override)
    }

    // --- SALE FUNCTIONS

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof)
        override
        public
        onlyAfterSale
        nonReentrant
    {
        address user = _msgSender();
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(
            whitelistRootHash == 0 || checkWhitelist(user, merkleProof),
            'proof invalid'
        );

        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            uint256 value = getUserStakeValue(user);
            claimable[user] = value;
            totalPurchased[user] = value;
        }
        uint256 saleTokenOwed = getCurrentClaimableToken(user);

        // send token and update states
        _withdraw(saleTokenOwed);
        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');
    }

    // purchase function when there is no whitelist
    function purchase(uint256 paymentAmount) override onlyDuringSale public {
        // there must not be a whitelist set (sales that use whitelist must be used with whitelistedPurchase)
        require(whitelistRootHash == 0, 'use whitelistedPurchase');
        uint256 remaining = getMaxPayment(_msgSender());
        _purchase(paymentAmount, remaining);
    }

    // --- CALCULATE ALLOCATION FROM STAKE WEIGHT

    function getUserStakeValue(address user) public view returns (uint256) {
        uint256 userWeight = allocationMaster.getUserStakeWeight(
            trackId,
            user,
            allocSnapshotTimestamp
        );
        uint256 totalWeight = allocationMaster.getTotalStakeWeight(
            trackId,
            allocSnapshotTimestamp
        );
        // total weight must be greater than 0
        require(totalWeight > 0, 'total weight is 0');

        // calculate max amount of obtainable sale token by user
        return (saleAmount * userWeight) / (totalWeight);
    }

    // Function to get the total allocation of a user in allocation sale
    // Allocation is calculated via the override if set, and otherwise
    // allocation is calculated by the allocation master data.
    function getTotalPaymentAllocation(address user)
        public
        view
        returns (uint256)
    {
        // get user allocation as ratio (multiply by 10**18, aka E18, for precision)
        uint256 userWeight = allocationMaster.getUserStakeWeight(
            trackId,
            user,
            allocSnapshotTimestamp
        );
        uint256 totalWeight = allocationMaster.getTotalStakeWeight(
            trackId,
            allocSnapshotTimestamp
        );

        // total weight must be greater than 0
        require(totalWeight > 0, 'total weight is 0');

        // determine TOTAL allocation (in payment token)
        uint256 paymentTokenAllocation;

        // calculate allocation (times 10**18)
        uint256 allocationE18 = (userWeight * 10**18) / totalWeight;

        // calculate max amount of obtainable sale token
        uint256 saleTokenAllocationE18 = (saleAmount * allocationE18);

        // calculate equivalent value in payment token
        paymentTokenAllocation =
            (saleTokenAllocationE18 * salePrice) /
            SALE_PRICE_DECIMALS /
            10**18;

        return paymentTokenAllocation;
    }

    // Function to get the MAX REMAINING amount of allocation for a user (in terms of payment token)
    // it is whichever is smaller:
    //      1. user's payment allocation
    //      2. maxTotalPayment
    function getMaxPayment(address user) public view returns (uint256) {
        // get the maximum total payment for a user
        uint256 max = getTotalPaymentAllocation(user);
        if (maxTotalPayment < max) {
            max = maxTotalPayment;
        }

        // calculate and return remaining
        return max - paymentReceived[user];
    }

}
