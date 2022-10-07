// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'
import chalk from 'chalk'

// Run this script on staging chain (testnet)
export async function main(): Promise<void> {
  const signerAddress = (await hre.ethers.getSigners())[0].address
  console.log(chalk.magentaBright('Deployer'), chalk.green(signerAddress))

  // MASTER
  const LoyaltyCardMasterFactory = await hre.ethers.getContractFactory(
    'LoyaltyCardMaster'
  )
  const loyaltyCardMaster = await LoyaltyCardMasterFactory.deploy(
    'ImpossibleLoyaltyCard',
    'ILC'
  )
  console.log(
    'LoyaltyCardMaster deployed to ',
    chalk.green(loyaltyCardMaster.address)
  )

  // LOOKUP
  const LoyaltyRewardsLookupFactory = await hre.ethers.getContractFactory(
    'LoyaltyRewardsLookup'
  )
  const loyaltyRewardsLookup = await LoyaltyRewardsLookupFactory.deploy()
  console.log(
    'LoyaltyRewardsLookup deployed to ',
    chalk.green(loyaltyRewardsLookup.address)
  )

  // REWARDER
  const LoyaltyCardRewarderFactory = await hre.ethers.getContractFactory(
    'LoyaltyCardRewarder'
  )
  const loyaltyCardRewarder = await LoyaltyCardRewarderFactory.deploy(
    loyaltyCardMaster.address,
    loyaltyRewardsLookup.address
  )
  console.log(
    'LoyaltyCardRewarder deployed to ',
    chalk.green(loyaltyCardRewarder.address)
  )

  // Setup credentials - use realistic values

  const credCodes = [1, 2, 3]
  const credPoints = [1, 1, 2]
  const credNames = ['dao', 'swap1', 'stake1']

  for (const i of [0, 1, 2])
    await loyaltyRewardsLookup.setCredential(
      credCodes[i],
      credPoints[i],
      credNames[i]
    )
  console.log('\nCredentials set in LoyaltyRewardsLookup')
  console.log('Codes: ', credCodes)
  console.log('Names: ', credNames)
  console.log('Points: ', credPoints)
  console.log()

  // Allow the deployer of the loyalty card master to mint the nfts
  // IN PRODUCTION WE NEED TO MAKE SURE THIS IS AS INTENDED
  await loyaltyCardMaster.setMinter(signerAddress)
  console.log(
    'Set',
    chalk.blueBright('Minter'),
    ' in LoyaltyCardMaster to',
    chalk.green(signerAddress)
  )
  // The rewarder contract is allowed to operate on loyalty card nfts
  await loyaltyCardMaster.addOperator(loyaltyCardRewarder.address)
  console.log(
    'Deployed LoyaltyCardRewarder added as',
    chalk.blueBright('Operator'),
    ' in LoyaltyCardMaster',
    chalk.green(loyaltyCardRewarder.address)
  )
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
