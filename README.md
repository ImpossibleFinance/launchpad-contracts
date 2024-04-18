# IDIA Launchpad Staking Contracts

In this repo, we will feature a new IDIA staking launchpad mechanism.

For documentation on our launchpad logic, please visit here:
https://docs.impossible.finance/launchpad/smart-contracts

## Setup

```
yarn install
forge build
```

## Test

### Running all tests

```
npx hardhat test
```

### Run foundry test

```
forge test --fork-url $BSC_URL
```

### Running specific tests

```
npx hardhat test --grep "<YOUR TARGET TESTS KEYWORD>"
```

### Inspect transactions on ethernal

Make sure ethernal is installed: https://doc.tryethernal.com/getting-started/quickstart

Spin up local node

```
npx hardhat node --fork <NODE RPC URL>
```

Turn on ethernal listener

```
ethernal listen
```

Import ethernal to the test script

```typescript
import 'hardhat-ethernal'
```

Run test case with ethernal credentials. Connect it to local node.

```
ETHERNAL_EMAIL=<YOUR EMAIL> ETHERNAL_PASSWORD=<YOUR PASSWORD> npx hardhat run <FILE PATH> --network localhost
```

Login and browse the transactions at https://app.tryethernal.com

## Deploy

### Deploy commands

MESSAGE_BUS is address of celer's message bus, this is required for cross chain data transfer
To get list of the address, follow this url https://im-docs.celer.network/developer/contract-addresses-and-rpc-info

```
# allocation master
MESSAGE_BUS=0xABC npx hardhat run ./scripts/IFAllocationMaster-deploy.ts --network bsc_test

# allocation master adapter
MESSAGE_BUS=0xABC SOURCE_ADDRESS=0xABC SOURCE_CHAIN={{chainId}} npx hardhat run ./scripts/IFAllocationMasterAdapter-deploy.ts --network bsc_test

# allocation sale
SELLER=0xABCD PAY_TOKEN=0xABCD SALE_TOKEN=0xABCD ALLOCATION_MASTER=0xABCD TRACK_ID=123 SNAP_BLOCK=123456 START_BLOCK=123456 END_BLOCK=123456 SALE_PRICE=100000000000000000000 MAX_TOTAL_PAYMENT=10000000000000000000000 npx hardhat run ./scripts/IFAllocationSale-deploy.ts --network bsc_test

# vIDIA
NAME=VIDIA SYMBOL=VIDIA ADMIN=0xABCD UNDERLYING=0xABCD npx hardhat run ./scripts/VIDIA-deploy.ts --network bsc_test

# Verify vIDIA
npx hardhat verify --network bsc_test <DEPLOYED_CONTRACT_ADDRESS> "VIDIA" "VIDIA" "<ADMIN_ADDRESS>" "<UNDERLYING_ADDRESS>"
```

### Production

For production, the deploy command is similar to the one for testnet but you must change the network to `bsc_main`.

You also must add a account / mnemonic in a file named `.env` in the root of the repo with the contents:

```
MAINNET_MNEMONIC='example example example example...'
```

## Other utilities

### Sending tokens

```
TOKEN=0x... TO=0x... AMOUNT=10000000000000000000000 npx hardhat run ./scripts/GenericToken-send.ts --network bsc_test
```

### Pausing and unpausing a pausable token

```
# pause
TOKEN=0x... npx hardhat run ./scripts/GenericToken-pause.ts --network bsc_test

# unpause
TOKEN=0x... npx hardhat run ./scripts/GenericToken-unpause.ts --network bsc_test
```

### Deploying a standard mintable pausable token

```
NAME='Token Name' SYMBOL='TKN1' INIT_SUPPLY=... npx hardhat run ./scripts/GenericToken-deploy.ts --network bsc_test
```

### Minting token

```
TOKEN=0x... TO=0x... AMOUNT=... npx hardhat run ./scripts/GenericToken-mint.ts --network bsc_test
```

### Adding an allocation master track

```
ALLOCATION_MASTER=0xABCD TRACK_NAME='Track Name' TOKEN=0xABCD ACCRUAL_RATE=1000 PASSIVE_RO_RATE=100000000000000000 ACTIVE_RO_RATE=200000000000000000 MAX_TOTAL_STAKE=1000000000000000000000000 npx hardhat run ./scripts/IFAllocationMaster-addTrack.ts --network bsc_test
```

### Bumping sale counter on track

```
ALLOCATION_MASTER=0xABCD TRACK_ID=n npx hardhat run ./scripts/IFAllocationMaster-bumpSaleCounter.ts --network bsc_test
```

### Funding an allocation sale

