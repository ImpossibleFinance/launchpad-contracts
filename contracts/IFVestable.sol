// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.9;

import '@openzeppelin/contracts/utils/math/Math.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

/**
  @dev Abstract contract containing vesting logics.
        To be implemented by IFSale.
  @notice There are two vesting types: linear and cliff
  @notice Can only set one vesting type
  @notice Once one of the vesting type is set, another one will be reset
  @notice Linear vesting unlocks tokens at a linear scale. Calculated by vesting end time
  @notice Cliff vesting unlocks tokens at a series of specific time. According to cliff period
 */
abstract contract IFVestable is Ownable {
    uint64 private constant TEN_YEARS = 315742060;

    // --- VESTING

    // Allow vesting to be editable after sale
    bool public vestingEditableOverride;

    // withdraw/cash delay timestamp (inclusive)
    uint256 public withdrawTime;
    // the most recent time the user claimed the saleToken
    mapping(address => uint256) public latestClaimTime;

    event OptInBuyback(address indexed user);

    // --- LINEAR VESTING

    // the time where the user can take all of the vested saleToken
    uint256 public linearVestingEndTime;
    event SetLinearVestingEndTime(uint256 indexed linearVestingEndTime);

    // --- CLIFF VESTING

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

    // --- CONSTRUCTOR

    constructor(
        uint256 _withdrawTime
    ) {
        // withdrawal delay is by default 0
        // it can be set using setWithdrawDelay
        withdrawTime = _withdrawTime;
    }

    // --- SETTER

    function setVestingEditable(bool _vestingEditableOverride) public onlyOwner {
        vestingEditableOverride = _vestingEditableOverride;
    }

    function setWithdrawTime(uint256 _withdrawTime) internal {
        withdrawTime = _withdrawTime;
    }

    // Function for owner to set a vesting end time
    function setLinearVestingEndTime(uint256 _linearVestingEndTime) virtual public onlyOwner {
        require(vestingEditableOverride || block.timestamp < withdrawTime, "Can't edit vesting after sale");
        require(_linearVestingEndTime > withdrawTime, "vesting end time has to be after withdrawal start time");
        require(withdrawTime > _linearVestingEndTime - TEN_YEARS, "vesting end time has to be within 10 years");
        linearVestingEndTime = _linearVestingEndTime;

        // unset cliff vesting
        delete cliffPeriod;
        emit SetLinearVestingEndTime(_linearVestingEndTime);
    }

    function setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct) virtual public onlyOwner {
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
        // pct is the release percentage, with a precision of 1%. Thus, the sum of all elements of pct must be equal to 100
        require(totalPct == 100, "total input percentage doesn't equal to 100");

        // unset linear vesting
        linearVestingEndTime = 0;
    }

    // --- VESTING LOGIC

    /**
      @notice Get the amount of token unlocked
      @param totalPurchased Total tokens purchased
      @param user Address of the user claiming the tokens
     */
    function getUnlockedToken(uint256 totalPurchased, uint256 claimable, address user) virtual public view returns (uint256) {
        // linear vesting
        if (linearVestingEndTime > block.timestamp) {
            // current claimable = total purchased * (now - last claimed time) / (total vesting time)
            return totalPurchased * (block.timestamp - Math.max(latestClaimTime[user], withdrawTime)) / (linearVestingEndTime - withdrawTime);
        }

        // cliff vesting
        uint256 cliffPeriodLength = cliffPeriod.length;
        if (cliffPeriodLength != 0 && (cliffPeriod[cliffPeriodLength - 1].claimTime > block.timestamp)) {
            uint8 claimablePct;
            for (uint8 i; i < cliffPeriodLength; i++) {
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

        // When vesting end, claim all of the remaining tokens.
        // Since all of the above calculations return a lower rounded number,
        // users will get a little bit less tokens.
        // Keeping track and returning the total remaining claimable makes sure the users will get the exact amount.
        return claimable;
    }
}