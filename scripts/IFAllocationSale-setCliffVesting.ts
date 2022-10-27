// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat

import IFAllocationSale from '../artifacts/contracts/IFAllocationSale.sol/IFAllocationSale.json'

// Runtime Environment's members available in the global scope.
import hre from 'hardhat'
import { Contract } from 'ethers'
import assert from 'assert'


async function setCliff(
    allocationSaleContract: Contract,
    withdrawTime: Date,
    duration: number,
    step: number,
    pct: number,
) {
  const start = withdrawTime

  const cliffCount = duration / step
  const cliffPct = Array(cliffCount).fill(pct)
  const cliffTimes = []
  for (let i = 0; i < cliffCount; i++) {
    withdrawTime.setDate(withdrawTime.getDate() + step)
    cliffTimes.push(withdrawTime.getTime() / 1000)
  }

  assert(cliffTimes.length === cliffPct.length)
  const cliffTimesHuman = cliffTimes.map((t) => {
    return new Date(t * 1000)
  })
  console.log('start at:', start)
  console.log('cliffTime:', cliffTimes)
  console.log('cliffTimesHuman:', cliffTimesHuman)
  console.log('Number of cliffs:', cliffTimes.length)

  // set sale token allocation override
  const result = await allocationSaleContract
    .connect((await hre.ethers.getSigners())[0])
    .setCliffPeriod(
        cliffTimes,
        cliffPct,
    )

  // wait for tx to be mined
  await result.wait()

  // log
  console.log('Cliff Period:', await allocationSaleContract.getCliffPeriod())
  console.log('Tx hash:', result.hash)
}

export default async function main(): Promise<void> {
  // when to start withdraw in unix timestamp
  // https://www.unixtimestamp.com/
  const WITHDRAW_TIME = new Date((parseInt(process.env.WITHDRAW_TIME || '0') * 1000))
  // how many days will it last
  const DURATION = parseInt(process.env.DURATION || '0')
  // how many days to start a each withdraw
  const STEP = parseInt(process.env.STEP || '0')
  // how many percent to unlock in each step
  const PCT = parseInt(process.env.PCT || '0')
  // sale contract address
  const allocationSale: string = process.env.SALE || '' // address

  // get allocationSale contract
  const allocationSaleContract = new hre.ethers.Contract(
    allocationSale,
    IFAllocationSale.abi
  )

  // await setCliff(allocationSaleContract, new Date(2022, 12, 1), 270, 3, 3)
  await setCliff(allocationSaleContract, WITHDRAW_TIME, DURATION, STEP, PCT)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
