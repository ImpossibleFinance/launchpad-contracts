// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'

export async function main(): Promise<void> {
  const messageBus = process.env.MESSAGE_BUS || '' // celer's message bus
  const omniAddress = process.env.OMNI_ADDRESS || '' // contract address on omni chain
  const omniPortal = process.env.OMNI_PORTAL || '0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985' // omni portal address

  // We get the contract to deploy
  const IFAllocationMasterFactory = await hre.ethers.getContractFactory(
    'IFAllocationMasterSource'
  )
  const IFAllocationMaster = await IFAllocationMasterFactory.deploy(messageBus)

  await IFAllocationMaster
    .connect((await hre.ethers.getSigners())[0])
    .setOmniAddress(omniAddress)


  await IFAllocationMaster
    .connect((await hre.ethers.getSigners())[0])
    .setOmniPortal(omniPortal)

  console.log('IFAllocationMasterSource deployed to ', IFAllocationMaster.address)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
