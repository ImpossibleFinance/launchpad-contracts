// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import './IFSale.sol';


/**
  @title Sale contract with user's allocation stored in merkle root
  @notice Regular purchase from IFSale is disabled here
  @notice Apart from merkle proof, users or the frontend has to supply allocation amount for verification
  @notice That means functions having `merkleProof` as param will also need `allocation`
 */
contract IFFixedSale is IFSale {
    // --- CONSTRUCTOR

    constructor(
        uint256 _salePrice,
        address _funder,
        ERC20 _paymentToken,
        ERC20 _saleToken,
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
    {}

    // --- DISABLED FUNCTIONS

    function purchase(uint256) virtual override public {
        revert("Use purchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function whitelistedPurchase(uint256, bytes32[] calldata) override public pure {
        revert("Use purchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function withdrawGiveaway(bytes32[] calldata) override public pure {
        revert("Use withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)");
    }

    // --- WHITELISTED ACTIONS

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof,
        uint256 allocation
    ) public onlyDuringSale {
        // require that user is whitelisted by checking proof
        require(checkWhitelist(_msgSender(), merkleProof, allocation), 'proof invalid');

        uint256 remaining = getMaxPayment(_msgSender(), allocation);
        _purchase(paymentAmount, remaining);
    }

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)
        external
        nonReentrant
        onlyAfterSale
    {
        address user = _msgSender();
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(
            whitelistRootHash == 0 || checkWhitelist(user, merkleProof, allocation),
            'proof invalid'
        );

        uint256 saleTokenOwed = 0;
        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            saleTokenOwed = allocation;
            claimable[user] = allocation;
            totalPurchased[user] = allocation;
        }

        // send token and update states
        _withdraw(saleTokenOwed);
        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');
    }

    // --- HELPER FUNCTIONS

    // Returns true if user's allocation matches the one in merkle root, otherwise false
    function checkWhitelist(address user, bytes32[] calldata merkleProof, uint256 allocation)
        public
        view
        returns (bool)
    {
        // compute merkle leaf from input
        bytes32 leaf = keccak256(abi.encodePacked(user, allocation));

        // verify merkle proof
        return MerkleProof.verify(merkleProof, whitelistRootHash, leaf);
    }

    // @dev get max payment from allocation
    function getMaxPayment(address user, uint256 allocation) public view returns (uint256) {
        // get the maximum total payment for a user
        uint256 max = (salePrice * allocation) / SALE_PRICE_DECIMALS;
        if (maxTotalPayment < max) {
            max = maxTotalPayment;
        }

        // calculate and return remaining
        return max - paymentReceived[user];
    }
}
