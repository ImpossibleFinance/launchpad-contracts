// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/math/Math.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

abstract contract IFVestable is Ownable {
    // 8 digits timestamp * 100 percent
    uint256 CLAIMABLE_PCT_DECIMAL = 10 ** 10 * 100;
    // seconds in 10 years
    uint64 private constant TEN_YEARS = 315569260;

    // store how many percentage of the token can be claimed at a certain cliff date
    struct Cliff {
        // the date when the percentage of token can be claimed
        uint256 claimTime;
        // the percentage token that can be claimed
        uint8 pct;
    }

    // withdraw/cash delay timestamp (inclusive)
    uint256 public withdrawTime;

    // the most recent time the user claimed the saleToken
    mapping(address => uint256) public latestClaimTime;
    // cliff vesting time and percentage
    Cliff[] public cliffPeriod;
    // the time where the user can take all of the vested saleToken
    uint256 public vestingEndTime;

    event SetVestingEndTime(uint256 indexed vestingEndTime);
    event SetCliffVestingPeriod(Cliff[] indexed cliffPeriod);

    constructor(
        // withdrawTIme is endTime + withdrawal delay 
        uint256 _withdrawTime
    ) {
        withdrawTime = _withdrawTime;
    }

    function setWithdrawTime(uint256 _withdrawTime) internal {
        withdrawTime = _withdrawTime;
    }

    // Function for owner to set a vesting end time
    function setVestingEndTime(uint256 _vestingEndTime) virtual public onlyOwner {
        require(_vestingEndTime > withdrawTime, "vesting end time has to be after withdrawal start time");
        require(withdrawTime > _vestingEndTime - TEN_YEARS, "vesting end time has to be within 10 years");
        vestingEndTime = _vestingEndTime;

        // unset cliff vesting
        delete cliffPeriod;

        // emit
        emit SetVestingEndTime(_vestingEndTime);
    }

    function setCliffPeriod(uint256[] calldata claimTimes, uint8[] calldata pct) virtual public onlyOwner {

        // lengths of claimTimes and pct must be equal
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
        vestingEndTime = 0;
    }

    function getCurrentClaimablePercentage(address user) public view returns (uint256) {
        // prevent returning a negative number
        require(block.timestamp > withdrawTime, 'claim not yet started');
        // linear vesting
        if (vestingEndTime > block.timestamp) {
            // current claimable = (now - last claimed time) / (total vesting time) * totalClaimable
            return CLAIMABLE_PCT_DECIMAL * (block.timestamp - Math.max(latestClaimTime[user], withdrawTime)) / (vestingEndTime - withdrawTime);
        }
        // cliff vesting
        uint256 cliffPeriodLength = cliffPeriod.length;
        // if cliff vesting is set  
        if (cliffPeriodLength != 0 && cliffPeriod[cliffPeriodLength - 1].claimTime > block.timestamp) {
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
            return CLAIMABLE_PCT_DECIMAL * claimablePct / 100;
        }
        // users can get all of the tokens after vestingEndTime
        return CLAIMABLE_PCT_DECIMAL;
    }

    function getCurrentClaimableToken(uint256 claimable, uint256 totalPurchased, address user) internal view returns (uint256) {
        uint256 pct = getCurrentClaimablePercentage(user);
        if (pct == CLAIMABLE_PCT_DECIMAL) {
            return claimable;
        } else {
            return totalPurchased * getCurrentClaimablePercentage(user) / CLAIMABLE_PCT_DECIMAL;
        }
    }

}