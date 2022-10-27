// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';

abstract contract IFSaleAbstract is Ownable, ReentrancyGuard {
    using SafeERC20 for ERC20;

    // optional whitelist setter (settable by owner)
    address public whitelistSetter;

    // --- SALE INFO

    // payment token
    ERC20 public immutable paymentToken;
    // sale token
    ERC20 public immutable saleToken;
    // price of the sale token
    uint256 public salePrice;
    // max for payment token amount
    uint256 public maxTotalPayment;
    // optional min for payment token amount
    uint256 public minTotalPayment;
    // whitelist merkle root; if not set, then sale is open to everyone that has allocation
    bytes32 public whitelistRootHash;
    // track id
    uint24 public trackId;

    // --- USER INFO

    // tracks amount purchased by each address
    mapping(address => uint256) public paymentReceived;
    // tracks whether user has already successfully withdrawn
    mapping(address => bool) public hasWithdrawn;
    // tracks amount of tokens owed to each address
    mapping(address => uint256) public claimable;
    // tracks amount of tokens purchased by each address
    mapping(address => uint256) public totalPurchased;

    // --- STAT

    // counter of unique purchasers
    uint32 public purchaserCount;
    // counter of unique withdrawers (doesn't count "cash"ing)
    uint32 public withdrawerCount;

    event Purchase(address indexed sender, uint256 paymentAmount);
    event Withdraw(address indexed sender, uint256 amount);
    event SetWhitelistSetter(address indexed whitelistSetter);
    event SetWhitelist(bytes32 indexed whitelistRootHash);
    event SetMinTotalPayment(uint256 indexed minTotalPayment);

    constructor(
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint256 _salePrice,
        uint256 _maxTotalPayment,
        uint24 _trackId
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
        saleToken = _saleToken;
        maxTotalPayment = _maxTotalPayment; // can be 0 (for giveaway)
        trackId = _trackId; // can be 0 (with allocation override)
    }

    // Throws if called by any account other than the whitelist setter.
    modifier onlyWhitelistSetterOrOwner() {
        require(
            _msgSender() == whitelistSetter || _msgSender() == owner(),
            'caller not whitelist setter or owner'
        );
        _;
    }

    // --- SETTER

    // Function for owner to set an optional, separate whitelist setter
    function setWhitelistSetter(address _whitelistSetter) public onlyOwner {
        whitelistSetter = _whitelistSetter;

        emit SetWhitelistSetter(_whitelistSetter);
    }

    // Function for owner or whitelist setter to set a whitelist; if not set, then everyone allowed
    function setWhitelist(bytes32 _whitelistRootHash)
        public
        onlyWhitelistSetterOrOwner
    {
        whitelistRootHash = _whitelistRootHash;

        emit SetWhitelist(_whitelistRootHash);
    }

    // Function for owner to set an optional, minTotalPayment
    // function setMinTotalPayment(uint256 _minTotalPayment) public onlyOwner onlyBeforeSale{
    function setMinTotalPayment(uint256 _minTotalPayment) public onlyOwner {
        // sale must not have started

        minTotalPayment = _minTotalPayment;

        emit SetMinTotalPayment(_minTotalPayment);
    }

    // --- SALE FUNCTIONS

    function purchase(uint256 paymentAmount) virtual public {}

    // Function for withdrawing purchased sale token after sale end
    function withdraw() virtual public nonReentrant {}   

    // purchase function when there is a whitelist
    function whitelistedPurchase(uint256 paymentAmount, bytes32[] calldata merkleProof) virtual public {}

    function withdrawGiveaway(bytes32[] calldata merkleProof) virtual public nonReentrant {}

    // --- INTENAL HELPER FUNCTIONS

    // Internal function for making purchase
    // Used by public functions `purchase`
    // function _purchase(uint256 paymentAmount, uint256 remaining) internal nonReentrant onlyDuringSale {
    function _purchase(uint256 paymentAmount, uint256 remaining) virtual internal nonReentrant {
        // amount must be greater than minTotalPayment
        // by default, minTotalPayment is 0 unless otherwise set
        require(paymentAmount >= minTotalPayment, 'amount below min');

        // payment must not exceed remaining
        require(paymentAmount <= remaining, 'exceeds max payment');

        // transfer specified amount from user to this contract
        paymentToken.safeTransferFrom(_msgSender(), address(this), paymentAmount);

        // // increase total payment received amount
        // totalPaymentReceived += paymentAmount;

        // if user is paying for the first time to this contract, increase counter
        if (paymentReceived[_msgSender()] == 0) purchaserCount += 1;

        // increase payment received amount
        paymentReceived[_msgSender()] += paymentAmount;

        emit Purchase(_msgSender(), paymentAmount);
    }

    function _withdraw(uint256 saleTokenOwed) virtual internal {
        require(saleTokenOwed != 0, 'no token to be withdrawn');

        // increment withdrawer count
        if (!hasWithdrawn[_msgSender()]) {
            withdrawerCount += 1;
            // set withdrawn to true
            hasWithdrawn[_msgSender()] = true;
        }

        saleToken.safeTransfer(_msgSender(), saleTokenOwed);

        emit Withdraw(_msgSender(), saleTokenOwed);
    }
}