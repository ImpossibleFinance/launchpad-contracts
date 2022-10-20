// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';

abstract contract IFFundable is Ownable, ReentrancyGuard {
    using SafeERC20 for ERC20;
    // CONSTANTS

    // number of decimals of sale price
    uint64 constant SALE_PRICE_DECIMALS = 10**18;
    // seconds in 1 hours
    uint64 constant ONE_HOUR = 3600;
    // seconds in 1 year
    uint64 constant ONE_YEAR = 31556926;
    // seconds in 5 years
    uint64 constant FIVE_YEARS = 157784630;
    // seconds in 10 years
    uint64 constant TEN_YEARS = 315569260;

    // Operators Info
    // funder
    address public funder;
    // optional casher (settable by owner)
    address public casher;

    // Sale Info

    // start timestamp when sale is active (inclusive)
    uint256 private immutable startTime;
    // end timestamp when sale is active (inclusive)
    uint256 private immutable endTime;
    // payment token
    ERC20 public immutable paymentToken;
    // sale token
    ERC20 public immutable saleToken;
    // amount of sale token to sell
    uint256 public saleAmount;
    // withdraw/cash delay timestamp (inclusive)
    uint24 private withdrawDelay;

    // STAT

    // tracks whether sale has been cashed
    bool public hasCashed;
    // total payment received for sale
    uint256 public totalPaymentReceived;

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
        require(_startTime < _endTime - ONE_HOUR, 'end timestamp before start should be least 1 hour');
        require(_endTime - TEN_YEARS < _startTime, 'end time has to be within 10 years');

        require(_funder != address(0), '0x0 funder');
        funder = _funder;

        paymentToken = _paymentToken; // can be 0 (for giveaway)
        saleToken = _saleToken;
        startTime = _startTime;
        endTime = _endTime;
    }

    // MODIFIERS

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

    // Function for owner to set an optional, separate casher
    function setCasher(address _casher) public onlyOwner {
        casher = _casher;

        // emit
        emit SetCasher(_casher);
    }

    modifier onlyBeforeSale() {
        require(block.timestamp < startTime, 'sale already started');
        _;
    }

    modifier onlyDuringClaim {
        require(block.timestamp > endTime + withdrawDelay, "can't withdraw before claim is started");
        _;
    }

    modifier onlyDuringSale {
        require(startTime <= block.timestamp, 'sale has not begun');
        require(block.timestamp <= endTime, 'sale over');
        _;
    }

    event SetCasher(address indexed casher);
    event Fund(address indexed sender, uint256 amount);
    event SetWithdrawDelay(uint24 indexed withdrawDelay);
    event Cash(
        address indexed sender,
        uint256 paymentTokenBalance,
        uint256 saleTokenBalance
    );
    event EmergencyTokenRetrieve(address indexed sender, uint256 amount);

    function getSaleTokensSold() internal virtual returns (uint256 amount);

    // Function for funding sale with sale token (called by project team)
    function fund(uint256 amount) public onlyFunder onlyBeforeSale{
        // transfer specified amount from funder to this contract
        saleToken.safeTransferFrom(_msgSender(), address(this), amount);

        // increase tracked sale amount
        saleAmount += amount;

        // emit
        emit Fund(_msgSender(), amount);
    }

    // Function for owner to set a withdraw delay
    function setWithdrawDelay(uint24 _withdrawDelay) public onlyOwner onlyBeforeSale{
        require(_withdrawDelay < FIVE_YEARS, "withdrawDelay has to be within 5 years");
        withdrawDelay = _withdrawDelay;

        // emit
        emit SetWithdrawDelay(_withdrawDelay);
    }

    // Function for funder to cash in payment token and unsold sale token
    function cash() external onlyCasherOrOwner {
        // must be past end timestamp plus withdraw delay
        require(
            endTime + withdrawDelay < block.timestamp,
            'cannot withdraw yet'
        );
        // prevent repeat cash
        require(!hasCashed, 'already cashed');

        // set hasCashed to true
        hasCashed = true;

        // get amount of payment token received
        uint256 paymentTokenBal = paymentToken.balanceOf(address(this));

        // transfer all
        paymentToken.safeTransfer(_msgSender(), paymentTokenBal);

        // get amount of sale token on contract
        uint256 saleTokenBal = saleToken.balanceOf(address(this));

        // get amount of sold token
        uint256 totalTokensSold = getTokensSold();

        // get principal (whichever is bigger between sale amount or amount on contract)
        uint256 principal = saleAmount < saleTokenBal
            ? saleTokenBal
            : saleAmount;

        // calculate amount of unsold sale token
        uint256 amountUnsold = principal - totalTokensSold;

        // transfer unsold
        saleToken.safeTransfer(_msgSender(), amountUnsold);

        // emit
        emit Cash(_msgSender(), paymentTokenBal, amountUnsold);
    }

    // retrieve tokens erroneously sent in to this address
    function emergencyTokenRetrieve(address token) public onlyOwner onlyDuringClaim {
        // cannot be sale tokens
        require(token != address(saleToken));

        uint256 tokenBalance = ERC20(token).balanceOf(address(this));

        // transfer all
        ERC20(token).safeTransfer(_msgSender(), tokenBalance);

        // emit
        emit EmergencyTokenRetrieve(_msgSender(), tokenBalance);
    }
}