// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import hre from 'hardhat'

export async function main(): Promise<void> {
  const messageBus = process.env.MESSAGE_BUS || '' // celer's message bus
  const srcAddress = process.env.SOURCE_ADDRESS || ''
  const srcChainId = process.env.SOURCE_CHAIN || ''

  // We get the contract to deploy
  const IFAllocationMasterAdapterFactory = await hre.ethers.getContractFactory(
    'IFAllocationMasterAdapter'
  )
  const IFAllocationMasterAdapter =
    await IFAllocationMasterAdapterFactory.deploy(
      messageBus,
      srcAddress,
      srcChainId
    )

  console.log(
    'IFAllocationMasterAdapter deployed to ',
    IFAllocationMasterAdapter.address
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
