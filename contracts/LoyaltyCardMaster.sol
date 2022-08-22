// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// TODO TBD we could also say that the owner should always automatically be allowed to mint and operate

/**
  @title Loyalty Points Cards NFTs for Impossible Finance users.
  @author Impossible Finance
  @notice Minting only allowed for specific minter
  @notice The cards are not transferrable except for whitelisted destinations.
  @notice The cards can be burned.
  @notice The cards acccumulate points which can be redeemed. Only whitelisted operators are allowed to do so.
  @notice This contract has no knowledge of particular benefits.
  @notice Benefits for redeeming loyalty points are handled entirely by operating contracts.
  @notice Benefits operators have to ensure that, when redeeming, they deduct correct point amounts and give the benefits accordingly.
 */

contract LoyaltyCardMaster is ERC721, Ownable {
    // --- MINTING / BURNING

    address public minter; // dedicated minting operator - TODO TBD ideally another contract so we can execute dedicated on-chain logic on mint
    address public burner; // dedicated burn operator - TODO TBD ideally another contract so we can execute dedicated on-chain logic on burn
    uint256 public mintCounter;
    uint256 public burnCounter; // may come in handy for statistics

    /// @dev ! --- We don't use tokenId 0 --- !

    /// Map IF user account to their loyalty card NFT tokenId
    /// @dev owner account => single owned tokenId. 
    /// @dev If value == 0, this user owns no loyalty card
    mapping(address => uint256) public originalOwnerToTokenId;

    event SetMinter(address minter);
    event SetBurner(address burner);

    modifier onlyMinter() {
        if (msg.sender != minter) revert NotAllowedToMint();
        _;
    }
    modifier onlyExistingToken(uint256 tokenId) {
        if (!_exists(tokenId)) revert TokenDoesntExist();
        _;
    }
    modifier onlyCardOwner(address account) {
        if (originalOwnerToTokenId[account] == 0) revert NoCardForUser();
        _;
    }

    error NotAllowedToMint();
    error NotAllowedToBurn();
    error TokenDoesntExist();
    error AlreadyOwnsCard();
    error NoCardForUser();
    error DestinationOwnsTokens();
    error CannotBurnStakedCard();
    error BatchRewardLengthsMismatch();

    // --- POINTS

    mapping(uint256 => uint256) public tokenIdToTotalPoints;
    mapping(uint256 => uint256) public tokenIdToCurrentPoints;

    event AddedPoints(uint256 tokenId, uint256 points, address operator);
    event RedeemedPoints(uint256 tokenId, uint256 points, address operator);

    error InsufficientPoints();

    // --- OPERATORS

    mapping(address => bool) whitelistedOperator;

    modifier onlyOperator() {
        if (!whitelistedOperator[msg.sender]) revert NotAllowedToOperate();
        _;
    }

    event AddedOperator(address operator);
    event RemovedOperator(address operator);

    error AlreadyWhitelistedOperator();
    error NotWhitelistedOperator();
    error NotAllowedToOperate();

    // --- TRANSFERS

    mapping(address => bool) whitelistedStakingDestination;

    event AddedDestination(address destination);
    event RemovedDestination(address destination);

    error AlreadyWhitelistedStakingDestination();
    error NotWhitelistedStakingDestination();
    error NotAllowedAsDestination();

    constructor(string memory name_, string memory symbol_)
        ERC721(name_, symbol_)
    {}

    // ================= MINTING / BURNING ================= //

    /**
      @notice Mint a new card to given account
      @param to The account to mint to
     */
    function mint(address to) external onlyMinter {
        if (originalOwnerToTokenId[to] != 0) revert AlreadyOwnsCard();
        uint256 tokenId = ++mintCounter; /// @dev first tokenId will be 1;
        _mint(to, tokenId);
        originalOwnerToTokenId[to] = tokenId;
    }

    function setMinter(address _minter) external onlyOwner {
        minter = _minter;
        emit SetMinter(minter);
    }

    function setBurner(address _burner) external onlyOwner {
        burner = _burner;
        emit SetBurner(burner);
    }

    /**
        @notice Burns a token with given tokenId
        @param tokenId The id of the token to be burned.

        @dev Following OZ's pattern: _burn() can only be called by the token owner or an approved party.
            Additionally, we say that the approved party must be our burner.
     */
    function burn(uint256 tokenId) external onlyExistingToken(tokenId) {
        /// @dev only allow burning while card is owned by original owner (not staked)
        if (whitelistedStakingDestination[ownerOf(tokenId)]) revert CannotBurnStakedCard();
        address spender = _msgSender();
        address owner = ERC721.ownerOf(tokenId);
        bool isOwner = spender == owner;
        bool isApprovedBurner;
        if (spender == burner) {
            isApprovedBurner = (isApprovedForAll(owner, spender) ||
                getApproved(tokenId) == spender);
        }
        if (!isOwner && !isApprovedBurner) revert NotAllowedToBurn();
        burnCounter++;
        _burn(tokenId);
        originalOwnerToTokenId[owner] = 0;
    }

    // ======================= POINTS ====================== //
    //                   (tokenID-centric)                   //

    /**
      @notice Retrieve total number of points this card has accumulated historically
      @param tokenId The tokenId of the card
     */
    function totalPointsCard(uint256 tokenId) external view returns (uint256) {
        return _totalPointsCard(tokenId);
    }

    /**
      @notice Retrieve current number of points this card has
      @param tokenId The tokenId of the card
     */
    function currentPointsCard(uint256 tokenId) external view returns (uint256) {
        return _currentPointsCard(tokenId);
    }

    /**
      @notice Add loyalty points to a given card
      @param tokenId The loyalty card to add points to
      @param points Number of points to add
     */
    function addPointsCard(uint256 tokenId, uint256 points) external {
        _addPointsCard(tokenId, points);
    }

    /**
      @notice Redeem loyalty points from a card
      @param tokenId The loyalty card to redeem points from
      @param points Number of points to redeem
     */
    function redeemPointsCard(uint256 tokenId, uint256 points) external {
        _redeemPointsCard(tokenId, points);
    }

    // ======================= POINTS ====================== //
    //                   (account-centric)                   //

    /**
      @notice Retrieve total number of points a user has accumulated historically on their current loyalty card
      @param account The IF user account
     */
    function totalPointsAccount(address account)
        external
        view
        onlyCardOwner(account)
        returns (uint256)
    {
        return _totalPointsCard(originalOwnerToTokenId[account]);
    }

    /**
      @notice Retrieve current number of points a user has on their loyalty card
      @param account The IF user account
     */
    function currentPointsAccount(address account)
        external
        view
        onlyCardOwner(account)
        returns (uint256)
    {
        return _currentPointsCard(originalOwnerToTokenId[account]);
    }

     /**
      @notice Add loyalty points to a given IF user account
      @param account The IF user account
      @param points Number of points to add
     */
    function addPointsAccount(address account, uint256 points)
        external 
    {
        _addPointsCard(originalOwnerToTokenId[account], points);
    }

    /**
      @notice Redeem loyalty points from an IF user account
      @param account The IF user account
      @param points Number of points to redeem
     */
    function redeemPointsAccount(address account, uint256 points)
        external
    {
        _redeemPointsCard(originalOwnerToTokenId[account], points);
    }

    /**
        @notice Add the same amount of points to multiple IF user accounts in a batch
        @param accounts IF user account addresses batch, as an array 
        @param pointsAmount The amount of points to add
     */
    function addPointsBatchAccSingleValue(address[] calldata accounts, uint256 pointsAmount) 
        external 
    {
        for (uint256 i = 0; i < accounts.length; i++) {
            _addPointsCard(originalOwnerToTokenId[accounts[i]], pointsAmount);
        }
    }

    /**
        @notice Add different amounts of points to multiple IF user accounts in a batch
        @param accounts IF user account addresses batch, as an array 
        @param pointsAmounts The amounts of points to add, as an array
     */
    function addPointsBatchAccMultiValues(address[] calldata accounts, uint256[] calldata pointsAmounts) 
        external 
    {
        if (accounts.length != pointsAmounts.length) revert BatchRewardLengthsMismatch();
        for (uint256 i = 0; i < accounts.length; i++) {
            _addPointsCard(originalOwnerToTokenId[accounts[i]], pointsAmounts[i]);
        }
    }

    // --------------------- POINTS ----------------------- // 
    //                     (internal)                       //

    function _totalPointsCard(uint256 tokenId)
        internal
        view
        onlyExistingToken(tokenId)
        returns (uint256)
    {
        return tokenIdToTotalPoints[tokenId];
    }

    function _currentPointsCard(uint256 tokenId)
        internal
        view
        onlyExistingToken(tokenId)
        returns (uint256)
    {
        return tokenIdToCurrentPoints[tokenId];
    }

    function _addPointsCard(uint256 tokenId, uint256 points)
        internal 
        onlyOperator 
        onlyExistingToken(tokenId)
    {
        tokenIdToTotalPoints[tokenId] += points;
        tokenIdToCurrentPoints[tokenId] += points;
        emit AddedPoints(tokenId, points, msg.sender);
    }

    function _redeemPointsCard(uint256 tokenId, uint256 points)
        internal
        onlyOperator
        onlyExistingToken(tokenId)
    {
        if (points > tokenIdToCurrentPoints[tokenId])
            revert InsufficientPoints();
        tokenIdToCurrentPoints[tokenId] -= points;
        emit RedeemedPoints(tokenId, points, msg.sender);
    }

    // ======================= TRANSFERS ================== //

    /**
      @dev Override the ERC721 default in order to block transfers to non-whitelisted destinations
     */
    function _transfer(
        address from,
        address to,
        uint256 tokenId
    ) internal override {
        // this only goes through if token is being sent to a WL destination
        // or a WL destination (related to Impossible Finance) transfers it (back)
        if (!whitelistedStakingDestination[to] && !whitelistedStakingDestination[from])
            revert NotAllowedAsDestination();
        super._transfer(from, to, tokenId);
    }

    /**
      @notice Returns whether given address may be receiver of tokens via transfer
      @param dest Address to check
     */
    function isDestination(address dest) external view returns (bool) {
        return whitelistedStakingDestination[dest];
    }

    /**
      @notice Returns whether given loyalty card NFT is currently staked
      @param tokenId Token id of the loyalty card to be checked
      @dev Staked Loyalty Card NFTs have a whitelisted contract as their token owner - i.e., ownerOf(tokenId),
        and originalOwnerToTokenId() returns the IF user which the card effectively belongs to
     */
    function isStaked(uint256 tokenId) external view returns (bool) {
        return whitelistedStakingDestination[ownerOf(tokenId)];
    }

    /**
      @notice Adds a destination to the whitelisted destinations
    */
    function addDestination(address destinationToAdd) external onlyOwner {
        whitelistedStakingDestination[destinationToAdd] = true;
        emit AddedDestination(destinationToAdd);
    }

    /**
      @notice Removes a destination from the whitelisted destinations
     */
    function removeDestination(address destinationToRemove) external onlyOwner {
        if (balanceOf(destinationToRemove) > 0) revert DestinationOwnsTokens();
        whitelistedStakingDestination[destinationToRemove] = false;
        emit RemovedDestination(destinationToRemove);
    }

    // ======================= OPERATORS ================== //

    /**
      @notice Returns whether given address may operate
      @param op Address to check
     */
    function isOperator(address op) external view returns (bool) {
        return whitelistedOperator[op];
    }

    /**
      @notice Adds an operator to the whitelisted operators
     */
    function addOperator(address operatorToAdd) external onlyOwner {
        whitelistedOperator[operatorToAdd] = true;
        emit AddedOperator(operatorToAdd);
    }

    /**
      @notice Removes an operator from the whitelisted operators
     */
    function removeOperator(address operatorToRemove) external onlyOwner {
        whitelistedOperator[operatorToRemove] = false;
        emit RemovedOperator(operatorToRemove);
    }
}
