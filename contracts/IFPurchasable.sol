// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';


/**
  @dev Abstract contract containing sale logics.
        To be implemented by IFSale.
  @notice Include virtual functions for regular and whitelisted purchase
  @notice Include virtual functions for whitelisted free token giveaway
  @notice Include sale state variables
  @notice Implemneted sale state variables changes on purchase and withdraw
 */
abstract contract IFPurchasable is Ownable, ReentrancyGuard {
    using SafeERC20 for ERC20;

    // --- SALE INFO

    // payment token
    ERC20 public immutable paymentToken;
    // price of the sale token
    uint256 public salePrice;
    // max for payment token amount
    uint256 public maxTotalPayment;
    // current purchased amount
    uint256 public saleTokenPurchased;
    // optional min for payment token amount
    uint256 public minTotalPayment;
    // optional max for total purchasable amount, default is 0 if there's no limit
    // assuming all users buy the token on the same price
    uint256 public maxTotalPurchasable;


    // --- USER INFO

    // tracks amount purchased by each address
    mapping(address => uint256) public paymentReceived;

    // promo code
    mapping(string => uint256) public amountPerCode;
    mapping(string => uint256) public uniqueUsePerCode;

    // --- STAT

    // counter of unique purchasers
    uint32 public purchaserCount;

    event Purchase(address indexed sender, uint256 paymentAmount);
    event PurchaseWithCode(address indexed sender, uint256 paymentAmount, string code);
    event SetMinTotalPayment(uint256 indexed minTotalPayment);
    event SetMaxTotalPurchasable(uint256 indexed _maxTotalPurchasable);

    constructor(
        ERC20 _paymentToken,
        uint256 _salePrice,
        uint256 _maxTotalPayment
    ) {
        require(
            _salePrice == 0 ||
                (_salePrice != 0 &&
                    address(_paymentToken) != address(0) &&
                    _maxTotalPayment >= _salePrice),
            'paymentToken or maxTotalPayment should not be 0 when salePrice is not 0'
        );
        salePrice = _salePrice; // can be 0 (for giveaway)
        paymentToken = _paymentToken; // can be 0 (for giveaway)
        maxTotalPayment = _maxTotalPayment; // can be 0 (for giveaway)
    }


    // Function for owner to set an optional, minTotalPayment
    // function setMinTotalPayment(uint256 _minTotalPayment) public onlyOwner onlyBeforeSale{
    function setMinTotalPayment(uint256 _minTotalPayment) public onlyOwner {
        // sale must not have started

        minTotalPayment = _minTotalPayment;

        emit SetMinTotalPayment(_minTotalPayment);
    }


    // Function for owner to set an optional, maxTotalPurchasable
    // The amount is calculated on salePrice. 
    function setMaxTotalPurchasable(uint256 _maxTotalPurchasable) public onlyOwner {
        maxTotalPurchasable = _maxTotalPurchasable * salePrice;

        require(maxTotalPurchasable >= saleTokenPurchased, 'Max purchasable should not be lower than the amount of token purchased');

        emit SetMaxTotalPurchasable(_maxTotalPurchasable);
    }

    // --- PURCHASE

    function purchase(uint256 paymentAmount) virtual public {}

    // Internal function for making purchase
    // Used by public functions `purchase`
    function _purchase(uint256 paymentAmount, uint256 remaining) virtual internal nonReentrant {
        // amount must be greater than minTotalPayment
        // by default, minTotalPayment is 0 unless otherwise set
        require(paymentAmount >= minTotalPayment, 'amount below min');

        // payment must not exceed remaining
        require(paymentAmount <= remaining, 'exceeds max payment');

        saleTokenPurchased += paymentAmount;
        require(maxTotalPurchasable == 0 || maxTotalPurchasable >= saleTokenPurchased, 'exceed max purchasable');

        // transfer specified amount from user to this contract
        paymentToken.safeTransferFrom(_msgSender(), address(this), paymentAmount);

        // if user is paying for the first time to this contract, increase counter
        if (paymentReceived[_msgSender()] == 0) purchaserCount += 1;

        // increase payment received amount
        paymentReceived[_msgSender()] += paymentAmount;

        emit Purchase(_msgSender(), paymentAmount);
    }

    function _purchaseWithCode(uint256 paymentAmount, uint256 remaining, string memory code) virtual internal {
        // This needs to be before anything else
        // ===
        _purchase(paymentAmount, remaining);
        // ====

        if (bytes(code).length > 0) {
            amountPerCode[code] += paymentAmount;
            if (paymentReceived[_msgSender()] == 0) {
                uniqueUsePerCode[code] += 1;
            }
        }

        emit PurchaseWithCode(_msgSender(), paymentAmount, code);
    }
}