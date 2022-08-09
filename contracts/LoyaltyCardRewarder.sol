// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "./interfaces/ILoyaltyCardMaster.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "./LoyaltyRewardsLookup.sol";

/**
  @title LoyaltyCardRewarder is meant to be used as a reward points - giving operator on the Impossible Finance LoyaltyCardMaster contract
  @notice This contract can add loyalty points for benefits associated to the loyalty card nft holders.
  @dev This contract uses the LoyaltyRewardsLookup contract to determine how many points to add for a given credential
 */

contract LoyaltyCardRewarder is Ownable {
    ILoyaltyCardMaster public loyaltyCardMaster;
    LoyaltyRewardsLookup public rewardsLookup;

    constructor(address loyaltyCardMasterAddress, address loyaltyRewardsLookupAddress) {
        loyaltyCardMaster = ILoyaltyCardMaster(loyaltyCardMasterAddress);
        rewardsLookup = LoyaltyRewardsLookup(loyaltyRewardsLookupAddress);
    }

    /**
      @notice Reward a IF user account with the appropriate amount of reward points for a specific credential.
      @notice It's required that the user currently has a Loyalty Card NFT.
      @param account The IF user account that should be rewarderd
      @param cred The credential that is to be rewarded
      @dev Typically this would be called by the IF backend.
      @dev Can be part of a regular task (daily updates from KNN3) or an isolated call (user has completed a L&E quiz)
     */
    function rewardAccount(address account, LoyaltyRewardsLookup.Credential cred) external onlyOwner {
        loyaltyCardMaster.addPointsAccount(account, rewardsLookup.pointsForCredential(cred));
    }
}
