// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat

import IFAllocationMasterOmni from '../artifacts/contracts/IFAllocationMasterOmni.sol/IFAllocationMasterOmni.json'

// Runtime Environment's members available in the global scope.
import hre from 'hardhat'

export async function main() {
  // params
  const allocationMaster = process.env.ALLOCATION_MASTER || '' // allocation master
  const whitelistContract = process.env.WHITELIST_CONTRACT || ''

  // get allocationMaster contract
  const allocationMasterContract = new hre.ethers.Contract(
    allocationMaster,
    IFAllocationMasterOmni.abi
  )

  // whitelist
  const result = await allocationMasterContract
    .connect((await hre.ethers.getSigners())[0])
    .addWhitelistedContract(whitelistContract)

  // log
  console.log('---- Output ----')
  console.log('Tx hash:', result.hash)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
