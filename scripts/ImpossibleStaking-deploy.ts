// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'

export async function main(): Promise<void> {
  const rewardToken: string = process.env.REWARD_TOKEN || '' // reward token address
  const rewardPerTimestamp: string = process.env.REWARD_PER_TIMESTAMP || '' // reward per timestamp
  const startTime: string = process.env.START_TIME || '' // start time of reward distribution
  const endTime: string = process.env.END_TIME || '' // end time of reward distribution

  // We get the contract to deploy
  const ImpossibleStakingFactory = await hre.ethers.getContractFactory(
    'ImpossibleStaking'
  )

  // deploy token
  const ImpossibleStaking = await ImpossibleStakingFactory.deploy(
    rewardToken,
    rewardPerTimestamp,
    startTime,
    endTime
  )

  // log deployed addresses
  console.log('Token deployed to ', ImpossibleStaking.address)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
