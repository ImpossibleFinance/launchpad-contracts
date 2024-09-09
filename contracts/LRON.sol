// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';
import { ILRON } from "./interfaces/ILRON.sol";
import { IWRON } from "./interfaces/IWRON.sol";
import { IWRONHelper } from "./interfaces/IWRONHelper.sol";
import { RONTransferHelper } from "./RONTransferHelper.sol";

contract LRON is ReentrancyGuard, ERC20, ILRON {
  /// @inheritdoc ILRON
  uint256 public constant MAX_LOCK_DURATION = 365 days;

  /// @dev Maximum active lock position for given account.
  uint256 private constant MAX_LOCK_POSITION = 1000;
  /// @dev Maximum amount constant.
  uint96 private constant MAX_AMOUNT = type(uint96).max;
  /// @dev Maximum lock timestamp constant.
  uint40 private constant MAX_TIMESTAMP = type(uint40).max;

  /// @inheritdoc ILRON
  IWRON public immutable WRON;

  /// @dev Total RON locked for each locker.
  mapping(address => uint256 ) private _vaultBalance;
  /// @dev Locked position record for each account.
  mapping(address => LockedPosition[]) private _lockedPositions;
  /// @dev Lockable allowance record for each account and locker.
  mapping(address => mapping(address => LockableAllowance)) private _lockableAllowance;

  modifier nonNullAmount(uint96 amount) {
    _requireNonNullAmount(amount);
    _;
  }

  modifier validDestination(address to) {
    _requireValidDestination(to);
    _;
  }

  receive() external payable {
    address sender = _msgSender();
    // Only `WRON` can deposit RON to this contract.
    if (sender != address(WRON)) revert UnauthorizedAction(sender);
  }

  constructor(address wron) ERC20("Lockable RON", "LRON") {
    WRON = IWRON(wron);
  }

  /**
   * @inheritdoc ILRON
   */
  function deposit(uint96 amount) external payable nonNullAmount(amount) nonReentrant {
    address account = _msgSender();

    if (msg.value == 0) {
      IWRON wron = WRON;

      uint256 selfBalanceBefore = address(this).balance;

      wron.transferFrom({ src: account, dst: address(this), wad: amount });
      wron.withdraw(amount);

      // Safety check if the amount is correctly deposited.
      uint256 acquired = address(this).balance - selfBalanceBefore;
      if (acquired != amount) revert IncorrectDepositAmount({ acquired: acquired, provided: amount });
    } else if (msg.value != amount) {
      revert InsufficientValue({ caller: account, sent: msg.value, needed: amount });
    }

    _mint(account, amount);

    emit Deposit({ dst: account, wad: amount });
  }

  /**
   * @inheritdoc ILRON
   */
  function approveLock(address locker, uint96 amount, uint40 maxUntil) external returns (bool) {
    _approveLock({ account: _msgSender(), locker: locker, amount: amount, maxUntil: maxUntil, shouldEmitEvent: true });

    return true;
  }

  /**
   * @inheritdoc ILRON
   */
  function lockFrom(address account, uint96 amount, uint40 until) external nonNullAmount(amount) returns (bool) {
    address locker = _msgSender();

    _spendLockAllowance({ account: account, locker: locker, amount: amount, requestedUntil: until });
    _lock({ account: account, locker: locker, amount: amount, until: until });

    return true;
  }

  /**
   * @inheritdoc ILRON
   */
  function unlockAndWithdraw(address locker, uint96 amount, bool shouldWrap) external {
    unlock(locker, amount);
    withdraw(amount, shouldWrap);
  }

  /**
   * @inheritdoc ILRON
   */
  function unlock(address locker, uint96 amount) public nonNullAmount(amount) {
    _unlock({ account: _msgSender(), locker: locker, amount: amount, tFrom: 0, tTo: uint40(block.timestamp) });
  }

  /**
   * @inheritdoc ILRON
   */
  function unlockFrom(address account, uint96 amount) external nonNullAmount(amount) {
    _unlock({ account: account, locker: _msgSender(), amount: amount, tFrom: 0, tTo: MAX_TIMESTAMP });
  }

  /**
   * @inheritdoc ILRON
   */
  function unlockFromWithinRange(address account, uint40 unlockedAtStart, uint40 unlockedAtEnd) external {
    address locker = _msgSender();

    (, uint256 unlockableWithinRange) = unlockableAmountWithinRange({
      account: account,
      locker: locker,
      unlockedAtStart: unlockedAtStart,
      unlockedAtEnd: unlockedAtEnd
    });

    _requireNonNullAmount(uint96(unlockableWithinRange));

    _unlock({
      account: account,
      locker: locker,
      amount: uint96(unlockableWithinRange),
      tFrom: unlockedAtStart,
      tTo: unlockedAtEnd
    });
  }

  /**
   * @inheritdoc ILRON
   */
  function withdraw(uint96 amount, bool shouldWrap) public nonNullAmount(amount) nonReentrant {
    address account = _msgSender();

    _burn(account, amount);

    if (shouldWrap) {
      IWRON wron = WRON;
      wron.deposit{ value: amount }();
      wron.transfer({ to: account, value: amount });
    } else {
      RONTransferHelper.safeTransfer({ _to: payable(account), _value: amount });
    }

    emit Withdrawal({ src: account, wad: amount, wrapped: shouldWrap });
  }

  /**
   * @dev See {IERC20-transfer}.
   */
  function transfer(address to, uint256 value) public virtual override validDestination(to) returns (bool) {
    return super.transfer(to, value);
  }

  /**
   * @dev See {IERC20-transferFrom}.
   */
  function transferFrom(address from, address to, uint256 value)
    public
    virtual
    override
    validDestination(to)
    returns (bool)
  {
    return super.transferFrom(from, to, value);
  }

  /**
   * @inheritdoc ILRON
   */
  function unlockableAmount(address account, address locker)
    external
    view
    returns (uint256 unlockableByOwner, uint256 unlockableByLocker)
  {
    return unlockableAmountWithinRange({
      account: account,
      locker: locker,
      unlockedAtStart: 0,
      unlockedAtEnd: MAX_TIMESTAMP
    });
  }

  /**
   * @inheritdoc ILRON
   */
  function unlockableAmountWithinRange(address account, address locker, uint40 unlockedAtStart, uint40 unlockedAtEnd)
    public
    view
    returns (uint256 unlockableByOwner, uint256 unlockableByLocker)
  {
    LockedPosition[] memory positions;
    (unlockableByLocker,, positions) = _filterLockedPositionsBy({
      account: account,
      locker: locker,
      tFrom: unlockedAtStart,
      tTo: unlockedAtEnd,
      stopLimit: MAX_AMOUNT
    });

    uint256 length = positions.length;
    uint256 currTimestamp = block.timestamp;
    for (uint256 i; i < length; ++i) {
      if (positions[i].until <= currTimestamp) {
        unlockableByOwner += positions[i].amount;
      }
    }
  }

  /**
   * @inheritdoc ILRON
   */
  function lockedPositionCount(address account) external view returns (uint256) {
    return _lockedPositions[account].length;
  }

  /**
   * @inheritdoc ILRON
   */
  function lockedPositionAt(address account, uint256 index) external view returns (LockedPosition memory) {
    return _lockedPositions[account][index];
  }

  /**
   * @inheritdoc ILRON
   */
  function totalLocked() external view returns (uint256) {
    return balanceOf(address(this));
  }

  /**
   * @inheritdoc ILRON
   */
  function totalDeposited(address account) external view returns (uint256) {
    return balanceOf(account) + lockedBalanceOf(account);
  }

  /**
   * @inheritdoc ILRON
   */
  function vaultBalanceOf(address locker) external view returns (uint256) {
    return _vaultBalance[locker];
  }

  /**
   * @inheritdoc ILRON
   */
  function lockedBalanceOf(address account) public view returns (uint256 lockedBalance) {
    (lockedBalance,,) = _filterLockedPositionsBy({
      account: account,
      locker: address(0x0),
      tFrom: 0,
      tTo: MAX_TIMESTAMP,
      stopLimit: MAX_AMOUNT
    });
  }

  /**
   * @inheritdoc ILRON
   */
  function lockableAllowance(address account, address locker) public view returns (LockableAllowance memory) {
    return _lockableAllowance[account][locker];
  }

  /**
   * @dev Lock the specified amount of RON for the account until the specified timestamp.
   *
   * Requirements:
   * - The account must not be the zero address.
   * - The amount must be greater than zero.
   *
   * Reverts if the lock duration exceeds the threshold.
   *
   * Emits a {Lock} event.
   *
   * @param account The account of the RON.
   * @param amount The amount to lock.
   */
  function _lock(address account, address locker, uint96 amount, uint40 until) private {
    uint256 currTimestamp = block.timestamp;

    if (locker == address(0x0) || locker == address(this)) revert InvalidLocker(locker);
    if (until <= currTimestamp) revert ForbiddenPastLockUntil(currTimestamp, until);
    if (until > currTimestamp + MAX_LOCK_DURATION) {
      revert ExceedLockDurationThreshold({ duration: until - uint40(currTimestamp), threshold: MAX_LOCK_DURATION });
    }

    _updateVaultBalance({ locker: locker, value: amount, shouldIncrease: true });

    // Finds similar locked positions.
    (, uint256[] memory indices,) =
      _filterLockedPositionsBy({ account: account, locker: locker, tFrom: until, tTo: until, stopLimit: MAX_AMOUNT });

    if (indices.length != 0) {
      _lockedPositions[account][indices[0]].amount += amount;
    } else {
      uint256 currPosCount = _lockedPositions[account].length;

      if (currPosCount >= MAX_LOCK_POSITION) {
        revert ExceedMaxLockPosition({ maxPosCount: MAX_LOCK_POSITION, currPosCount: currPosCount });
      }

      _lockedPositions[account].push(LockedPosition({ amount: amount, until: until, locker: locker }));
    }

    _transfer(account, address(this), amount);

    emit Lock({ account: account, locker: locker, amount: amount, lockedAt: currTimestamp, lockedUntil: until });
  }

  /**
   * @dev Unlock the specified amount of RON for the account from the locker.
   *
   * Requirements:
   * - The account must not be the zero address.
   * - The amount must be greater than zero.
   *
   * If `locker` is not provided, unlock from all lockers.
   *
   * Reverts if the account does not have enough locked balance.
   *
   * Emits an {Unlock} event.
   */
  function _unlock(address account, address locker, uint40 tFrom, uint40 tTo, uint96 amount) private {
    (uint256 totalUnlockable, uint256[] memory unlockIndices,) =
      _filterLockedPositionsBy({ account: account, locker: locker, tFrom: tFrom, tTo: tTo, stopLimit: amount });

    if (totalUnlockable < amount) {
      revert InsufficientLockedBalance({ account: account, balance: totalUnlockable, needed: amount });
    }

    // Remove unlocked positions.
    _removeUnlockedPositions({ account: account, unlockIndices: unlockIndices, stopLimit: amount });

    // Transfer the unlocked amount to the account.
    _transfer(address(this), account, amount);
  }

  /**
   * @dev Update the lockable allowance for the locker.
   *
   * - Does not update the allowance value in case of infinite allowance.
   * - Revert if not enough allowance is available.
   * - Does not emit an {LockApproval} event.
   */
  function _spendLockAllowance(address account, address locker, uint96 amount, uint40 requestedUntil) private {
    LockableAllowance memory mLockableAllowance = lockableAllowance({ account: account, locker: locker });
    uint96 allowed = mLockableAllowance.amount;
    uint40 maxUntil = mLockableAllowance.maxUntil;

    if (maxUntil < requestedUntil) {
      revert ExceedMaxLockUntilAllowance({ maxUntilAllowed: maxUntil, requestedUntil: requestedUntil });
    }

    if (allowed != MAX_AMOUNT) {
      if (allowed < amount) {
        revert InsufficientLockableAllowance({ locker: locker, current: allowed, spent: amount });
      }

      unchecked {
        _approveLock({
          account: account,
          locker: locker,
          amount: allowed - amount,
          maxUntil: maxUntil,
          shouldEmitEvent: false
        });
      }
    }
  }

  /**
   * @dev Approve a specified locker for locking account's RON for specific lock until timestamp.
   *
   * Requirements:
   * - The account must not be the zero address.
   * - The locker must not be the zero address.
   * - The locker must not be the account.
   *
   * Emits a {LockApproval} event.
   * If the `shouldEmitEvent` is false, the event will not be emitted.
   */
  function _approveLock(address account, address locker, uint96 amount, uint40 maxUntil, bool shouldEmitEvent) private {
    if (locker == account) revert SelfApproval();
    if (locker == address(0x0)) revert InvalidLocker(locker);
    if (account == address(0x0)) revert InvalidApprover(account);
    if (amount != 0 && maxUntil <= block.timestamp) revert ForbiddenPastLockUntil(block.timestamp, maxUntil);

    LockableAllowance storage $ = _lockableAllowance[account][locker];

    $.maxUntil = maxUntil;
    $.amount = amount;

    if (!shouldEmitEvent) return;

    emit LockApproval({ account: account, locker: locker, amount: amount, maxUntil: maxUntil });
  }

  /**
   * @dev Delete the unlocked positions and update the locked positions.
   *
   * WARNING: `unlockIndices` must be in descending order.
   *
   * @param unlockIndices The indices of the locked positions to unlock (desc order).
   * @param stopLimit The upper bound limit amount to stop updating.
   */
  function _removeUnlockedPositions(address account, uint256[] memory unlockIndices, uint96 stopLimit) private {
    uint96 unlockedAmount;
    LockedPosition memory mLP;

    address by = _msgSender();
    bool selfUnlock = by == account; // Whether the unlock is done by the account owner.
    LockedPosition[] storage $lockedPositions = _lockedPositions[account];
    uint256 mLength = unlockIndices.length;
    uint256 sLength = $lockedPositions.length;

    for (uint256 i; i < mLength; ++i) {
      mLP = $lockedPositions[unlockIndices[i]];

      unlockedAmount += mLP.amount;

      unchecked {
        if (unlockedAmount > stopLimit) {
          // Overwrite the locked position with the remaining amount.
          uint96 remain = unlockedAmount - stopLimit;
          uint96 deducted = mLP.amount - remain;
          // Update `unlockedAmount` with the remaining amount.
          unlockedAmount -= remain;

          $lockedPositions[unlockIndices[i]].amount = remain;

          _updateVaultBalance({ locker: mLP.locker, value: deducted, shouldIncrease: false });

          emit Unlock({
            by: by,
            account: account,
            locker: mLP.locker,
            unlockedAt: block.timestamp,
            amount: deducted,
            remain: remain,
            selfUnlock: selfUnlock,
            actualLockedUntil: mLP.until
          });

          break;
        }
      }

      // Replace the unlocked position with the last locked position.
      $lockedPositions[unlockIndices[i]] = $lockedPositions[--sLength];
      // Delete the last locked position.
      delete $lockedPositions[sLength];

      _updateVaultBalance({ locker: mLP.locker, value: mLP.amount, shouldIncrease: false });

      emit Unlock({
        by: by,
        account: account,
        locker: mLP.locker,
        unlockedAt: block.timestamp,
        amount: mLP.amount,
        remain: 0,
        selfUnlock: selfUnlock,
        actualLockedUntil: mLP.until
      });
    }

    // Update storage length
    assembly {
      sstore($lockedPositions.slot, sLength)
    }

    // Safety check if the amount is correctly unlocked.
    if (unlockedAmount != stopLimit) revert IncorrectUnlockAmount({ acquired: unlockedAmount, provided: stopLimit });
  }

  /**
   * @dev Update the vault balance of the locker.
   *
   * Emits a {VaultBalanceUpdated} event.
   *
   * @param locker The locker address.
   * @param value The value to update.
   * @param shouldIncrease The flag to increase or decrease the balance.
   */
  function _updateVaultBalance(address locker, uint256 value, bool shouldIncrease) private {
    uint256 prevBalance = _vaultBalance[locker];
    uint256 newBalance = shouldIncrease ? prevBalance + value : prevBalance - value;

    _vaultBalance[locker] = newBalance;

    emit VaultBalanceUpdated({ locker: locker, prevBalance: prevBalance, newBalance: newBalance });
  }

  /**
   * @dev Filter locked positions in range [tFrom, tTo] with stop limit amount.
   *
   * Filter conditions:
   * - Locked position until timestamp is in the provided unlock time range.
   * - If `stopLimit` is reached, stop searching.
   * - If `locker` is provided, filter the locked positions by the locker.
   *
   * @param account The account address.
   * @param locker The locker address.
   * @param stopLimit The upper bound limit amount to stop searching.
   * @param tFrom The start timestamp to search.
   * @param tTo The end timestamp to search.
   *
   * @return sum Total locked amount of the filtered positions. Can be greater than the limit amount.
   * @return indices The indices of the filtered positions by time range.
   */
  function _filterLockedPositionsBy(address account, address locker, uint40 tFrom, uint40 tTo, uint256 stopLimit)
    private
    view
    returns (uint256 sum, uint256[] memory indices, LockedPosition[] memory positions)
  {
    unchecked {
      if (tFrom > tTo) revert InvalidTimeRange({ start: tFrom, end: tTo });

      LockedPosition[] storage $ = _lockedPositions[account];
      uint256 length = $.length;

      // Return early if there are no locked positions.
      if (length == 0) return (0, new uint256[](0), new LockedPosition[](0));

      uint256 tailIdx = length - 1;
      uint256 posCount;
      indices = new uint256[](length);
      positions = new LockedPosition[](length);

      for (uint256 i = tailIdx; i >= 0; --i) {
        LockedPosition memory mLP = $[i];

        // Include the locked position if the until timestamp is in the range.
        if (mLP.until >= tFrom && mLP.until <= tTo && (locker == mLP.locker || locker == address(0x0))) {
          indices[posCount] = i;
          positions[posCount] = mLP;
          sum += mLP.amount;
          ++posCount;
        }

        // Break if the total unlockable amount is greater than or equal to the limit amount.
        if (sum >= stopLimit) break;
        if (i == 0) break;
      }

      // Update the unlockable indices and positions array if the position count is less than the length.
      if (posCount < length) {
        assembly {
          mstore(indices, posCount)
          mstore(positions, posCount)
        }
      }
    }
  }

  /**
   * @dev Revert if the destination address is the lock address.
   */
  function _requireValidDestination(address to) private view {
    if (to == address(this)) revert UnauthorizedAction(to);
  }

  /**
   * @dev Revert if zero amount is provided.
   */
  function _requireNonNullAmount(uint96 amount) private pure {
    if (amount == 0) revert NullAmountProvided();
  }
}
