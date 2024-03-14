# Privileged Roles Guide

## `IFFixedSale.sol`

### Owner
- **`setVestedGiveaway()`**  
  Toggle between making the giveaway a standard fixed payout-giveaway or a vested-giveaway, activating the vesting logic.
  
- **`setPublicAllocation()`**  
  Set the public allocation granted to all users in case users are not on the whitelist.

## `IFSale.sol`

### Owner
- **`setWithdrawDelay()`**  
  Set an additional withdraw delay after the sale period.
  
- **`setLinearVestingEndTime()`**  
  Activate and define the terms of the vesting schedule, potentially deactivating a previous cliff schedule.
  
- **`setCliffPeriod()`**  
  Activate and define the terms of the cliff schedule, potentially deactivating a previous vesting schedule.

## `IFVestable.sol`

### Owner
- **`setVestingEditable()`**  
  Toggle the ability to switch between the cliff and vesting schedule, even if the withdrawal period has started.
  
- **`setLinearVestingEndTime()`**  
  Adjust the end time of the vesting schedule, before the withdrawal period or if `vestingEditable` is enabled.
  
- **`setCliffPeriod()`**  
  Adjust the end time of the cliff schedule, before the withdrawal period or if `vestingEditable` is enabled.

## `IFFundable.sol`

### Owner
- **`setCasher()`**  
  Update the casher address at all times.
  
- **`setFunder()`**  
  Update the funder address at all times.
  
- **`setWithdrawDelay()`**  
  Before the sale period, add and adjust an additional withdraw delay, not exceeding five years.
  
- **`emergencyTokenRetrieve()`**  
  After the sale period, withdraw funds from arbitrary tokens which are not the saleToken.

### Owner's Privileges with Casher
- The owner shares all privileges of Casher. Notably, `cash()` and `cashPaymentToken()` send funds to `msg.sender`, not the funder address, allowing the owner to technically front-run the Casher address.

### Casher
- **`cash()`**  
  After the sale period, withdraw all payment token and unsold saleToken.
  
- **`cashPaymentToken()`**  
  Withdraw up to the whole balance of payment token the contract has received.

### Funder
- **`fund()`**  
  Add additional tokens for sale up until the start time of sale.

## `IFPurchasable.sol`

### Owner
- **`setIsPurchaseHalted()`**  
  Pause all purchasing activities.
  
- **`setMinTotalPayment()`**  
  Set an optional minimum payment a user must make to purchase saleTokens.
  
- **`setMaxTotalPurchasable()`**  
  Set an optional maximum payment a user can make to limit the amount of purchasable tokens.

## `IFWhitelistable.sol`

### Owner
- **`setWhitelistSetter()`**  
  Set an address to hold the WhitelistSetter-role. The owner shares all privileges with the WhitelistSetter.

### WhitelistSetter
- **`setWhitelist()`**  
  Set and unset the root of the Merkle tree used for the whitelisting feature at all times.
