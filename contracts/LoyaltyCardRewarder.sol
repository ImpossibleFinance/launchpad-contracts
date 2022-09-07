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
    ILoyaltyCardMaster public immutable loyaltyCardMaster;
    LoyaltyRewardsLookup public immutable rewardsLookup;

    error BatchRewardLengthsMismatch();

    constructor(address loyaltyCardMasterAddress, address loyaltyRewardsLookupAddress) {
        loyaltyCardMaster = ILoyaltyCardMaster(loyaltyCardMasterAddress);
        rewardsLookup = LoyaltyRewardsLookup(loyaltyRewardsLookupAddress);
    }

    /**
      @notice Reward a IF user account with the appropriate amount of reward points for a specific credential.
      @notice It's required that the user currently has a Loyalty Card NFT.
      @param account The IF user account that should be rewarderd
      @param credCode The numeric credential code of the credential that is to be rewarded
      @param credName The name of the credential that is to be rewarded
      @notice The reward is given based on the provided numeric credential code. 
        An additional check is performed to see if the provided credential name matches 
        the name which the rewards lookup contract itself associates with the provided credential code
      @dev Typically this would be called by the IF backend.
      @dev Can be part of a regular task (daily updates from KNN3) or an isolated call (user has completed a L&E quiz)
     */
    function rewardAccount(address account, uint256 credCode, string calldata credName) external onlyOwner {
        loyaltyCardMaster.addPointsAccount(account, rewardsLookup.getPoints(credCode, credName));
    }

    /**
        @notice Reward multiple IF users with the same amount of reward points, for the same credential
        @notice analogous to rewardAccount()
        @param accounts IF user account addresses, as an array
        @param credCode The numeric credential code that identifies given credential
        @param credName The name of given credential (to double check that given reward is intended)
     */
    function rewardBatchSingleCredential(
        address[] calldata accounts,
        uint256 credCode,
        string calldata credName
    ) external onlyOwner {
        loyaltyCardMaster.addPointsBatchAccSingleValue(accounts, rewardsLookup.getPoints(credCode, credName));
    }

    /**
        @notice Reward multiple IF users each having different amounts of reward points
        @notice When there's only 1 type of credential in the batch, use the SingleCredential version of this function
        @param accounts An array of unique IF user account addresses
        @param credCodes An array of credential code arrays, corresponding to the array of addresses. 
                        Each sub-array represents one IF user's fulfilled credentials codes (to be rewarded)
        @param credNames An array of credential name arrays, corresponding to the array of addresses.
                        Each sub-array represents one IF user's fulfilled credentials names (to be rewarded)
     */
    function rewardBatchMultiCredentials(
        address[] calldata accounts,
        uint256[][] calldata credCodes,
        string[][] calldata credNames
    ) external onlyOwner {
        if (accounts.length != credCodes.length) revert BatchRewardLengthsMismatch();
        if (accounts.length != credNames.length) revert BatchRewardLengthsMismatch();
        uint256[] memory pointsPerUser = new uint256[](credCodes.length);
        for (uint256 i = 0; i < credCodes.length; i++) {
            for (uint256 j = 0; j < credCodes[i].length; j++) {
                pointsPerUser[i] += rewardsLookup.getPoints(credCodes[i][j], credNames[i][j]);
            }
        }
        loyaltyCardMaster.addPointsBatchAccMultiValues(accounts, pointsPerUser);
    }
}
