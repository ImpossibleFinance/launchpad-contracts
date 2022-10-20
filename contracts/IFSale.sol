// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "hardhat/console.sol";
import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import './IFFundable.sol';
import './IFSaleAbstract.sol';
import './IFVestable.sol';

contract IFAllocationSale is IFSaleAbstract, IFVestable, IFFundable {

    // CONSTRUCTOR
    constructor(
        address _funder,
        uint256 _salePrice,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint24 _trackId,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment,
        IIFRetrievableStakeWeight _allocationMaster,
        uint80 _allocSnapshotTimestamp
    ) 
        IFSaleAbstract(_salePrice, _paymentToken, _saleToken, _trackId, _startTime, _endTime, _maxTotalPayment, _funder) 
        IFVestable(_startTime, _endTime)
    { }

    // FUNCTIONS
    function setWithdrawDelay(uint256 _withdrawDelay) public onlyOwner onlyBeforeSale {
        setWithdrawTime(endTime + _withdrawDelay);
        super.setWithdrawDelay(_withdrawDelay);
    }

    function purchase(uint256 paymentAmount) external {
        _purchase(paymentAmount, type(uint256).max);
    }

    function withdraw() virtual public nonReentrant {
        // send token and update states
        uint256 tokenOwed = getCurrentClaimableToken();
        _withdraw(tokenOwed);
        // sale token owed must be greater than 0
        require(tokenOwed != 0, 'no token to be withdrawn');
    }   

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof
    ) external {
        // require that user is whitelisted by checking proof
        require(checkWhitelist(_msgSender(), merkleProof), 'proof invalid');

        _purchase(paymentAmount);
    }

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof)
        external
        nonReentrant
    {
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(
            whitelistRootHash == 0 || checkWhitelist(_msgSender(), merkleProof),
            'proof invalid'
        );

        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[_msgSender()]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            claimable[_msgSender()] = getUserStakeValue(_msgSender());
            totalPurchased[_msgSender()] = claimable[_msgSender()];
        }

        // send token and update states
        uint256 saleTokenOwed = sendSaleToken();
        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');
    }

    function getSaleTokensSold() internal returns (uint256 amount) {
        return (totalPaymentReceived * SALE_PRICE_DECIMALS) /
            salePrice;
    }

    function _purchase(uint256 paymentAmount, uint256 remaining) internal nonReentrant onlyDuringSale {
        super._purchase(paymentAmount, remaining);
        updateVestingOnPurchase((paymentReceived[_msgSender()] * SALE_PRICE_DECIMALS) / salePrice);
    }

    function _withdraw(uint256 tokenOwed) internal onlyDuringClaim {
        super._withdraw(tokenOwed);
        updateVestingOnWithdraw(tokenOwed);
    }

    // Returns true if user is on whitelist, otherwise false
    function checkWhitelist(address user, bytes32[] calldata merkleProof)
        public
        view
        returns (bool)
    {
        // compute merkle leaf from input
        bytes32 leaf = keccak256(abi.encodePacked(user));

        // verify merkle proof
        return MerkleProof.verify(merkleProof, whitelistRootHash, leaf);
    }
}
