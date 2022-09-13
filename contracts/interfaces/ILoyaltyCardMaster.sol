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

    function originalOwnerToTokenId(address owner)
        external
        view
        returns (uint256);

    function isStaked(uint256 tokenId) external view returns (bool);

    function totalPointsCard(uint256 tokenId) external;

    function currentPointsCard(uint256 tokenId) external;

    function addPointsCard(uint256 tokenId, uint256 points) external;

    function redeemPointsCard(uint256 tokenId, uint256 points) external;

    function totalPointsAccount(address account) external;

    function currentPointsAccount(address account) external;

    function addPointsAccount(address account, uint256 points) external;

    function addPointsBatchAccSingleValue(
        address[] calldata accounts,
        uint256 points,
        uint256[] calldata multipliers
    ) external;

    function addPointsBatchAccMultiValues(
        address[] calldata accounts,
        uint256[] calldata pointsAmounts
    ) external;

    function redeemPointsAccount(address account, uint256 points) external;

    function isDestination(address dest) external;

    function addDestination(address destinationToAdd) external;

    function removeDestination(address destinationToRemove) external;

    function isOperator(address op) external;

    function addOperator(address operatorToAdd) external;

    function removeOperator(address operatorToRemove) external;
}
