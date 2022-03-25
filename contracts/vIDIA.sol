// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import 'hardhat/console.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/access/AccessControlEnumerable.sol';
import '../library/IFTokenStandard.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';

contract vIDIA is AccessControlEnumerable, IFTokenStandard {
    using SafeERC20 for ERC20;

    uint256 private constant FACTOR = 10**30;
    uint256 private constant ONE_HUNDRED = 10000; // one hundred in basis points

    // delay for unstaking token
    uint256 public unstakingDelay = 86400 * 14; // 2 weeks in seconds

    // Fees for different actions. All fees denoted in basis points
    uint256 public skipDelayFee = 2000; // initialzed at 20%
    uint256 public cancelUnstakeFee = 200; // initialized at 2%

    uint256 public accumulatedFee;
    uint256 public totalStakedAmount;
    uint256 public rewardSum; // (1/T1 + 1/T2 + 1/T3)
    address public tokenAddress;

    address admin;

    struct UserInfo {
        uint256 stakedAmount;
        uint256 unstakeAt;
        uint256 unstakedAmount;
        uint256 lastRewardSum;
    }

    bytes32 public constant FEE_SETTER_ROLE = keccak256('FEE_SETTER_ROLE');

    bytes32 public constant DELAY_SETTER_ROLE = keccak256('DELAY_SETTER_ROLE');

    bytes32 public constant WHITELIST_SETTER_ROLE =
        keccak256('WHITELIST_SETTER_ROLE');

    EnumerableSet.AddressSet private whitelistAddresses;

    // user info mapping (user addr => token addr => user info)
    mapping(address => UserInfo) public userInfo;

    // Events

    event Stake(address _from, uint256 amount);

    event Unstake(address _from, uint256 amount);

    event ClaimStaked(address _from, uint256 fee, uint256 withdrawAmount);

    event SetWhitelistSetter(address whitelistSetter);

    event SetWhitelist(bytes32 whitelistRootHash);

    event ClaimUnstaked(address _from, uint256 withdrawAmount);

    event ClaimPendingUnstake(
        address _from,
        uint256 fee,
        uint256 withdrawAmount
    );

    event CancelPendingUnstake(
        address _from,
        uint256 fee,
        uint256 stakedAmount
    );

    event ClaimReward(address _from, uint256 amount);

    constructor(
        string memory _name,
        string memory _symbol,
        address _admin,
        address _tokenAddress
    ) AccessControlEnumerable() IFTokenStandard(_name, _symbol, _admin) {
        _setupRole(FEE_SETTER_ROLE, _admin);
        _setupRole(DELAY_SETTER_ROLE, _admin);
        _setupRole(WHITELIST_SETTER_ROLE, _admin);
        tokenAddress = _tokenAddress;
        admin = _admin;
    }

    function stake(uint256 amount) public {
        claimReward();
        totalStakedAmount += amount;
        userInfo[_msgSender()].stakedAmount += amount;
        _mint(_msgSender(), amount);
        ERC20 stakedTokens = ERC20(tokenAddress);
        stakedTokens.safeTransferFrom(_msgSender(), address(this), amount);
        emit Stake(_msgSender(), amount);
    }

    /** 
     @notice Function for a user unstake tokens and put them in unstaking queue
     @param amount the amount of tokens to unstake from staked tokens
     */
    function unstake(uint256 amount) public {
        require(
            userInfo[_msgSender()].unstakedAmount == 0 ||
                userInfo[_msgSender()].unstakeAt == 0,
            'User has pending tokens unstaking'
        );
        claimReward();
        totalStakedAmount -= amount;
        userInfo[_msgSender()].stakedAmount -= amount;
        //start unvesting period
        userInfo[_msgSender()].unstakeAt = block.timestamp + unstakingDelay;

        userInfo[_msgSender()].unstakedAmount = amount;
        burn(userInfo[_msgSender()].unstakedAmount);
        emit Unstake(_msgSender(), amount);
    }

    /** 
     @notice Function for a user to retrieve underlying tokens after waiting for the unstake delay
     @notice *no* fees required
     @notice For tokens in the unstaking queue, use instantUnstakePending()
     */
    function claimUnstaked() public {
        //require curr time more than unstaking delay
        require(
            block.timestamp > userInfo[_msgSender()].unstakeAt,
            'Tokens have not finished vesting'
        );
        uint256 withdrawAmount = userInfo[_msgSender()].unstakedAmount;
        ERC20(tokenAddress).safeTransfer(
            _msgSender(),
            withdrawAmount
        );
        userInfo[_msgSender()].unstakedAmount = 0;
        userInfo[_msgSender()].unstakeAt = 0;

        emit ClaimUnstaked(_msgSender(), withdrawAmount);
    }

    /** 
     @notice Function for a user to pay fee and receive underlying tokens *NOT* in the unstaking queue
     @notice fees required
     @notice For tokens in the unstaking queue, use claimPendingUnstake()
     @param amount the amount of tokens to instantly withdraw from staked tokens
     */
    function claimStaked(uint256 amount) public {
        claimReward();

        uint256 fee = (amount * skipDelayFee) / ONE_HUNDRED;
        uint256 withdrawAmount = amount - fee;
        uint256 divisor = totalStakedAmount - userInfo[_msgSender()].stakedAmount;

        if (divisor != 0) {
            // mul by FACTOR of 10**30 to reduce truncation
            rewardSum += (fee * FACTOR) / divisor;
            userInfo[_msgSender()].lastRewardSum = rewardSum;
        }
        
        totalStakedAmount -= amount;
        userInfo[_msgSender()].stakedAmount -= amount;
        accumulatedFee += fee;

        burn(amount);
        ERC20(tokenAddress).safeTransfer(_msgSender(), withdrawAmount);
        emit ClaimStaked(_msgSender(), fee, withdrawAmount);
    }

    /** 
     @notice Function for a user to retrieve underlying tokens associated with vidia in the unstaking queue
     @notice fees required
     @dev Requires user to have tokens in the unstake queue which cannot be claimed now
     @param amount the amount of tokens to instantly withdraw from unstake queue
     */
    function claimPendingUnstake(uint256 amount) public {
        require(
            userInfo[_msgSender()].unstakeAt > block.timestamp,
            'Can unstake without paying fee'
        );
        claimReward();

        uint256 fee = (amount * skipDelayFee) / ONE_HUNDRED;
        uint256 withdrawAmount = amount - fee;
        uint256 divisor = totalStakedAmount - userInfo[_msgSender()].stakedAmount;

        if (divisor != 0) {
            // mul by FACTOR of 10**30 to reduce truncation
            rewardSum += (fee * FACTOR) / divisor;
            userInfo[_msgSender()].lastRewardSum = rewardSum;
        }
        accumulatedFee += fee;

        userInfo[_msgSender()].unstakedAmount -= amount;
        if (userInfo[_msgSender()].unstakedAmount == 0) {
            userInfo[_msgSender()].unstakeAt = 0;
        }
        ERC20(tokenAddress).safeTransfer(_msgSender(), withdrawAmount);
        emit ClaimPendingUnstake(_msgSender(), fee, withdrawAmount);
    }

    /** 
     @notice Function for a user to cancel unstaking process for vidia
     @notice fees required
     @dev Requires user to have tokens in the unstake queue which cannot be claimed now
     @param amount the amount of tokens to cancel unstaking process for
     */
    function cancelPendingUnstake(uint256 amount) public {
        require(
            userInfo[_msgSender()].unstakeAt > block.timestamp,
            'Can restake without paying fee'
        );
        claimReward();

        uint256 fee = (amount * cancelUnstakeFee) / ONE_HUNDRED;
        uint256 stakeAmount = amount - fee;
        uint256 divisor = totalStakedAmount - userInfo[_msgSender()].stakedAmount;

        if (divisor != 0) {
            // mul by FACTOR of 10**30 to reduce truncation
            rewardSum += (fee * FACTOR) / divisor;
            userInfo[_msgSender()].lastRewardSum = rewardSum;
        }
        accumulatedFee += fee;

        userInfo[_msgSender()].unstakedAmount -= amount;
        if (userInfo[_msgSender()].unstakedAmount == 0) {
            userInfo[_msgSender()].unstakeAt = 0;
        }

        userInfo[_msgSender()].stakedAmount += stakeAmount;
        totalStakedAmount += stakeAmount;
        _mint(_msgSender(), stakeAmount);
        emit CancelPendingUnstake(_msgSender(), fee, stakeAmount);
    }

    // claim reward and reset user's reward sum
    function claimReward() public {
        uint256 reward = calculateUserReward();
        // reset user's rewards sum
        userInfo[_msgSender()].lastRewardSum = rewardSum;
        // transfer reward to user
        ERC20 claimedTokens = ERC20(tokenAddress);
        claimedTokens.safeTransfer(_msgSender(), reward);
        emit ClaimReward(_msgSender(), reward);
    }

    /** 
     @notice Update fee levied for instantly unstaking. Fee is in basis points
     @dev Requires fee setter role and fee must be below 10000 basis pts
     @param newFee the new fee
     */
    function updateSkipDelayFee(uint256 newFee) external {
        require(
            hasRole(FEE_SETTER_ROLE, _msgSender()),
            'Must have fee setter role'
        );
        require(newFee <= 10000, 'Fee must be less than 100%');
        skipDelayFee = newFee;
    }

    /** 
     @notice Update fee levied for cancelling unstaking. Fee is in basis points
     @dev Requires fee setter role and fee must be below 10000 basis pts
     @param newFee the new fee
     */
    function updateCancelUnstakeFee(uint256 newFee) external {
        require(
            hasRole(FEE_SETTER_ROLE, _msgSender()),
            'Must have fee setter role'
        );
        require(newFee <= 10000, 'Fee must be less than 100%');
        cancelUnstakeFee = newFee;
    }

    /** 
     @notice Update wait period required for fee-free unvesting. initialized at 2 weeks
     @dev Requires delay setter role and existing wait times will not change
     @param newDelay the new delay
     */
    function updateUnstakingDelay(uint24 newDelay) external {
        require(
            hasRole(DELAY_SETTER_ROLE, _msgSender()),
            'Must have delay setter role'
        );
        unstakingDelay = newDelay;
    }

    /** 
     @notice Calculates user reward
     @dev formula: amount * (global_reward_sum - user_reward_sum) / 10**30
     @dev we perform div 10**30 as rewardsum is inflated by 10**30 to reduce truncation
     @return uint256 amount of underlying tokens the user has earned from fees
     */
    function calculateUserReward() public view returns (uint256) {
        return
            (userInfo[_msgSender()].stakedAmount *
                (rewardSum - userInfo[_msgSender()].lastRewardSum)) / FACTOR;
    }

    /** 
     @notice Adds an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to add to whitelist
     @return boolean. True = account was added, False = account already exists in set
     */
    function addToWhitelist(address account) public returns (bool) {
        require(
            hasRole(WHITELIST_SETTER_ROLE, _msgSender()),
            'Must have whitelist setter role'
        );
        return EnumerableSet.add(whitelistAddresses, account);
    }

    /** 
     @notice Removes an address to the transfer whitelist
     @dev requires whitelist setter role
     @param account is the address to remove from whitelist
     @return boolean. True = account was removed, False = account doesnt exist in set
     */
    function removeFromWhitelist(address account) public returns (bool) {
        require(
            hasRole(WHITELIST_SETTER_ROLE, _msgSender()),
            'Must have whitelist setter role'
        );
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
    function transfer(address to, uint256 amount)
        public
        override
        returns (bool)
    {
        require(
            EnumerableSet.contains(whitelistAddresses, to) ||
                EnumerableSet.contains(whitelistAddresses, _msgSender()),
            'Origin and dest address not in whitelist'
        );
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
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public override returns (bool) {
        require(
            EnumerableSet.contains(whitelistAddresses, from) ||
                EnumerableSet.contains(whitelistAddresses, to),
            'Origin and dest address not in whitelist'
        );
        return ERC20.transferFrom(from, to, amount);
    }

    /** 
     @notice msg.sender for EIP2771 meta transactions. Parses out original msg.sender when transaction is sent by relayer
     @return address of msg.sender
     */
    function _msgSender()
        internal
        view
        override(IFTokenStandard, Context)
        returns (address)
    {
        return ERC2771ContextUpdateable._msgSender();
    }

    /** 
     @notice msg.data for EIP2771 meta transactions. Parses out original msg.data when transaction is sent by relayer
     @return bytes of msg.data
     */
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
