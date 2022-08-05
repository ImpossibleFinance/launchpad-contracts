// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';

/**
  @title LoyaltyRewardsLookup is meant to be used in order to look up reward points given for any particular credential
  @notice This contract is needed by the LoyaltyCardRewarder so it can add to a loyalty card NFT the correct amount of points for any given credential.
  @dev We chose to have a separate lookup contract for a more flexible architecture with loose coupling.
 */

contract LoyaltyRewardsLookup is Ownable {
    enum Credential {
        KYC,
        LAUNCHPAD_STAKE_STANDARD_AURIGAMI,
        LAUNCHPAD_STAKE_UNLIMITED_AURIGAMI,
        LAUNCHPAD_STAKE_IDIA,
        LAUNCHPAD_STAKE_BLOCTO,
        LAUNCHPAD_STAKE_HIGHSTREET,
        LAUNCHPAD_STAKE_GENOPETS,
        LAUNCHPAD_PURCHASE,
        LAUNCHPAD_PURCHASE_FULL,
        DIAMOND_HAND1,
        DIAMOND_HAND2,
        FOLLOW_TWITTER,
        FOLLOW_INSTAGRAM,
        IN_TELEGRAM,
        IN_FRIENDS,
        IN_DISCORD,
        DISCORD_SCORE,
        SWAP_USER,
        SWAP_STAKER,
        SWAP_USER2, // traded over x
        SWAP_STAKER2, // staked over x
        REFERRAL,
        ATTEND_A_METAVERSE_EVENT,
        AURIGAMI_POWER_USER, // ( IF AURIGAMI + WNEAR/PLY LP STAKING % OR #)
        REGIONAL_MARKETING_CAMPAIGN,
        IF_MASTER, // complete l & e for impossible finance
        PROJECT_X_MASTER, // complete l & e for a project we launch/incubate
        IMPOSSIBLE_BULLS,
        VOTE_ON_DAO_PROPOSALS,
        SUBMIT_NEW_PROJ,
        SUBMIT_NEW_PROJ_WITH_HIGH_QUALITY,
        XX_POWER_USER,
        LAUNCHPAD_EARLY_STAKE,
        IMPOWER_ACCESS //（FROM RANKING BUT NOT IN LOYALTY)
    }

    mapping(Credential => uint256) public pointsForCredential;

    /// @notice Set how many reward points the fulfillment of a specific credential should bring the user   
    function setPoints(Credential cred, uint256 points) external onlyOwner {
        pointsForCredential[cred] = points;
    }
}
