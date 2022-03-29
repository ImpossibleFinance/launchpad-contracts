import '@nomiclabs/hardhat-ethers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { Contract } from '@ethersproject/contracts'
import { asyncWriteFile, mineNext, readFile, unparseCsv } from './helpers'
import { simAllocationMaster } from './simulator'

import sim1Input from './simulationData/sim1Input.json'
import sim2Input from './simulationData/sim2Input.json'
import sim3Input from './simulationData/sim3Input.json'
import sim4Input from './simulationData/sim4Input.json'

// array of simulations input/output maps
const simulations = [
  { in: sim1Input, out: './test/simulationData/sim1ExpectedOutput.csv' },
  { in: sim2Input, out: './test/simulationData/sim2ExpectedOutput.csv' },
  { in: sim3Input, out: './test/simulationData/sim3ExpectedOutput.csv' },
  { in: sim4Input, out: './test/simulationData/sim4ExpectedOutput.csv' },
]

export default describe('IFAllocationMaster', function () {
  // unset timeout from the test
  this.timeout(0)

  // vars for all tests
  let owner: SignerWithAddress
  let nonOwner: SignerWithAddress
  let simUser1: SignerWithAddress
  let simUser2: SignerWithAddress
  let TestToken: Contract
  let IFAllocationMaster: Contract

  // setup for each test
  beforeEach(async () => {
    // get test accounts
    owner = (await ethers.getSigners())[0]
    nonOwner = (await ethers.getSigners())[1]
    simUser1 = (await ethers.getSigners())[2]
    simUser2 = (await ethers.getSigners())[3]

    // deploy test token
    const TestTokenFactory = await ethers.getContractFactory('GenericToken')
    TestToken = await TestTokenFactory.deploy(
      'test token',
      'TEST',
      '21000000000000000000000000000' // 21 billion * 10**18
    )

    // deploy allocation master
    const IFAllocationMasterFactory = await ethers.getContractFactory(
      'IFAllocationMaster'
    )
    IFAllocationMaster = await IFAllocationMasterFactory.deploy()
  })

  // TESTS

  it('counts tracks', async () => {
    // num tracks should be 0
    mineNext()
    expect(await IFAllocationMaster.trackCount()).to.equal(0)

    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      1000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )

    // num tracks should be 1
    mineNext()
    expect(await IFAllocationMaster.trackCount()).to.equal(1)
  })

  it('can bump sale counter', async () => {
    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      1000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )
    const trackNum = 0

    // bump sale counter
    mineNext()
    await IFAllocationMaster.bumpSaleCounter(trackNum)
    mineNext()

    // update track as non-owner (should fail)
    mineNext()
    await IFAllocationMaster.connect(nonOwner).bumpSaleCounter(trackNum)
    mineNext()

    // sale counter should update only by owner
    const nTrackCheckpoints = await IFAllocationMaster.trackCheckpointCounts(
      trackNum
    )
    const latestTrackCp = await IFAllocationMaster.trackCheckpoints(
      trackNum,
      nTrackCheckpoints - 1
    )
    mineNext()
    expect(latestTrackCp.numFinishedSales).to.equal(1) // only 1 not 2

    //// user checkpoint should record latest sale count

    // approve
    await TestToken.approve(IFAllocationMaster.address, '1000')
    // stake
    await IFAllocationMaster.stake(trackNum, '1000')
    mineNext()

    // get newly generated checkpoint info
    const nUserCheckpoints = await IFAllocationMaster.userCheckpointCounts(
      trackNum,
      owner.address
    )
    const userCp = await IFAllocationMaster.userCheckpoints(
      trackNum,
      owner.address,
      nUserCheckpoints - 1
    )

    // new user checkpoint's numFinishedSales should match
    expect(userCp.numFinishedSales).to.equal(1)
  })

  it('simulation 1: general staking and unstaking', async () => {
    // allocate stake token to simulation user1 and user2
    mineNext()
    await TestToken.transfer(simUser1.address, '10000000000000000000000000000') // 10B tokens
    await TestToken.transfer(simUser2.address, '10000000000000000000000000000') // 10B tokens

    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      '10000000', // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '10000000000000000000000000000' // max total stake (10B)
    )

    //// block-by-block simulation

    // simulation reference inputs and outputs
    const simIn = simulations[0].in
    const simExpectedOut = simulations[0].out

    // run
    const simOutput = await simAllocationMaster(
      IFAllocationMaster, // staking contract
      TestToken, // stake token
      await IFAllocationMaster.trackCount(), // track number
      [simUser1, simUser2], // simulation users
      simIn
    )

    // // write output to CSV
    await asyncWriteFile(
      './test/simulationData',
      '.tmp.out1.csv',
      unparseCsv(simOutput)
    )

    //// check simulation output against output csv
    // get lines of expected output and simulation
    const expectedLines = (await readFile(simExpectedOut)).split(/\r?\n/)
    const simOutLines = unparseCsv(simOutput).split(/\r?\n/)

    // compare each line
    expectedLines.map((expectedLine, i) => {
      expect(expectedLine).to.equal(simOutLines[i])
    })
  })

  it('simulation 2: rollovers', async () => {
    // allocate stake token to simulation user1 and user2
    mineNext()
    await TestToken.transfer(simUser1.address, '10000000000000000000000000000') // 10B tokens
    await TestToken.transfer(simUser2.address, '10000000000000000000000000000') // 10B tokens

    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      '10000000', // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '10000000000000000000000000000' // max total stake (10B)
    )

    //// block-by-block simulation

    // simulation reference inputs and outputs
    const simIn = simulations[1].in
    const simExpectedOut = simulations[1].out

    // run
    const simOutput = await simAllocationMaster(
      IFAllocationMaster, // staking contract
      TestToken, // stake token
      await IFAllocationMaster.trackCount(), // track number
      [simUser1, simUser2], // simulation users
      simIn
    )

    // // write output to CSV
    await asyncWriteFile(
      './test/simulationData',
      '.tmp.out2.csv',
      unparseCsv(simOutput)
    )

    //// check simulation output against output csv
    // get lines of expected output and simulation
    const expectedLines = (await readFile(simExpectedOut)).split(/\r?\n/)
    const simOutLines = unparseCsv(simOutput).split(/\r?\n/)

    // compare each line
    expectedLines.map((expectedLine, i) => {
      expect(expectedLine).to.equal(simOutLines[i])
    })
  })

  it('simulation 3: disabling track', async () => {
    // allocate stake token to simulation user1 and user2
    mineNext()
    await TestToken.transfer(simUser1.address, '10000000000000000000000000000') // 10B tokens
    await TestToken.transfer(simUser2.address, '10000000000000000000000000000') // 10B tokens

    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      '10000000', // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '10000000000000000000000000000' // max total stake (10B)
    )

    //// block-by-block simulation

    // simulation reference inputs and outputs
    const simIn = simulations[2].in
    const simExpectedOut = simulations[2].out

    // run
    const simOutput = await simAllocationMaster(
      IFAllocationMaster, // staking contract
      TestToken, // stake token
      await IFAllocationMaster.trackCount(), // track number
      [simUser1, simUser2], // simulation users
      simIn
    )

    // // write output to CSV
    await asyncWriteFile(
      './test/simulationData',
      '.tmp.out3.csv',
      unparseCsv(simOutput)
    )

    //// check simulation output against output csv
    // get lines of expected output and simulation
    const expectedLines = (await readFile(simExpectedOut)).split(/\r?\n/)
    const simOutLines = unparseCsv(simOutput).split(/\r?\n/)

    // compare each line
    expectedLines.map((expectedLine, i) => {
      expect(expectedLine).to.equal(simOutLines[i])
    })
  })

  it('simulation 4: emergency withdraws', async () => {
    // allocate stake token to simulation user1 and user2
    mineNext()
    await TestToken.transfer(simUser1.address, '10000000000000000000000000000') // 10B tokens
    await TestToken.transfer(simUser2.address, '10000000000000000000000000000') // 10B tokens

    // add a track
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      '10000000', // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '10000000000000000000000000000' // max total stake (10B)
    )

    //// block-by-block simulation

    // simulation reference inputs and outputs
    const simIn = simulations[3].in
    const simExpectedOut = simulations[3].out

    // run
    const simOutput = await simAllocationMaster(
      IFAllocationMaster, // staking contract
      TestToken, // stake token
      await IFAllocationMaster.trackCount(), // track number
      [simUser1, simUser2], // simulation users
      simIn
    )

    // // write output to CSV
    await asyncWriteFile(
      './test/simulationData',
      '.tmp.out4.csv',
      unparseCsv(simOutput)
    )

    //// check simulation output against output csv
    // get lines of expected output and simulation
    const expectedLines = (await readFile(simExpectedOut)).split(/\r?\n/)
    const simOutLines = unparseCsv(simOutput).split(/\r?\n/)

    // compare each line
    expectedLines.map((expectedLine, i) => {
      expect(expectedLine).to.equal(simOutLines[i])
    })
  })

  it('can add track checkpoint', async () => {
    mineNext()
    await IFAllocationMaster.addTrack(
      'TEST Track', // name
      TestToken.address, // stake token
      1000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )

    //check for track number/check id

    const trackNum = 0

    //call the addTrackCheckpoint function
    let _bumpSaleCounter = true
    let disabled = true
    mineNext()
    await IFAllocationMaster.addTrackCheckpoint(
      trackNum,
      '2000',
      true,
      disabled,
      _bumpSaleCounter
    )

    mineNext()
    const nTrackCheckpoints = await IFAllocationMaster.trackCheckpointCounts(
      trackNum
    )

    mineNext()
    if (nTrackCheckpoints == 0) {
      await IFAllocationMaster.trackCheckpoints(trackNum, 0)
      const nTrackCheckpoints2 = await IFAllocationMaster.trackCheckpointCounts(
        trackNum
      )
      expect(nTrackCheckpoints2).to.equal(nTrackCheckpoints + 1)
    } else {
      mineNext()
      //call track
      const latestTrackCp = await IFAllocationMaster.trackCheckpoints(
        trackNum,
        nTrackCheckpoints - 1
      )

      if (latestTrackCp.disabled) {
        //should revert
      } else {
        mineNext()
        //calculate additional block
        //calculate additional block
        // current block number
        const currBlockNum = await ethers.provider.getBlockNumber()
        //check lastblock
        const lastBlockNum = latestTrackCp.blockNumber.toNumber()
        //additional block
        const additionalBlocks = currBlockNum - lastBlockNum

        mineNext()
        //get weight acrrual rate in track
        //get track

        const track = await IFAllocationMaster.tracks(trackNum)
        const weightAccrualRate = track.weightAccrualRate
        console.log(track, 'track details')

        //calculate total staked in last track
        const totalStaked = latestTrackCp.totalStaked.toNumber()
        console.log(totalStaked, 'total staked converted')

        //calculate margin accrued stake weight
        const marginalAccruedStakeWeight =
          (additionalBlocks * weightAccrualRate * totalStaked) / 10 ** 18
        //calculate new stake weight
        const newStakeWeight =
          latestTrackCp.totalStakeWeight.toNumber() + marginalAccruedStakeWeight

        console.log(
          marginalAccruedStakeWeight,
          newStakeWeight,
          'margin and new stake weight'
        )

        if (_bumpSaleCounter) {
          mineNext()
          //get active roll over weight
          /*  mineNext()
          const activeRollOverWeight = (await IFAllocationMaster.trackTotalActiveRollOvers(
            trackNum, latestTrackCp.numFinishedSales)).toNumber();

            mineNext()
            const passiveRollOverWeight = track.passiveRolloverRate;
            console.log(activeRollOverWeight, passiveRollOverWeight, 'active roll over weight without')
          
            //calculate another newStake weight then run assert
           mineNext()
            const b_NewStakeWeight = 
            (activeRollOverWeight * track.activeRollOverWeight)/ 10**18 + 
            (newStakeWeight - activeRollOverWeight)

         */
          //    console.log(latestTrackCp, 'last track cp')
          expect(
            await IFAllocationMaster.addTrackCheckpoint(
              trackNum,
              '2000',
              true,
              true,
              _bumpSaleCounter
            )
          )
            .to.emit(IFAllocationMaster, 'BumpSaleCounter')
            .withArgs(trackNum, latestTrackCp.numFinishedSales + 1)
        }
        mineNext()

        if (additionalBlocks == 0) {
          //
        } else {
          //new track check point should increase
          await IFAllocationMaster.trackCheckpoints(trackNum, nTrackCheckpoints)
          const nTrackCheckpoints2 =
            await IFAllocationMaster.trackCheckpointCounts(trackNum)
          expect(nTrackCheckpoints2).to.equal(nTrackCheckpoints + 1)
        }

        if (!latestTrackCp.disabled && disabled) {
          expect(
            await IFAllocationMaster.addTrackCheckpoint(
              trackNum,
              '2000',
              true,
              true,
              _bumpSaleCounter
            )
          )
            .to.emit(IFAllocationMaster, 'DisableTrack')
            .withArgs(trackNum)
        }
      }
    }

    expect(
      await IFAllocationMaster.addTrackCheckpoint(
        trackNum,
        '2000',
        true,
        true,
        _bumpSaleCounter
      )
    )
      .to.emit(IFAllocationMaster, 'AddTrackCheckpoint')
      .withArgs(trackNum, await ethers.provider.getBlockNumber())
  })
})
