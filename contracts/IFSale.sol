// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import './IFFundable.sol';
import './IFPurchasable.sol';
import './IFVestable.sol';
import './IFWhitelistable.sol';

/**
  @dev Vanilla Sale contract compositing vesting, funding, and whitelisting functions.
  @notice Features:
  @notice 1. Funder related actions like fund and cash
  @notice 2. Regular and whitelisted purchase
  @notice 3. Whitelisted free token giveaway
  @notice 4. Vest tokens in linear or cliff mode
 */
contract IFSale is IFPurchasable, IFVestable, IFFundable, IFWhitelistable {
    // tracks amount of tokens owed to each address
    mapping(address => uint256) public claimable;
    // tracks amount of tokens purchased by each address
    mapping(address => uint256) public totalPurchased;

    // --- CONSTRUCTOR

    constructor(
        address _funder,
        uint256 _salePrice,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    ) 
        IFPurchasable(_paymentToken, _salePrice, _maxTotalPayment)
        IFVestable(_endTime)
        IFFundable(_paymentToken, _saleToken, _startTime, _endTime, _funder)
        IFWhitelistable()
    {}

    // --- SETTERS

    function setWithdrawDelay(uint24 _withdrawDelay) override public onlyOwner onlyBeforeSale {
        setWithdrawTime(endTime + _withdrawDelay);
        super.setWithdrawDelay(_withdrawDelay);
    }

    function setLinearVestingEndTime(uint256 _vestingEndTime) override public onlyOwner onlyBeforeSale {
        super.setLinearVestingEndTime(_vestingEndTime);
    }

    function setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct) override public onlyOwner onlyBeforeSale {
        super.setCliffPeriod(claimTimes, pct);
    }

    // --- PURCHASE

    function purchase(uint256 paymentAmount) virtual override public onlyDuringSale {
        require(whitelistRootHash == 0, 'use whitelistedPurchase');
        _purchase(paymentAmount, maxTotalPayment);
    }

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof
    ) virtual override public onlyDuringSale {
        // the user has to be whitelisted
        require(checkWhitelist(_msgSender(), merkleProof), 'proof invalid');
        _purchase(paymentAmount, maxTotalPayment);
    }

    // --- WITHDRAW

    function withdraw() virtual override public onlyAfterSale nonReentrant {
        address user = _msgSender();
        // must not be a zero price sale
        require(salePrice != 0, 'use withdrawGiveaway');

        uint256 tokenOwed = getCurrentClaimableToken(user);
        // send token and update states
        _withdraw(tokenOwed);
        require(tokenOwed != 0, 'no token to be withdrawn');
    }   

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof) virtual override public onlyAfterSale nonReentrant
    {
        address user = _msgSender();
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(whitelistRootHash == 0 || checkWhitelist(user, merkleProof), 'proof invalid');

        uint256 tokenOwed = getCurrentClaimableToken(user);
        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            claimable[user] = tokenOwed;
            totalPurchased[user] = tokenOwed;
        }
        // send token and update states
        _withdraw(tokenOwed);
        require(tokenOwed != 0, 'withdraw giveaway amount 0');
    }

    // --- UPDATE SALE STATES

    function _purchase(uint256 paymentAmount, uint256 remaining) override internal {
        totalPaymentReceived += paymentAmount;
        super._purchase(paymentAmount, remaining);
        // Update vesting variables
        uint256 tokenPurchased = (paymentReceived[_msgSender()] * SALE_PRICE_DECIMALS) / salePrice;
        totalPurchased[_msgSender()] = tokenPurchased;
        claimable[_msgSender()] = tokenPurchased;
    }

    function _withdraw(uint256 tokenOwed) override internal {
        super._withdraw(tokenOwed);
        // Update vesting variables
        latestClaimTime[_msgSender()] = block.timestamp;
        claimable[_msgSender()] -= tokenOwed;
    }

    // --- HELPER FUNCTIONS

    function getSaleTokensSold() override internal view returns (uint256 amount) {
        return (totalPaymentReceived * SALE_PRICE_DECIMALS) /
            salePrice;
    }

    // A helper function to get the amount of unlocked token by providing user's address
    function getCurrentClaimableToken (address user) public view returns (uint256) {
        return getUnlockedToken(totalPurchased[user], claimable[user], user);
    }

    // Returns true if user is on whitelist, otherwise false
    function checkWhitelist(address user, bytes32[] calldata merkleProof) virtual public view returns (bool)
    {
        // compute merkle leaf from input
        bytes32 leaf = keccak256(abi.encodePacked(user));

        // verify merkle proof
        return MerkleProof.verify(merkleProof, whitelistRootHash, leaf);
    }
}