```
SALE=0xABCD AMOUNT=10000000000000000000000 npx hardhat run ./scripts/IFAllocationSale-fund.ts --network bsc_test
```

### Setting whitelist on allocation sale

```
# via command line, for a short list
# Note: whitelist passed in as comma separated list (end comma optional). No space allowed after comma.
SALE=0xABCD WHITELIST=0xABCD,0xBCDE,0xCDEF, npx hardhat run ./scripts/IFAllocationSale-setWhitelist.ts --network bsc_test

# via file containing JSON list of address strings, for a long list
SALE=0xABCD WHITELIST_JSON_FILE=/path/to/addresses.json npx hardhat run ./scripts/IFAllocationSale-setWhitelist.ts --network bsc_test

# using optional second whitelist for intersection
SALE=0xABCD WHITELIST_JSON_FILE=/path/to/addresses.json WHITELIST_JSON_FILE_2=/path/to/addresses2.json npx hardhat run ./scripts/IFAllocationSale-setWhitelist.ts --network bsc_test
```

### Overriding Sale Token Allocation

```
SALE=0xABCD ALLOCATION=1000000000000000000000 npx hardhat run ./scripts/IFAllocationSale-setSaleTokenAllocationOverride.ts --network bsc_test
```

### Setting a delay for claim

```
SALE=0xABCD DELAY=100 npx hardhat run ./scripts/IFAllocationSale-setWithdrawDelay.ts --network bsc_test
```

### Setting a casher

```
SALE=0xABCD CASHER=0xABCD npx hardhat run ./scripts/IFAllocationSale-setCasher.ts --network bsc_test
```

### Transfering ownership

```
SALE=0xABCD NEW_OWNER=0xABCD npx hardhat run ./scripts/IFAllocationSale-transferOwnership.ts --network bsc_test
```

### Cashing

```
SALE=0xABCD npx hardhat run ./scripts/IFAllocationSale-cash.ts --network bsc_test
```

### Setting cliff periods

```
// To set cliff starting at 2022-OCT-27, lasting for 270 days, unlock every 3 days with 1 percent
SALE=0xABCD WITHDRAW_TIME=1666843153 DURATION=270 STEP=3 PCT=1 npx hardhat run scripts/IFAllocationSale-setCliffVesting.ts                            
```

## Local Development

### Init Loyalty Program contracts

This will deploy the following contracts locally with addresses that those preset in backend-service for local development
- LoyaltyCardMaster.sol
- LoyaltyRewardsLookup.sol
- LoyaltyCardRewarder.sol

Additionally it will initialize the RewardsLookup contract with some credential values which should match the preset values used in the backend-service to populate the loyalty_credential table in dev-mode

```
CODE    POINTS   NAME

1       11       'dao'
2       12       'swap1'
3       13       'stake1'
```

```shell
# terminal 1 - keeps logging blockchain
ganache-cli --deterministic

# terminal 2 - loyalty setup output
npx hardhat run --network localhost ./scripts/loyalty-dev-setup.ts
```


## Setup local contracts

Start a hardhat node. It will fork from bsc mainnet and start a JSON-RPC server at http://127.0.0.1:8545/
```bash
npx hardhat node
```

Specify rpc url and block number to fork from bsc testnet.
```bash
npx hardhat node --fork https://data-seed-prebsc-1-s3.binance.org:8545 --fork-block-height <BLOCK_NUMBER>
```

Deploy allocation master. We'll need celer message bus address. It can be found here: https://im-docs.celer.network/developer/contract-addresses-and-rpc-info
```bash
MESSAGE_BUS=0x95714818fdd7a5454F73Da9c777B3ee6EbAEEa6B npx hardhat run scripts/IFAllocationMaster-deploy.ts --network localhost
```

Deploy sale contract. Get the allocation master address from the previous deployment. Put it to ALLOCATION_MASTER.
```bash
SELLER=0x54F5A04417E29FF5D7141a6d33cb286F50d5d50e PAY_TOKEN=0x0b15Ddf19D47E6a86A56148fb4aFFFc6929BcB89 SALE_TOKEN=0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 ALLOCATION_MASTER=<ALLOCATION_MASTER_ADDRESS> TRACK_ID=1 SNAP_BLOCK=1667377037 START_BLOCK=1667377037 END_BLOCK=1767377037 SALE_PRICE=100000000000000000000 MAX_TOTAL_PAYMENT=10000000000000000000000 npx hardhat run ./scripts/IFAllocationSale-deploy.ts --network localhost
```

## Compile contracts into go files
The base path taken by the compile script is `./contract`. 
```bash
bash scripts/compile-file.sh <CONTRACT_NAME>
```