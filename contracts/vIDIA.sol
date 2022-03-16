// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import 'hardhat/console.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/access/AccessControlEnumerable.sol';
import '../library/IFTokenStandard.sol';
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract vIDIA is AccessControlEnumerable, IFTokenStandard {
    using SafeERC20 for ERC20;

    // delay for unvesting token
    uint24 unstakingDelay;
    // constant penalty for early unvesting
    uint256 penalty;
    // if token is enabled for staking
    bool enabled;
    uint256 accumulatedPenalty;
    uint256 totalStakedAmount;
    uint256 totalUnstakedAmount;
    uint256 totalStakers;
    uint256 rewardSum; // (1/T1 + 1/T2 + 1/T3)
    address public tokenAddress;
    address admin;

    struct UserInfo {
        uint256 stakedAmount;
        uint256 unstakeAt;
        uint256 unstakedAmount;
    }

    bytes32 public constant PENALTY_SETTER_ROLE =
        keccak256('PENALTY_SETTER_ROLE');

    bytes32 public constant DELAY_SETTER_ROLE = keccak256('DELAY_SETTER_ROLE');

    bytes32 public constant WHITELIST_SETTER_ROLE =
        keccak256('WHITELIST_SETTER_ROLE');

    EnumerableSet.AddressSet private whitelistAddresses;

    // user info mapping (user addr => token addr => user info)
    mapping(address => mapping(address => UserInfo)) public userInfo;

    // Events

    event Stake(address _from, uint256 amount);

    event Unstake(address _from, uint256 amount);

    event ImmediateUnstake(address _from, uint256 amount);

    event SetWhitelistSetter(address whitelistSetter);

    event SetWhitelist(bytes32 whitelistRootHash);

    event Claim(address _from);

    event ImmediateClaim(address _from);

    event ClaimReward(address _from, uint256 amount);

    constructor(
        string memory _name,
        string memory _symbol,
        address _admin,
        address _tokenAddress
    ) AccessControlEnumerable() IFTokenStandard(_name, _symbol, _admin, _tokenAddress) {
        _setupRole(PENALTY_SETTER_ROLE, _msgSender();
        _setupRole(DELAY_SETTER_ROLE, _msgSender());
        _setupRole(WHITELIST_SETTER_ROLE, _msgSender());
        tokenAddress = _tokenAddress;
        admin = _admin;
    }

    function stake(uint256 amount) public {
        emit Stake(_msgSender(), amount);
    }

    function unstake(uint256 amount) public {
        emit Unstake(_msgSender(), amount);
    }

    function immediateUnstake(uint256 amount) public {
        emit ImmediateUnstake(_msgSender(), amount);
    }

    function claim() public {
        emit Claim(_msgSender());
    }

    function immediateClaim() public {
        emit ImmediateClaim(_msgSender());
    }

    // claim reward and reset user's reward sum
    function claimReward(address token) public {
        require(
            tokenConfigurations[token].enabled,
            'Invalid token for claiming reward'
        );
        uint256 reward = calculateUserReward(token);
        require(reward <= 0, 'No reward to claim');
        // reset user's rewards sum
        userInfo[msg.sender][token].lastRewardSum = tokenStats[token].rewardSum;
        // transfer reward to user
        ERC20 claimedTokens = ERC20(token);
        claimedTokens.safeTransfer(_msgSender(), reward);

        emit ClaimReward(_msgSender(), reward, token);
    }

    function setPenalty(uint256 newPenalty) external {
        require(
            hasRole(PENALTY_SETTER_ROLE, _msgSender()),
            'Must have penalty setter role'
        );
        penalty = newPenalty;
    }

    function setUnvestingDelay(uint24 newDelay) external {
        require(
            hasRole(DELAY_SETTER_ROLE, _msgSender()),
            'Must have delay setter role'
        );
        unvestingDelay = newDelay;
    }

    /** 
     @notice Adds an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to add to whitelist
     @return boolean. True = account was added, False = account already exists in set
     */
    function addToWhitelist(address account) public returns (bool) {
        require(hasRole(WHITELIST_SETTER_ROLE, _msgSender()), 'Must have whitelist setter role');
        isWhitelistedAddr[account] = true;
        return EnumerableSet.add(whitelistAddresses, account);
    }

    /** 
     @notice Removes an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to remove from whitelist
     @return boolean. True = account was removed, False = account doesnt exist in set
     */
    function removeFromWhitelist(address account) public returns (bool) {
        require(hasRole(WHITELIST_SETTER_ROLE, _msgSender()), 'Must have whitelist setter role');
        isWhitelistedAddr[account] = false;
        return EnumerableSet.remove(whitelistAddresses, account);
    }

    /** 
     @notice Getter for all transfer whitelisted addresses
     @return Array of all transfer whitelisted addresses
     */
    function getAllWhitelistedAddrs() public view returns (address[]) {
        return EnumerableSet.values(whitelistAddresses);
    }

    /** 
     @notice Standard ERC20 transfer but only to/fro whitelisted addresses
     @dev purpose is to enable transfers to and fro launchpad contract only
     @param to address to send tokens to
     @param amount transfer amount
     @return boolean representing if transfer was successful
     */
    function transfer(address to, uint256 amount) public override returns (bool) {
        require(EnumerableSet.contains(whitelistAddresses, to) || EnumerableSet.contains(whitelistAddresses, _msgSender()), 'Origin and dest address not in whitelist');
        return ERC20.transfer(to, amount);
    }

    /** 
     @notice Standard ERC20 transferFrom but only to/fro whitelisted addresses
     @dev purpose is to enable transfers to and fro launchpad contract only
     @param to address the tokens are sent from 
     @param to address to send tokens to
     @param amount transfer amount
     @return boolean representing if transfer was successful
     */
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        require(EnumerableSet.contains(whitelistAddresses, from) || EnumerableSet.contains(whitelistAddresses, to), 'Origin and dest address not in whitelist');
        return ERC20.transferFrom(from, to, amount);
    }

    /** 
     @notice Adds an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to add to whitelist
     @return boolean. True = account was added, False = account already exists in set
     */
    function addToWhitelist(address account) public returns (bool) {
        require(hasRole(WHITELIST_SETTER_ROLE, _msgSender()), 'Must have whitelist setter role');
        return EnumerableSet.add(whitelistAddresses, account);
    }

    /** 
     @notice Removes an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to remove from whitelist
     @return boolean. True = account was removed, False = account doesnt exist in set
     */
    function removeFromWhitelist(address account) public returns (bool) {
        require(hasRole(WHITELIST_SETTER_ROLE, _msgSender()), 'Must have whitelist setter role');
        return EnumerableSet.remove(whitelistAddresses, account);
    }

    /** 
     @notice Getter for all transfer whitelisted addresses
     @return Array of all transfer whitelisted addresses
     */
    function getAllWhitelistedAddrs() public view returns (address[] memory) {
        return EnumerableSet.values(whitelistAddresses);
    }

    /** 
     @notice Standard ERC20 transfer but only to/fro whitelisted addresses
     @dev purpose is to enable transfers to and fro launchpad contract only
     @param to address to send tokens to
     @param amount transfer amount
     @return boolean representing if transfer was successful
     */
    function transfer(address to, uint256 amount) public override returns (bool) {
        require(EnumerableSet.contains(whitelistAddresses, to) || EnumerableSet.contains(whitelistAddresses, _msgSender()), 'Origin and dest address not in whitelist');
        return ERC20.transfer(to, amount);
    }

    /** 
     @notice Standard ERC20 transferFrom but only to/fro whitelisted addresses
     @dev purpose is to enable transfers to and fro launchpad contract only
     @param from address the tokens are sent from 
     @param to address to send tokens to
     @param amount transfer amount
     @return boolean representing if transfer was successful
     */
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        require(EnumerableSet.contains(whitelistAddresses, from) || EnumerableSet.contains(whitelistAddresses, to), 'Origin and dest address not in whitelist');
        return ERC20.transferFrom(from, to, amount);
    }

    //// EIP2771 meta transactions

    function _msgSender()
        internal
        view
        override(IFTokenStandard, Context)
        returns (address)
    {
        return ERC2771ContextUpdateable._msgSender();
    }

    function _msgData()
        internal
        view
        override(IFTokenStandard, Context)
        returns (bytes calldata)
    {
        return ERC2771ContextUpdateable._msgData();
    }

    //// EIP1363 payable token

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(AccessControlEnumerable, IFTokenStandard)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
