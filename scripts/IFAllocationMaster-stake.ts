// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat

import IFAllocationMaster from '../artifacts/contracts/IFAllocationMasterSource.sol/IFAllocationMasterSource.json'
import GenericToken from '../artifacts/contracts/GenericToken.sol/GenericToken.json'

// Runtime Environment's members available in the global scope.
import hre from 'hardhat'

export async function main() {
  // params
  const allocationMaster = process.env.ALLOCATION_MASTER || '' // allocation master
  const trackId: string = process.env.TRACK_ID || '0' // track ID
  const token = process.env.TOKEN || '' // address
  const stakeAmount = process.env.AMOUNT || 0

  // get allocationMaster contract
  const allocationMasterContract = new hre.ethers.Contract(
    allocationMaster,
    IFAllocationMaster.abi
  )

  // get token contract
  const tokenContract = new hre.ethers.Contract(token, GenericToken.abi)

  // approve
  await tokenContract
    .connect((await hre.ethers.getSigners())[0])
    .approve(allocationMasterContract.address, stakeAmount)

  // stake
  const result = await allocationMasterContract
    .connect((await hre.ethers.getSigners())[0])
    .stake(trackId, stakeAmount)

  // log
  console.log('Track ID:', trackId)
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
