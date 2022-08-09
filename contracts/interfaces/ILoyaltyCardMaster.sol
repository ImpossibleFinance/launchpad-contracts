// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/**
  @title Interface for the loyalty card use-case-specific functions 
         of the Impossible Finance LoyaltyCardMaster contract
  @author Impossible Finance
 */
interface ILoyaltyCardMaster {
    function mint(address to) external;

    function setMinter(address _minter) external;

    function setBurner(address _burner) external;

    function burn(uint256 tokenId) external;

    /// @notice Returns the IF user which the card was initially minted to, and effectively belongs to
    function originalOwnerToTokenId(address owner)
        external
        view
        returns (uint256);

    /**
      @notice Returns whether given loyalty card NFT is currently staked
      @param tokenId Token id of the loyalty card to be checked
      @dev Staked Loyalty Card NFTs have a whitelisted contract as their token owner - i.e., ownerOf(tokenId),
        and originalOwnerToTokenId() returns the IF user which the card effectively belongs to
     */
    function isStaked(uint256 tokenId) external view returns (bool);

    function totalPoints(uint256 tokenId) external;

    function currentPoints(uint256 tokenId) external;

    function addPoints(uint256 tokenId, uint256 points) external;

    function redeemPoints(uint256 tokenId, uint256 points) external;

    function isDestination(address dest) external;

    function addDestination(address destinationToAdd) external;

    function removeDestination(address destinationToRemove) external;

    function isOperator(address op) external;

    function addOperator(address operatorToAdd) external;

    function removeOperator(address operatorToRemove) external;
}
