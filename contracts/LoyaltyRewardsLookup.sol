// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';

/**
  @title LoyaltyRewardsLookup is meant to be used in order to look up reward points given for any particular credential
  @notice This contract is needed by the LoyaltyCardRewarder so it can add to a loyalty card NFT the correct amount of points for any given credential.
  @dev We chose to have a separate lookup contract for a more flexible architecture with loose coupling.
 */

contract LoyaltyRewardsLookup is Ownable {

    // ================ FLEXIBLE DESIGN ============== // 

    /// @dev For on-chain transparency we track not only the points, both also the name associated with a credential
    struct Credential {
        string name;
        uint256 points;
    }

    /// @dev credential code => credential data (name and points)
    mapping(uint256 => Credential) public credentialByCode;

    event SetCredential(uint256 code, string name, uint256 points);
    event UpdateCredentialName(uint256 code, string name);
    event UpdateCredentialPoints(uint256 code, uint256 points);

    error CredentialMismatch();

    modifier intendedCredential(uint256 credCode, string calldata credName) {
        if (keccak256(abi.encodePacked(credentialByCode[credCode].name)) != 
            keccak256(abi.encodePacked((credName)))) revert CredentialMismatch();
        _;
    }

    /// @notice Set the name of a credential (based on its numeric code) and how many 
    /// reward points the fulfillment should bring the user
    /// @param credCode Credential code
    /// @param points How many reward points are associated   
    /// @param credName The name of the credential
    function setCredential(uint256 credCode, uint256 points, string calldata credName)
        external
        onlyOwner
    {
        credentialByCode[credCode].name = credName;
        credentialByCode[credCode].points = points;
        emit SetCredential(credCode, credName, points);
    }

    /// @notice Update the name of a specific credential, as it identified by its credential code
    /// @param credCode The numeric credential code
    /// @param credName The credential name
    function updateCredentialName(uint256 credCode, string memory credName) external onlyOwner {
        credentialByCode[credCode].name = credName;
        emit UpdateCredentialName(credCode, credName);
    }

    /// @notice Update how many reward points the fulfillment of a specific credential should bring the user
    /// while also ensuring it's the intended credential based on its name
    /// @param credCode Credential code
    /// @param points How many reward points are associated   
    /// @param credName The name of the credential
    function updateCredentialPoints(uint256 credCode, uint256 points, string calldata credName) 
        external 
        onlyOwner 
        intendedCredential(credCode, credName) 
    {
        credentialByCode[credCode].points = points;
    }

    /// @notice Retrieve a credential's name based on its code
    function getName(uint256 credCode)
        external view
        returns (string memory) 
    {
        return credentialByCode[credCode].name;
    }

    /// @notice Retrieve a credential's reward points based on its code
    /// while also ensuring it's the intended credential based on its name
    function getPoints(uint256 credCode, string calldata credName)
        external view 
        intendedCredential(credCode, credName)
        returns (uint256) 
    {
        return credentialByCode[credCode].points;
    }

    // ================ INITIAL DESIGN ============== // 


    //               (lacks flexibility)                     

    // enum Credential {
    //     KYC,
    //     LAUNCHPAD_STAKE_STANDARD_AURIGAMI,
    //     LAUNCHPAD_STAKE_UNLIMITED_AURIGAMI,
    //     LAUNCHPAD_STAKE_IDIA,
    //     LAUNCHPAD_STAKE_BLOCTO,
    //     LAUNCHPAD_STAKE_HIGHSTREET,
    //     LAUNCHPAD_STAKE_GENOPETS,
    //     LAUNCHPAD_PURCHASE,
    //     LAUNCHPAD_PURCHASE_FULL,
    //     DIAMOND_HAND1,
    //     DIAMOND_HAND2,
    //     FOLLOW_TWITTER,
    //     FOLLOW_INSTAGRAM,
    //     IN_TELEGRAM,
    //     IN_FRIENDS,
    //     IN_DISCORD,
    //     DISCORD_SCORE,
    //     SWAP_USER,
    //     SWAP_STAKER,
    //     SWAP_USER2, // traded over x
    //     SWAP_STAKER2, // staked over x
    //     REFERRAL,
    //     ATTEND_A_METAVERSE_EVENT,
    //     AURIGAMI_POWER_USER, // ( IF AURIGAMI + WNEAR/PLY LP STAKING % OR #)
    //     REGIONAL_MARKETING_CAMPAIGN,
    //     IF_MASTER, // complete l & e for impossible finance
    //     PROJECT_X_MASTER, // complete l & e for a project we launch/incubate
    //     IMPOSSIBLE_BULLS,
    //     VOTE_ON_DAO_PROPOSALS,
    //     SUBMIT_NEW_PROJ,
    //     SUBMIT_NEW_PROJ_WITH_HIGH_QUALITY,
    //     XX_POWER_USER,
    //     LAUNCHPAD_EARLY_STAKE,
    //     IMPOWER_ACCESS //（FROM RANKING BUT NOT IN LOYALTY)
    // }

    // mapping(Credential => uint256) public credentialToPoints;

    // /// @notice Set how many reward points the fulfillment of a specific credential should bring the user   
    // function setPoints(Credential cred, uint256 points) external onlyOwner {
    //     credentialToPoints[cred] = points;
    // }
}
