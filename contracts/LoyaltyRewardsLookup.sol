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
        uint256 points;
        string name;
    }

    /// @dev credential code => credential data (name and points)
    mapping(uint256 => Credential) public credentialByCode;

    /// @dev credential name => whether it's already in use (i.e. associated with a credential code)
    mapping(string => bool) public credentialNameInUse;

    event SetCredential(uint256 code, string name, uint256 points);
    event UpdateCredentialName(uint256 code, string name);
    event UpdateCredentialPoints(uint256 code, uint256 points);

    error CredentialMismatch();
    error DuplicateCredentialName();
    error EmptyCredentialName();
    error ZeroCredentialPoints();
    error CredentialCodeAlreadyInUse();
    error CredentialNotSet(uint256);

    modifier intendedCredential(uint256 credCode, string calldata credName) {
        if (keccak256(abi.encodePacked(credentialByCode[credCode].name)) != 
            keccak256(abi.encodePacked((credName)))) revert CredentialMismatch();
        _;
    }

    modifier noDuplicate(string calldata credName) {
        if (credentialNameInUse[credName]) revert DuplicateCredentialName();
        _;
    }

    modifier nonEmpty(string calldata credName) {
        if (bytes(credName).length == 0) revert EmptyCredentialName();
        _;
    }

    modifier nonZero(uint256 points) {
        if (points == 0) revert ZeroCredentialPoints();
        _;
    }

    modifier onlySetCredential(uint256 credCode) {
        string memory name = credentialByCode[credCode].name;
        if (bytes(name).length == 0) revert CredentialNotSet(credCode);
        _;
    }

    /// @notice Set the name of a credential (based on its numeric code) and how many 
    /// reward points the fulfillment should bring the user
    /// @notice Works as a DB insert - only as "first write" for given credential code
    /// @param credCode Credential code
    /// @param points How many reward points are associated   
    /// @param credName The name of the credential
    function setCredential(uint256 credCode, uint256 points, string calldata credName)
        external
        onlyOwner
        noDuplicate(credName)
        nonEmpty(credName)
        nonZero(points)
    {
        /// @dev since names must be non empty, the checked condition is reliable
        if (bytes(credentialByCode[credCode].name).length > 0) revert CredentialCodeAlreadyInUse();
        credentialByCode[credCode].name = credName;
        credentialByCode[credCode].points = points;
        credentialNameInUse[credName] = true;
        emit SetCredential(credCode, credName, points);
    }

    /// @notice Update the name of a specific credential, as it is identified by its credential code
    /// @param credCode The numeric credential code
    /// @param credName The credential name
    function updateCredentialName(uint256 credCode, string calldata credName) 
        external
        onlyOwner
        noDuplicate(credName)
        nonEmpty(credName)
    {
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
        nonZero(points)
    {
        credentialByCode[credCode].points = points;
        emit UpdateCredentialPoints(credCode, points);
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
        onlySetCredential(credCode)
        intendedCredential(credCode, credName)
        returns (uint256) 
    {
        return credentialByCode[credCode].points;
    }

    /// @notice Convenience function to read all credential info for given codes
    /// @dev Each tuple from return data is decoded as a 2-element array with ethers.js
    function getNamesAndPointsForAll(uint256[] calldata credCodes)
        external view
        returns (Credential[] memory credentials) 
    {
        credentials = new Credential[](credCodes.length);
        for (uint256 i = 0; i < credCodes.length; i++) {
            string memory credName = credentialByCode[credCodes[i]].name;
            if (bytes(credName).length == 0) revert CredentialNotSet(credCodes[i]);
            credentials[i] = credentialByCode[credCodes[i]];
        }
    }
}


