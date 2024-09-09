// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IWRON} from './IWRON.sol';
import {IWRONHelper} from './IWRONHelper.sol';

interface ILRON {
    /// @dev Throws if the provided array is empty.
    error EmptyArray();
    /// @dev Throws if start time > end time.
    error InvalidTimeRange(uint256 start, uint256 end);
    /// @dev Throws if the current lock position count exceeds the maximum lock position count.
    error ExceedMaxLockPosition(uint256 maxPosCount, uint256 currPosCount);
    /// @dev Throws if the lock until timestamp is greater than the maximum lock allowance.
    error ExceedMaxLockUntilAllowance(
        uint256 maxUntilAllowed,
        uint256 requestedUntil
    );
    /// @dev Throws if the lock duration exceeds the threshold.
    error ExceedLockDurationThreshold(uint256 threshold, uint256 duration);
    /// @dev Throws if the unlocked amounts are different from than the provided amount.
    error IncorrectUnlockAmount(uint256 acquired, uint256 provided);
    /// @dev Throws if the balance after deposit - balance before deposit is less than the provided amount.
    error IncorrectDepositAmount(uint256 acquired, uint256 provided);
    /// @dev Throws if the amount is zero.
    error NullAmountProvided();
    /// @dev Throws if the caller is not authorized for given action.
    error UnauthorizedAction(address caller);
    /// @dev Throws if the unlocker is invalid.
    error InvalidUnlocker(address unlocker);
    /// @dev Throws if the locker is invalid.
    error InvalidLocker(address locker);
    /// @dev Throws if the approver is invalid.
    error InvalidApprover(address approver);
    /// @dev Throws if the account is approving self.
    error SelfApproval();
    /// @dev Throws if the `until` timestamp is less than the current timestamp.
    error ForbiddenPastLockUntil(uint256 current, uint256 until);
    /// @dev Throws if insufficient value is sent.
    error InsufficientValue(address caller, uint256 sent, uint256 needed);
    /// @dev Throws if not enough locked balance to unlock.
    error InsufficientLockedBalance(
        address account,
        uint256 balance,
        uint256 needed
    );
    /// @dev Throws if the lockable allowance is insufficient.
    error InsufficientLockableAllowance(
        address locker,
        uint256 current,
        uint256 spent
    );

    /// @dev Lockable allowance data structure.
    struct LockableAllowance {
        // Amount approved for the locker to lock account fund.
        uint96 amount;
        // Maximum lock duration for locker to lock account fund.
        uint40 maxUntil;
    }

    /// @dev Locked position data structure.
    struct LockedPosition {
        // The locker address.
        address locker;
        // Amount of RON locked.
        uint96 amount;
        // The timestamp until the lock is invalidated.
        uint40 until;
    }

    /// @dev Emitted when the user deposits (W)RON into the contract.
    event Deposit(address indexed dst, uint256 wad);
    /// @dev Emitted when the user withdraws (W)RON from the contract.
    event Withdrawal(address indexed src, uint256 wad, bool wrapped);
    /// @dev Emitted when the locker's vault balance is updated.
    event VaultBalanceUpdated(
        address indexed locker,
        uint256 prevBalance,
        uint256 newBalance
    );
    /**
     * @dev Emitted when the user or locker unlocks RON from the locker's vault.
     *
     * @param by The address who unlocks the RON.
     * @param account The address of the owner of locked RONs.
     * @param locker The address of the locker.
     * @param unlockedAt The timestamp when the RON is unlocked.
     * @param actualLockedUntil The actual locked until timestamp.
     * @param amount The amount of RON unlocked.
     * @param remain The remaining locked RON amount at given lock position.
     * @param selfUnlock Whether the unlock is initiated by the owner.
     *
     */
    event Unlock(
        address indexed by,
        address indexed account,
        address indexed locker,
        uint256 unlockedAt,
        uint256 actualLockedUntil,
        uint256 amount,
        uint256 remain,
        bool selfUnlock
    );
    /// @dev Emitted when the user locks their RON.
    event Lock(
        address indexed account,
        address indexed locker,
        uint256 amount,
        uint256 lockedAt,
        uint256 lockedUntil
    );
    /// @dev Emitted when the user approves the lock.
    event LockApproval(
        address indexed account,
        address indexed locker,
        uint256 amount,
        uint256 maxUntil
    );

    /**
     * @notice The maximum duration threshold to lock RON amount.
     */
    function MAX_LOCK_DURATION() external view returns (uint256);

    /**
     * @notice The address of the WRON contract.
     */
    function WRON() external view returns (IWRON);

    /**
     * @notice Deposit (W)RON into the contract.
     *
     * @dev
     * - If `msg.value` = 0, contract will infer provided `amount` field as WRON amount.
     * - Emits a {Deposit} event.
     * - All WRON deposits are converted into RON.
     */
    function deposit(uint96 amount) external payable;

