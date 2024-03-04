// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/math/Math.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

/**
  @dev Abstract contract containing vesting logics.
        To be implemented by IFSale.
  @notice There are two vesting types: linear and cliff
  @notice Non-dual mode:
  @notice - Can only set one vesting type
  @notice - Once one of the vesting type is set, another one will be unset
  @notice - Linear vesting unlocks tokens at a linear scale. Calculated by vesting end time
  @notice - Cliff vesting unlocks tokens at a series of specific time. According to cliff period
  @notice Dual mode:
  @notice - Can set both vesting type for dual mode
  @notice - Once dual mode is set, non dual mode params will be unset
  @notice Buyback:
  @notice - Only applicable for cliff vesting for now
  @notice - Users can opt in to buyback
  @notice - If user has opted in to buyback, they cannot claim their token after buybackClaimableNumber vesting phase
 */
abstract contract IFVestable is Ownable {
    uint64 private constant TEN_YEARS = 315360000;

    // --- VESTING PARAMS

    // Allow vesting to be editable after sale
    bool public vestingEditableOverride;

    // whether the user has opted in to buy back
    // if true, the user cannot claim the token after `buybackClaimableNumber` vesting phase
    // only applicable to cliff vesting for now
    mapping(address => bool) public hasOptInBuyback;

    // the number of vesting phase the user can claim the token if they have opted in to buyback
    // default 0
    uint256 public buybackClaimableNumber;

    // withdraw/cash delay timestamp (inclusive)
    uint256 public withdrawTime;
    // the most recent time the user claimed the saleToken
    mapping(address => uint256) public latestClaimTime;

    event OptInBuyback(address indexed user);
    event SetDualMode(bool indexed dualMode);
    event SetLinearVestingPctDualMode(uint8 indexed linearVestingPctDualMode);

    // --- LINEAR VESTING PARAMS

    // the time where the user can take all of the vested saleToken
    uint256 public linearVestingEndTime;
    event SetLinearVestingEndTime(uint256 indexed linearVestingEndTime);

    // --- CLIFF VESTING PARAMS

    // store how many percentage of the token can be claimed at a certain cliff date
    struct Cliff {
        // the date when the percentage of token can be claimed
        uint256 claimTime;
        // the percentage token that can be claimed
        uint8 pct;
    }
    // cliff vesting time and percentage
    Cliff[] public cliffPeriod;
    event SetCliffVestingPeriod(Cliff[] indexed cliffPeriod);

    function getCliffPeriod() public view returns (Cliff[] memory){
        return cliffPeriod;
    }

    // --- DUAL MODE PARAMS

    // Allow vesting on both cliff and linear mode
    bool public dualMode;
    uint256 public cliffVestingEndTimeDualMode;
    uint256 public linearVestingEndTimeDualMode;
    uint8 public linearVestingPctDualMode;


    // --- CONSTRUCTOR

    constructor(
        // withdrawTime is endTime + withdrawal delay 
        uint256 _withdrawTime
    ) {
        withdrawTime = _withdrawTime;
    }

    // --- SETTER

    /**
      @notice Set dual mode to enable both cliff and linear vesting
      @notice Also require dual mode vesting end time of cliff and linear vesting if dual mode is enabled
      @notice Making sure the admin knows if they want to enable or disable dual mode
      @param _dualMode: whether to enable dual mode
      @param _linearVestingEndTimeDualMode: end time of linear vesting
      @param _cliffVestingEndTimeDualMode: end time of cliff vesting
    */
    function setDualMode(
        bool _dualMode,
        uint256 _linearVestingEndTimeDualMode,
        uint256 _cliffVestingEndTimeDualMode
    ) public onlyOwner {
        require(vestingEditableOverride || block.timestamp < withdrawTime, "Can't edit vesting after sale");
        // Require that dual mode vesting end time of cliff and linear vesting if dual mode is going to be enabled
        // Cannot set dual mode vesting end time of cliff and linear vesting if dual mode is going to be disabled
        if (_dualMode) {
            require(_linearVestingEndTimeDualMode > withdrawTime, "vesting end time has to be after withdrawal start time");
            require(_cliffVestingEndTimeDualMode > withdrawTime, "vesting end time has to be after withdrawal start time");
        } else {
            require(_linearVestingEndTimeDualMode == 0 && _cliffVestingEndTimeDualMode == 0, "vesting end time has to be 0");
        }

        // Unset params
        if (_dualMode) {
            // unset non dual mode linear vesting
            linearVestingEndTime = 0;
        } else {
            // unset dual mode linear vesting percentage
            linearVestingPctDualMode = 0;
            // unset dual mode cliff vesting end time
            cliffVestingEndTimeDualMode = 0;
            // unset dual mode linear vesting end time
            linearVestingEndTimeDualMode = 0;
        }
        dualMode = _dualMode;
        emit SetDualMode(_dualMode);
    }

    function setVestingEditable(bool _vestingEditableOverride) public onlyOwner {
        vestingEditableOverride = _vestingEditableOverride;
    }

    function setWithdrawTime(uint256 _withdrawTime) internal {
        withdrawTime = _withdrawTime;
    }

    function setBuybackClaimableNumber(uint256 _buybackClaimableNumber) public onlyOwner {
        require(_buybackClaimableNumber < cliffPeriod.length, "buyback claimable number cannot exceed number of cliff period");
        buybackClaimableNumber = _buybackClaimableNumber;
    }

    // --- LINEAR VESTING SETTER

    /**
      @notice Set linear vesting end time. Only applicable for non dual mode
      @param _linearVestingEndTime: the time where the user can take all of the vested saleToken
     */
    function setLinearVestingEndTime(uint256 _linearVestingEndTime) virtual public onlyOwner {
        require(dualMode == false, "Must disable dual mode to set linear vesting end time");
        require(vestingEditableOverride || block.timestamp < withdrawTime, "Can't edit vesting after sale");
        require(_linearVestingEndTime > withdrawTime, "vesting end time has to be after withdrawal start time");
        require(withdrawTime > _linearVestingEndTime - TEN_YEARS, "vesting end time has to be within 10 years");
        linearVestingEndTime = _linearVestingEndTime;

        // unset cliff vesting
        delete cliffPeriod;
        emit SetLinearVestingEndTime(_linearVestingEndTime);
    }

    // --- CLIFF VESTING SETTER

    /**
      notice: Set cliff period. Only applicable for non dual mode
      @param claimTimes: the percentage of token that can be claimed at a certain cliff date
      @param pct: the percentage of token that can be claimed at a certain cliff date. Must be summed to 100
     */
    function setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct) virtual public onlyOwner {
        require(dualMode == false, "Must disable dual mode to set cliff period");
        _setCliffPeriod(claimTimes, pct, 100);
    }

    // --- DUAL MODE SETTER

    /**
      @notice Set dual mode parameters:
      @param _linearVestingPctDualMode: the percentage of token that can be claimed linearly
      @param _claimTimes: the percentage of token that can be claimed at a certain cliff date
      @param _pct: the percentage of token that can be claimed at a certain cliff date
     */
    function setDualModeParameters(uint8 _linearVestingPctDualMode, uint256[] calldata _claimTimes, uint8[] calldata _pct) virtual public onlyOwner {
        setLinearVestingPercentage(_linearVestingPctDualMode);
        _setCliffPeriodDualMode(_claimTimes, _pct);
    }

    /**
      @notice Internal function to set linear vesting percentage for dual mode
      @param _linearVestingPctDualMode: the percentage of token that can be claimed linearly
     */
    function setLinearVestingPercentage(uint8 _linearVestingPctDualMode) internal {
        require(vestingEditableOverride || block.timestamp < withdrawTime, "Can't edit vesting after sale");
        require(_linearVestingPctDualMode <= 100, "percentage cannot exceed 100");
        require(_linearVestingPctDualMode > 0, "percentage cannot be 0");
        require(dualMode == true, "Must enable dual mode to set linear vesting percentage");
        linearVestingPctDualMode = _linearVestingPctDualMode;
        emit SetLinearVestingPctDualMode(_linearVestingPctDualMode);
    }

    /**
      notice: Internal function to set cliff period on dual mode.
      @param claimTimes: the percentage of token that can be claimed at a certain cliff date
      @param pct: the percentage of token that can be claimed at a certain cliff date. Must be summed to 100 - linearVestingPctDualMode
     */
    function _setCliffPeriodDualMode(uint256[] calldata claimTimes, uint8[] calldata pct) internal {
        require(dualMode == true, "Must enable dual mode to set cliff period");
        _setCliffPeriod(claimTimes, pct, 100 - linearVestingPctDualMode);
    }

    // -- GENERAL SETTER

    /**
      notice: Set cliff period. Applicable for both dual mode and non dual mode
      @param claimTimes: the percentage of token that can be claimed at a certain cliff date
      @param pct: the percentage of token that can be claimed at a certain cliff date
      @param cliffTotalPct: the total percentage of token that can be claimed at a certain cliff date. Default 100 for non dual mode
     */
    function _setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct, uint8 cliffTotalPct) virtual public onlyOwner {
        require(vestingEditableOverride || block.timestamp < withdrawTime, "Can't edit vesting after sale");
        require(claimTimes.length == pct.length, "dates and pct doesn't match");
        require(claimTimes.length > 0, "input is empty");
        require(claimTimes.length <= 100, "input length cannot exceed 100");

        // clear the past entry
        delete cliffPeriod;

        uint256 maxDate;
        uint8 totalPct;
        require(claimTimes[0] > withdrawTime, "first claim time is before end time + withdraw delay");
        for (uint i = 0; i < claimTimes.length; i++) {
            require(maxDate < claimTimes[i], "dates not in ascending order");
            maxDate = claimTimes[i];
            totalPct += pct[i];
            cliffPeriod.push(Cliff(claimTimes[i], pct[i]));
        }
        require(withdrawTime > maxDate - TEN_YEARS, "vesting end time has to be within 10 years");
        // pct is the release percentage, with a precision of 1%. Thus, the sum of all elements of pct must be equal to cliffTotalPct.
        // cliffTotalPct is 100 for non dual mode and (100 - linearVestingPctDualMode) for dual mode
        require(totalPct == cliffTotalPct, "total percentage doesn't equal to 100");

        if (dualMode != true) {
            // unset linear vesting
            linearVestingEndTime = 0;
        }
    }

    // --- VESTING LOGIC

    // Opt in buyback. If called, the user will not be able to claim their token afer 
    // Emit an event OptInBuyback(user) if the user has successfully opted in
    // Only applicable to cliff vesting for now
    function optInBuyback() public {
        address user = _msgSender();
        require(hasOptInBuyback[user] == false, "user has already opted in");
        require(buybackClaimableNumber != 0, "buyback is not enabled");
        hasOptInBuyback[user] = true;
        emit OptInBuyback(user);
    }

    /**
      @notice Get the amount of token unlocked
      @param totalPurchased Total tokens purchased
      @param claimable Total claimable tokens
      @param user Address of the user claiming the tokens
     */
    function getUnlockedToken(uint256 totalPurchased, uint256 claimable, address user) virtual public view returns (uint256) {
        // Check if it is dual mode
        if (dualMode == true) {
            // Check if it is in linear vesting period
            if (linearVestingEndTimeDualMode > block.timestamp && cliffVestingEndTimeDualMode < block.timestamp) {
                return _getUnlockedTokenLinear(totalPurchased, user);
            // Check if it is in cliff vesting period
            } else if (cliffVestingEndTimeDualMode > block.timestamp && linearVestingEndTimeDualMode < block.timestamp) {
                return _getUnlockedTokenCliff(totalPurchased, user);
            }
            // If it is not in any vesting period, return claimable at the end
        } else {
            // Check if linear vesting is enabled and it is in linear vesting period
            if (linearVestingEndTime != 0 && linearVestingEndTime > block.timestamp) {
                return _getUnlockedTokenLinear(totalPurchased, user);
            // Check if cliff vesting is enabled and it is in cliff vesting period
            } else if (cliffPeriod.length != 0 && (cliffPeriod[cliffPeriod.length - 1].claimTime > block.timestamp)) {
                return _getUnlockedTokenCliff(totalPurchased, user);
            }
            // If it is not in any vesting period, return claimable at the end
        }

        // When vesting end, claim all of the remaining tokens.
        // Since all of the above calculations return a lower rounded number,
        // users will get a little bit less tokens.
        // Keeping track and returning the total remaining claimable makes sure the users will get the exact amount.
        return claimable;
    }

    // --- INTERNAL VESTING LOGIC
    /**
      @notice Get the amount of token unlocked in linear vesting
      @param totalPurchased Total tokens purchased
      @param user Address of the user claiming the tokens
     */
    function _getUnlockedTokenLinear(uint256 totalPurchased, address user) internal view returns (uint256) {
        // Always make sure if linearVestingEndTime > block.timestamp to prevent overflow
        uint256 _linearVestingEndTime;
        uint256 _withdrawTime;
        uint8 _linearVestingPct;
        if (dualMode) {
            _linearVestingEndTime = linearVestingEndTimeDualMode;
            _withdrawTime = Math.max(withdrawTime, cliffVestingEndTimeDualMode);
            _linearVestingPct = linearVestingPctDualMode;
        } else {
            _linearVestingEndTime = linearVestingEndTime;
            _withdrawTime = withdrawTime;
            _linearVestingPct = 100;
        }

        if (_linearVestingEndTime > block.timestamp) {
            // current claimable = total purchased * (now - last claimed time) / (total vesting time) * linearVestingPctDualMode / 100
            return totalPurchased * (block.timestamp - Math.max(latestClaimTime[user], _withdrawTime)) / (_linearVestingEndTime - withdrawTime) * _linearVestingPct / 100;
        }

        // Default return 0 if linear vesting end time is over
        // Will be handled by getUnlockedToken function to return claimable at the end
        return 0;
    }

    /**
      @notice Get the amount of token unlocked in cliff vesting
      @param totalPurchased Total tokens purchased
      @param user Address of the user claiming the tokens
     */
    function _getUnlockedTokenCliff(uint256 totalPurchased, address user) internal view returns (uint256) {
        uint256 cliffPeriodLength = cliffPeriod.length;
        // Always make sure the last cliff period > block.timestamp to prevent miscalculation
        // Calculate the claimable percentage by looping through the cliff period
        // If the user has opted in to buyback, the user can only claim the token before buybackClaimableNumber vesting phase
        // buybackClaimableNumber is 0 means buyback is not enabled. Thus, user can claim all of the token
        if (cliffPeriodLength != 0 && (cliffPeriod[cliffPeriodLength - 1].claimTime > block.timestamp || hasOptInBuyback[user] == true)) {
            uint8 claimablePct;
            for (uint8 i; i < cliffPeriodLength; i++) {
                // if hasOptInBuyback is true, user's claimable phase will be limited by buyBackClaimableNumber
                // buyBackClaimableNumber is 0 means buyback is not enabled. Thus, user can claim all of the token
                // if (hasOptInBuyback[user] == true && buybackClaimableNumber <= i && buybackClaimableNumber != 0) {
                if (hasOptInBuyback[user] == true && buybackClaimableNumber <= i) {
                    return totalPurchased * claimablePct / 100;
                }
                // if the cliff timestamp has been passed, add the claimable percentage
                if (cliffPeriod[i].claimTime > block.timestamp) { break; }
                if (latestClaimTime[user] < cliffPeriod[i].claimTime) {
                    claimablePct += cliffPeriod[i].pct;
                }
            }
            // current claimable = total * claimiable percentage
            if (claimablePct == 0) {
                return 0;
            }
            return totalPurchased * claimablePct / 100;
        }

        // Default return 0 if cliff vesting end time is over
        // Will be handled by getUnlockedToken function to return claimable at the end
        return 0;
    }
}