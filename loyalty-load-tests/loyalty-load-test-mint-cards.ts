// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'
import chalk from 'chalk'
import fs from 'fs'

export async function main(): Promise<void> {
  // values for local chain run with $ ganache-cli --deterministic
  const GANACHE_CARD_MASTER_ADDRESS =
    '0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab'
  const LIMIT = 200

  const loyaltyCardMaster = await hre.ethers.getContractAt(
    'LoyaltyCardMaster',
    GANACHE_CARD_MASTER_ADDRESS
  )

  const addressesAsText = fs.readFileSync(
    'loyalty-load-tests/addresses',
    'utf8'
  )
  const addresses = addressesAsText.split('\n')
  const subset = addresses.filter((v, i) => i < LIMIT)
  const tx = await loyaltyCardMaster.mintForNonOwners(subset, {
    gasLimit: 20000000,
  })
  console.log(tx)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
