// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.9;

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
    // flag to enable integer sale
    bool public isIntegerSale = false;

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

    function setIsIntegerSale(bool _isIntegerSale) public onlyOwner onlyBeforeSale {
        isIntegerSale = _isIntegerSale;
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
        // must not be a zero price sale
        require(salePrice != 0, 'use withdrawGiveaway');

        address user = _msgSender();

        uint256 tokenOwed = getCurrentClaimableToken(user);
        require(tokenOwed != 0, 'no token to be withdrawn');
        // send token and update states
        _withdraw(tokenOwed);
    }   

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof) virtual override public onlyAfterSale nonReentrant
    {
        revert("Not implemented");
    }

    // --- UPDATE SALE STATES

    function _purchase(uint256 paymentAmount, uint256 remaining) override internal {
        if (isIntegerSale) {
            require(isIntegerPayment(paymentAmount), 'can only buy integer amount of sale tokens');
        }
        totalPaymentReceived += paymentAmount;
        super._purchase(paymentAmount, remaining);
        // Update vesting variables
        uint256 tokenPurchased = (paymentReceived[_msgSender()] * SALE_PRICE_DECIMALS) / salePrice;
        totalPurchased[_msgSender()] = tokenPurchased;
        claimable[_msgSender()] = tokenPurchased;
    }

    function _withdraw(uint256 tokenOwed) override internal {
        // Update vesting variables
        latestClaimTime[_msgSender()] = block.timestamp;
        claimable[_msgSender()] -= tokenOwed;
        super._withdraw(tokenOwed);
    }

    // --- HELPER FUNCTIONS

    function getSaleTokensSold() override internal view returns (uint256 amount) {
        // if salePrice is 0, no tokens will be sold
        if (salePrice == 0) {
            return 0;
        } else {
            return totalPaymentReceived * SALE_PRICE_DECIMALS / salePrice;
        }
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

    // a function to check if the payment amount can buy integer amount of sale tokens, accounting the token decimals
    function isIntegerPayment(uint256 paymentAmount) public view returns (bool) {
        return (
            // has to at least purchase 1 token
            paymentAmount % salePrice == 0
            && paymentAmount > salePrice
        );
    }

    // Override the renounceOwnership function to disable it
    function renounceOwnership() public pure override{
        revert("ownership renunciation is disabled");
    }

}