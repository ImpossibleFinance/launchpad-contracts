// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "hardhat/console.sol";
import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import './IFFundable.sol';
import './IFSaleAbstract.sol';
import './IFVestable.sol';

contract IFSale is IFSaleAbstract, IFVestable, IFFundable {
    // CONSTRUCTOR
    constructor(
        address _funder,
        uint256 _salePrice,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint24 _trackId,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    ) 
        IFSaleAbstract(_paymentToken, _saleToken, _salePrice, _maxTotalPayment, _trackId)
        IFVestable(_endTime)
        IFFundable(_paymentToken, _saleToken, _startTime, _endTime, _funder)
    {}

    // FUNCTIONS

    function setWithdrawDelay(uint24 _withdrawDelay) override public onlyOwner onlyBeforeSale {
        setWithdrawTime(endTime + _withdrawDelay);
        super.setWithdrawDelay(_withdrawDelay);
    }

    function setVestingEndTime(uint256 _vestingEndTime) override public onlyOwner onlyBeforeSale {
        super.setVestingEndTime(_vestingEndTime);
    }

    function setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct) override public onlyOwner onlyBeforeSale {
        super.setCliffPeriod(claimTimes, pct);
    }

    function purchase(uint256 paymentAmount) virtual override public {
        require(whitelistRootHash == 0, 'use whitelistedPurchase');
        _purchase(paymentAmount, type(uint256).max);
    }

    function withdraw() virtual override public nonReentrant {
        // must not be a zero price sale
        require(salePrice != 0, 'use withdrawGiveaway');
        // send token and update states
        uint256 tokenOwed = getCurrentClaimableToken(_msgSender());
        _withdraw(tokenOwed);
        // sale token owed must be greater than 0
        require(tokenOwed != 0, 'no token to be withdrawn');
    }   

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof
    ) virtual override public {
        // require that user is whitelisted by checking proof
        require(checkWhitelist(_msgSender(), merkleProof), 'proof invalid');
        _purchase(paymentAmount, type(uint256).max);
    }

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof) virtual override public nonReentrant
    {
        address user = _msgSender();
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(
            whitelistRootHash == 0 || checkWhitelist(user, merkleProof),
            'proof invalid'
        );

        uint256 saleTokenOwed = getCurrentClaimableToken(user);
        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            // claimable[_msgSender()] = getUserStakeValue(_msgSender());
            claimable[user] = saleTokenOwed;
            totalPurchased[user] = saleTokenOwed;
        }

        // send token and update states
        _withdraw(saleTokenOwed);
        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');
    }

    function getCurrentClaimableToken(address user) internal view returns (uint256) {
        uint256 pct = getCurrentClaimablePercentage(user);
        if (pct == CLAIMABLE_PCT_DECIMAL) {
            return claimable[user];
        } else {
            return totalPurchased[user] * getCurrentClaimablePercentage(user) / CLAIMABLE_PCT_DECIMAL;
        }
    }

    // Returns true if user is on whitelist, otherwise false
    function checkWhitelist(address user, bytes32[] calldata merkleProof) virtual public view returns (bool)
    {
        // compute merkle leaf from input
        bytes32 leaf = keccak256(abi.encodePacked(user));

        // verify merkle proof
        return MerkleProof.verify(merkleProof, whitelistRootHash, leaf);
    }

    function getSaleTokensSold() override internal view returns (uint256 amount) {
        return (totalPaymentReceived * SALE_PRICE_DECIMALS) /
            salePrice;
    }

    function _purchase(uint256 paymentAmount, uint256 remaining) override internal onlyDuringSale {
        totalPaymentReceived += paymentAmount;
        super._purchase(paymentAmount, remaining);
        updateVestingOnPurchase((paymentReceived[_msgSender()] * SALE_PRICE_DECIMALS) / salePrice, _msgSender());
    }

    function _withdraw(uint256 tokenOwed) override internal onlyDuringClaim {
        require(endTime + withdrawDelay < block.timestamp, 'cannot withdraw yet');
        if (cliffPeriod.length != 0) {
            require(cliffPeriod[0].claimTime < block.timestamp, 'cannot withdraw yet');
        }
        super._withdraw(tokenOwed);
        updateVestingOnWithdraw(tokenOwed, _msgSender());
    }

    function updateVestingOnPurchase(uint256 tokenPurchased, address user) internal {
        totalPurchased[user] = tokenPurchased;
        claimable[user] = tokenPurchased;
    }

    function updateVestingOnWithdraw(uint256 tokenSent, address user) internal {
        // update claimable
        claimable[user] -= tokenSent;
        // update last claimed time
        latestClaimTime[user] = block.timestamp;
        // transfer giveaway sale token to participant
    }
}
