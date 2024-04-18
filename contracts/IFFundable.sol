// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.9;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';

/**
  @title Abstract contract providing funder related functions in a sale
  @notice To be implemented by IFSale.
 */
abstract contract IFFundable is Ownable, ReentrancyGuard {
    using SafeERC20 for ERC20;

    // --- CONSTANTS

    // number of decimals of sale price
    uint64 constant SALE_PRICE_DECIMALS = 10**18;
    uint64 private constant ONE_HOUR = 3600;
    uint64 private constant ONE_YEAR = 31556926;
    uint64 private constant FIVE_YEARS = 157784630;
    uint64 private constant TEN_YEARS = 315742060;

    // --- OPERATOR ADDRESSES

    address public funder;
    // optional casher (settable by owner)
    address public casher;

    // --- SALE INFO

    // start timestamp when sale is active (inclusive)
    uint256 public immutable startTime;
    // end timestamp when sale is active (inclusive)
    uint256 public immutable endTime;
    // payment token
    ERC20 private immutable paymentToken;
    // sale token
    ERC20 private immutable saleToken;
    // withdraw/cash delay timestamp (inclusive)
    uint24 public withdrawDelay;
    // tracks whether user has already successfully withdrawn
    mapping(address => bool) public hasWithdrawn;

    // --- STATS

    // amount of sale token to sell
    uint256 public saleAmount;
    // tracks whether sale has been cashed
    bool public hasCashed;
    // total payment received for sale
    uint256 public totalPaymentReceived;
    // counter of unique withdrawers (doesn't count "cash"ing)
    uint32 public withdrawerCount;

    // --- CONSTRUCTOR

    constructor(
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint256 _startTime,
        uint256 _endTime,
        address _funder
    ) {
        // saleToken shouldn't be the same as paymentToken
        require(_saleToken != _paymentToken, 'saleToken = paymentToken');
        // when salePrice != 0, paymentToken and maxTotalPayment shouldn't be 0
        // sale token cannot be 0
        require(address(_saleToken) != address(0), '0x0 saleToken');
        // start timestamp must be in future
        require(block.timestamp < _startTime, 'start timestamp too early');
        require(_startTime - ONE_YEAR < block.timestamp, 'start time has to be within 1 year');
        // end timestamp must be after start timestamp
        require(_startTime < _endTime, 'end timestamp must be after start timestamp');
        require(_endTime - TEN_YEARS < _startTime, 'end time has to be within 10 years');

        require(_funder != address(0), '0x0 funder');
        funder = _funder;

        paymentToken = _paymentToken; // can be 0 (for giveaway)
        saleToken = _saleToken;
        startTime = _startTime;
        endTime = _endTime;
    }

    // --- MODIFIERS

    // Throws if called by any account other than the funder.
    modifier onlyFunder() {
        require(_msgSender() == funder, 'caller not funder');
        _;
    }

    // Throws if called by any account other than the casher.
    modifier onlyCasherOrOwner() {
        require(
            _msgSender() == casher || _msgSender() == owner(),
            'caller not casher or owner'
        );
        _;
    }

    // Throws if called during or after sale
    modifier onlyBeforeSale() {
        require(block.timestamp < startTime, 'sale already started');
        _;
    }

    // Throws if called outside of claim period
    modifier onlyAfterSale {
        require(block.timestamp > endTime + withdrawDelay, "can't withdraw before claim is started");
        _;
    }

    // Throws if called outside of sale period
    modifier onlyDuringSale {
        require(startTime <= block.timestamp, 'sale has not begun');
        require(block.timestamp <= endTime, 'sale over');
        _;
    }

    // --- EVENTS

    event SetCasher(address indexed casher);
    event SetFunder(address indexed funder);
    event Fund(address indexed sender, uint256 amount);
    event SetWithdrawDelay(uint24 indexed withdrawDelay);
    event Cash(
        address indexed sender,
        uint256 paymentTokenBalance,
        uint256 saleTokenBalance
    );
    event EmergencyTokenRetrieve(address indexed sender, uint256 amount);
    event Withdraw(address indexed sender, uint256 amount);

    // --- SETTER

    // Function for owner to set an optional, separate casher
    function setCasher(address _casher) public onlyOwner {
        casher = _casher;

        emit SetCasher(_casher);
    }

    function setFunder(address _funder) public onlyOwner {
        require(_funder != address(0), '0x0 funder');
        funder = _funder;

        emit SetFunder(_funder);
    }

    // Function for owner to set a withdraw delay
    function setWithdrawDelay(uint24 _withdrawDelay) virtual public onlyOwner onlyBeforeSale{
        require(_withdrawDelay < FIVE_YEARS, "withdrawDelay has to be within 5 years");
        withdrawDelay = _withdrawDelay;

        emit SetWithdrawDelay(_withdrawDelay);
    }

    // --- FUNDER'S LOGIC

    // Virtual function to be implemented by IFSale.
    //   To calculate the amount of cashable tokens.
    function getSaleTokensSold() internal virtual returns (uint256 amount);

    // Function for funding sale with sale token (called by project team)
    function fund(uint256 amount) public onlyFunder onlyBeforeSale{
        // transfer specified amount from funder to this contract
        saleToken.safeTransferFrom(_msgSender(), address(this), amount);

        // increase tracked sale amount
        saleAmount += amount;

        emit Fund(_msgSender(), amount);
    }


    // Function for funder to cash in payment token and unsold sale token
    function cash() external onlyCasherOrOwner onlyAfterSale {
        // prevent repeat cash
        require(!hasCashed, 'already cashed');

        hasCashed = true;

        // get amount of payment token received
        uint256 paymentTokenBal = paymentToken.balanceOf(address(this));

        // transfer all
        paymentToken.safeTransfer(_msgSender(), paymentTokenBal);

        // get amount of sale token on contract
        uint256 saleTokenBal = saleToken.balanceOf(address(this));

        // get amount of sold token
        uint256 totalTokensSold = getSaleTokensSold();

        // get principal (whichever is bigger between sale amount or amount on contract)
        uint256 principal = saleAmount < saleTokenBal
            ? saleTokenBal
            : saleAmount;

        // calculate amount of unsold sale token
        uint256 amountUnsold = principal - totalTokensSold;

        // transfer unsold
        saleToken.safeTransfer(_msgSender(), amountUnsold);

        emit Cash(_msgSender(), paymentTokenBal, amountUnsold);
    }

    function cashPaymentToken(uint256 amount) external onlyCasherOrOwner {
        // Get amount of payment token received
        uint256 paymentTokenBal = paymentToken.balanceOf(address(this));

        // Ensure there's enough payment tokens to cash
        require(paymentTokenBal >= amount, "No enough payment tokens to cash");

        // Transfer payment tokens to the caller
        paymentToken.safeTransfer(_msgSender(), amount);

        // Emit an event for this cashing
        emit Cash(_msgSender(), amount, 0);
    }


    // Retrieve tokens erroneously sent in to this address
    function emergencyTokenRetrieve(address token) public onlyOwner onlyAfterSale {
        // cannot be sale tokens
        require(token != address(saleToken));

        uint256 tokenBalance = ERC20(token).balanceOf(address(this));

        // transfer all
        ERC20(token).safeTransfer(_msgSender(), tokenBalance);

        emit EmergencyTokenRetrieve(_msgSender(), tokenBalance);
    }

    // Function for withdrawing purchased sale token after sale end
    function withdraw() virtual public nonReentrant {}   

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