    /**
     * @notice Withdraw (W)RON from the contract.
     *
     * @dev
     *
     * Requirements:
     * - The caller is the owner.
     * - The provided `amount` is less than or equal to the balance.
     * - Emits a {Withdrawal} event.
     */
    function withdraw(uint96 amount, bool shouldWrap) external;

    /**
     * @notice Unlock and withdraw the specified amount of RON from the locker's vault.
     *
     * @dev
     *
     * Requirements:
     * - The caller is the owner.
     * - If `locker` is not provided, unlock and withdraw from all lockers.
     * - If `shouldWrap` is true, the withdrawn amount will be wrapped to WRON.
     * - The provided `amount` must be less than or equal to the unlockable amount.
     *
     * Emits a {Unlock} event.
     * Emits a {Withdrawal} event.
     */
    function unlockAndWithdraw(
        address locker,
        uint96 amount,
        bool shouldWrap
    ) external;

    /**
     * @notice Unlock the specified amount of RON from the locker's vault.
     *
     * @dev
     *
     * Requirements:
     * - The caller is the owner.
     * - If `locker` is not provided, unlock from all lockers.
     * - The provided `amount` is less than or equal to the unlockable amount.
     *
     * Emits a {Unlock} event.
     */
    function unlock(address locker, uint96 amount) external;

    /**
     * @notice Locker unlock the specified amount of RON from the locker's vault.
     *
     * @dev
     *
     * Requirements:
     * - The caller is the locker.
     * - The provided `amount` is less than or equal to the unlockable amount.
     *
     * NOTE: Locker can unlock before `until` timestamp passed.
     *
     * Emits a {Unlock} event.
     */
    function unlockFrom(address account, uint96 amount) external;

    /**
     * @notice Locker unlock the specified amount of RON from the locker's vault
     * with specified unlock range [`unlockedAtStart`,`unlockedAtEnd`].
     *
     * @dev
     *
     * Requirements:
     * - The caller is the locker.
     * - The provided `amount` is less than or equal to the unlockable amount.
     * - `unlockedAtStart` must be less than or equal to `unlockedAtEnd`.
     *
     * NOTE: Locker can unlock before `unlockedAtStart` or `unlockedAtEnd` timestamp passed.
     *
     * Emits a {Unlock} event.
     */
    function unlockFromWithinRange(
        address account,
        uint40 unlockedAtStart,
        uint40 unlockedAtEnd
    ) external;

    /**
     * @notice Spender lock the specified amount of RON for the account until the specified timestamp.
     *
     * @dev
     *
     * Requirements:
     * - The caller is the approved locker.
     * - The provided `amount` is less than or equal to the account's lockable balance.
     * - The provided `until` is less than or equal to the maximum allowed lock until timestamp.
     *
     * Emits a {Lock} event.
     */
    function lockFrom(
        address account,
        uint96 amount,
        uint40 until
    ) external returns (bool);

    /**
     * @notice Approve `locker` to lock `amount` of RON until `until` timestamp expired.
     *
     * @dev
     *
     * Requirements:
     * - The caller must be the owner.
     * - The locker must differ from the caller.
     * - The `until` timestamp must be greater than the current timestamp.
     *
     * `amount` = 0 means the approval is revoked.
     * `amount` = `type(uint96).max` means the approval is infinite.
     *
     * Emits a {LockApproval} event.
     */
    function approveLock(
        address locker,
        uint96 amount,
        uint40 until
    ) external returns (bool);

    /**
     * @notice Get lockable allowance data of the locker from the account.
     */
    function lockableAllowance(address account, address locker)
        external
        view
        returns (LockableAllowance memory);

    /**
     * @notice Get total of locked RON in the contract.
     */
    function totalLocked() external view returns (uint256);

    /**
     * @notice Get the locked balance of the account.
     */
    function lockedBalanceOf(address account)
        external
        view
        returns (uint256 lockedBalance);

    /**
     * @notice Get the total lockable RON and actively locked RON for given `account`.
     */
    function totalDeposited(address account)
        external
        view
        returns (uint256 deposited);

    /**
     * @notice Total of RON locked by given `locker`.
     */
    function vaultBalanceOf(address locker) external view returns (uint256);

    /**
     * @notice Get the total unlockable amount of RON that can be unlocked for current timestamp.
     * @dev
     * - If `locker` is not provided, the unlockable amount from all lockers will be included.
     */
    function unlockableAmount(address account, address locker)
        external
        view
        returns (uint256 unlockableByOwner, uint256 unlockableByLocker);

    /**
     * @notice Get the unlockable amount of RON that can be unlocked for given `account` and `locker` within the specified range.
     */
    function unlockableAmountWithinRange(
        address account,
        address locker,
        uint40 unlockedAtStart,
        uint40 unlockedAtEnd
    )
        external
        view
        returns (uint256 unlockableByOwner, uint256 unlockableByLocker);

    /**
     * @notice Get the count of locked positions.
     */
    function lockedPositionCount(address account)
        external
        view
        returns (uint256);

    /**
     * @notice Get the locked position at the specified index.
     */
    function lockedPositionAt(address account, uint256 index)
        external
        view
        returns (LockedPosition memory);
}
