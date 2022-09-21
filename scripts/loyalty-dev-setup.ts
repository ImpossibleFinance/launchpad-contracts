// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'
import chalk from 'chalk'

export async function main(): Promise<void> {
  // MASTER
  const LoyaltyCardMasterFactory = await hre.ethers.getContractFactory(
    'LoyaltyCardMaster'
  )
  const loyaltyCardMaster = await LoyaltyCardMasterFactory.deploy(
    'ImpossibleLoyaltyCard',
    'ILC'
  )
  console.log('LoyaltyCardMaster deployed to ', loyaltyCardMaster.address)

  // LOOKUP
  const LoyaltyRewardsLookupFactory = await hre.ethers.getContractFactory(
    'LoyaltyRewardsLookup'
  )
  const loyaltyRewardsLookup = await LoyaltyRewardsLookupFactory.deploy()
  console.log('LoyaltyRewardsLookup deployed to ', loyaltyRewardsLookup.address)

  // REWARDER
  const LoyaltyCardRewarderFactory = await hre.ethers.getContractFactory(
    'LoyaltyCardRewarder'
  )
  const loyaltyCardRewarder = await LoyaltyCardRewarderFactory.deploy(
    loyaltyCardMaster.address,
    loyaltyRewardsLookup.address
  )
  console.log('LoyaltyCardRewarder deployed to ', loyaltyCardRewarder.address)

  // make sure ganache cli --deterministic and addresses are as expected
  const deploymentOK =
    loyaltyCardMaster.address == '0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab' &&
    loyaltyRewardsLookup.address ==
      '0x5b1869D9A4C187F2EAa108f3062412ecf0526b24' &&
    loyaltyCardRewarder.address == '0xCfEB869F69431e42cdB54A4F4f105C19C080A601'
  if (!deploymentOK) {
    console.log(
      chalk.red(
        'Contract addresses not as expected. Are you using ganace-cli --deterministic?'
      )
    )
  }

  const credCodes = [1, 2, 3]
  const credPoints = [11, 12, 13]
  const credNames = ['dao', 'swap1', 'stake1']

  for (const i of [0, 1, 2])
    await loyaltyRewardsLookup.setCredential(
      credCodes[i],
      credPoints[i],
      credNames[i]
    )

  // The deployer of the loyalty card master is allowed to mint the nfts
  // IN PRODUCTION WE NEED TO MAKE SURE THIS IS AS INTENDED
  // (who deploys LoyaltyCardMaster vs. who is the loyalty RewarderWallet)
  const signerAddress = (await hre.ethers.getSigners())[0].address
  await loyaltyCardMaster.setMinter(signerAddress)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
