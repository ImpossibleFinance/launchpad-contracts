import '@nomiclabs/hardhat-ethers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import {
  getBlockTime,
  mineNext,
  mineTimeDelta,
  minePause,
  mineStart,
  setAutomine,
} from './helpers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { Contract } from '@ethersproject/contracts'
import { ALREADY_CASHED, ALREADY_OPTED_IN, BUY_BACK_NOT_ENABLED, NO_TOKEN_TO_BE_WITHDRAWN, NOT_CASHER_OR_OWNER, NOT_OWNER, NOT_FUNDER, USE_WITHDRAWGIVEAWAY, CANNOT_WITHDRAW_BEFORE_CLAIM, ADDRESS_ZERO_FUNDER, NOT_ENOUGH_PAYMENT_TOKEN_TO_CASH, PURCHASE_IS_HALTED, CAN_ONLY_BUY_INTEGER_AMOUNT } from './reverts/msg-IFAllocationSale'

export const _ctx ={
  owner: SignerWithAddress,
  buyer: SignerWithAddress,
  buyer2: SignerWithAddress,
  seller: SignerWithAddress,
  casher: SignerWithAddress,
  StakeToken: Contract,
  PaymentToken: Contract,
  SaleToken: Contract,
  IFAllocationMaster: Contract,
  IFAllocationSale: Contract,
  trackId: 0,
  // sale contract vars
  snapshotTimestamp: 0,// block at which to take allocation snapshot
  startTime: 0, // start timestamp of sale (inclusive)
  endTime: 0, // end timestamp of sale (inclusive)
  linearVestingEndTime: 0, // end timestamp of vesting
  salePrice: '10000000000000000000', // 10 PAY per SALE
  maxTotalDeposit: '25000000000000000000000000', // max deposit
  // other vars
  // const ctx.fundAmount = '33333'
  fundAmount: '1000000000'
}

export const _ctxFree = {
  owner: SignerWithAddress,
  buyer: SignerWithAddress,
  buyer2: SignerWithAddress,
  seller: SignerWithAddress,
  casher: SignerWithAddress,
  StakeToken: Contract,
  PaymentToken: Contract,
  SaleToken: Contract,
  IFAllocationMaster: Contract,
  IFAllocationSale: Contract,
  trackId: 0,
  // sale contract vars
  snapshotTimestamp: 0,// block at which to take allocation snapshot
  startTime: 0, // start timestamp of sale (inclusive)
  endTime: 0, // end timestamp of sale (inclusive)
  linearVestingEndTime: 0, // end timestamp of vesting
  salePrice: '0', // FREE
  maxTotalDeposit: '25000000000000000000000000', // max deposit
  // other vars
  // const ctx.fundAmount = '33333'
  fundAmount: '1000000000'
}

const salePrice = '3700000000000000000' // 3.7 PAY per SALE

export const _ctxSale ={
  owner: SignerWithAddress,
  buyer: SignerWithAddress,
  buyer2: SignerWithAddress,
  seller: SignerWithAddress,
  casher: SignerWithAddress,
  StakeToken: Contract,
  PaymentToken: Contract,
  SaleToken: Contract,
  IFAllocationMaster: Contract,
  IFAllocationSale: Contract,
  trackId: 0,
  // sale contract vars
  snapshotTimestamp: 0,// block at which to take allocation snapshot
  startTime: 0, // start timestamp of sale (inclusive)
  endTime: 0, // end timestamp of sale (inclusive)
  linearVestingEndTime: 0, // end timestamp of vesting
  salePrice: salePrice,
  maxTotalDeposit: '25000000000000000000000000', // max deposit
  // other vars
  // const ctx.fundAmount = '33333'
  fundAmount: '1000000000',
  paymentTokenPerSaleToken: 3.7
}

export default function (_this: Mocha.Suite, contractName: string, ctx: any, ctxFree: any, ctxSale: any) {
  // unset timeout from the test
  _this.timeout(0)

  _this.beforeAll(async () => {
    await setAutomine(false)
  })

  _this.afterAll(async () => {
    await setAutomine(true)
  })
  // setup for each test
  beforeEach(async () => {
    // set launchpad blocks in future
    mineNext()
    let currTime = await getBlockTime()
    mineNext()
    ctx.snapshotTimestamp = currTime + 5000
    ctx.startTime = currTime + 10000
    ctx.endTime = currTime + 20000
    ctx.linearVestingEndTime = currTime + 50000

    // get test accounts
    ctx.owner = (await ethers.getSigners())[0]
    ctx.buyer = (await ethers.getSigners())[1]
    ctx.seller = (await ethers.getSigners())[2]
    ctx.casher = (await ethers.getSigners())[3]
    ctx.buyer2 = (await ethers.getSigners())[4]

    // deploy test tokens
    const TestTokenFactory = await ethers.getContractFactory('GenericToken')
    ctx.StakeToken = await TestTokenFactory.connect(ctx.buyer).deploy(
      'Test Stake Token',
      'STAKE',
      '21000000000000000000000000' // 21 million * 10**18
    )
    ctx.PaymentToken = await TestTokenFactory.connect(ctx.buyer).deploy(
      'Test Payment Token',
      'PAY',
      '21000000000000000000000000' // 21 million * 10**18
    )
    ctx.SaleToken = await TestTokenFactory.connect(ctx.seller).deploy(
      'Test Sale Token',
      'SALE',
      '21000000000000000000000000' // 21 million * 10**18
    )

    // redistribute tokens
    mineNext()
    ctx.StakeToken.connect(ctx.buyer).transfer(
      ctx.buyer2.address,
      '1000000000000000000000000'
    )
    ctx.PaymentToken.connect(ctx.buyer).transfer(
      ctx.buyer2.address,
      '1000000000000000000000000'
    )

    // deploy allocation master
    const IFAllocationMasterFactory = await ethers.getContractFactory(
      'IFAllocationMaster'
    )
    ctx.IFAllocationMaster = await IFAllocationMasterFactory.deploy(
      ethers.constants.AddressZero
    )

    // add track on allocation master
    mineNext()
    mineNext()
    await ctx.IFAllocationMaster.addTrack(
      'IDIA track', // name
      ctx.StakeToken.address, // stake token
      10000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )

    // get new track id
    mineNext()
    ctx.trackId = (await ctx.IFAllocationMaster.trackCount()) - 1

    // deploy sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      contractName
    )
    ctx.IFAllocationSale = await IFAllocationSaleFactory.deploy(
      ctx.salePrice,
      ctx.seller.address,
      ctx.PaymentToken.address,
      ctx.SaleToken.address,
      ctx.IFAllocationMaster.address,
      ctx.trackId,
      ctx.snapshotTimestamp,
      ctx.startTime,
      ctx.endTime,
      ctx.maxTotalDeposit
    )
    mineNext()

    // set the ctx.casher address
    await ctx.IFAllocationSale.setCasher(ctx.casher.address)
    mineNext()

    // fund sale
    mineNext()
    await ctx.SaleToken.connect(ctx.seller).approve(
      ctx.IFAllocationSale.address,
      ctx.fundAmount
    ) // approve
    await ctx.IFAllocationSale.connect(ctx.seller).fund(ctx.fundAmount) // fund
    //
    // stake and accrue stake weight
    mineNext()
    const stakeAmount = 100000000000000
    minePause()
    // ctx.buyer 1
    await ctx.StakeToken.connect(ctx.buyer).approve(
      ctx.IFAllocationMaster.address,
      3 * stakeAmount
    ) // approve
    await ctx.IFAllocationMaster.connect(ctx.buyer).stake(ctx.trackId, 3 * stakeAmount) // stake
    // ctx.buyer 2
    await ctx.StakeToken.connect(ctx.buyer2).approve(
      ctx.IFAllocationMaster.address,
      stakeAmount
    ) // approve
    await ctx.IFAllocationMaster.connect(ctx.buyer2).stake(ctx.trackId, stakeAmount) // stake
    mineStart()

    // expect staked amount to match
    mineNext()
    expect(
      (await ctx.StakeToken.balanceOf(ctx.IFAllocationMaster.address)).toString()
    ).to.equal((stakeAmount * 4).toString())

    //fastforward from current block to after snapshot block
    mineTimeDelta(ctx.snapshotTimestamp - (await getBlockTime()))

    //Setup ctxFree
    if (!ctxFree) return

    currTime = await getBlockTime()
    ctxFree.snapshotTimestamp = currTime + 5000
    ctxFree.startTime = currTime + 10000
    ctxFree.endTime = currTime + 20000
    ctxFree.linearVestingEndTime = currTime + 50000

    //Free get test accounts
    ctxFree.owner = (await ethers.getSigners())[0]
    ctxFree.buyer = (await ethers.getSigners())[1]
    ctxFree.seller = (await ethers.getSigners())[2]
    ctxFree.casher = (await ethers.getSigners())[3]
    ctxFree.buyer2 = (await ethers.getSigners())[4]
    ctxFree.StakeToken = ctx.StakeToken
    ctxFree.PaymentToken = ctx.PaymentToken
    ctxFree.SaleToken = ctx.SaleToken

    //Free redistribute tokens
    mineNext()
    ctxFree.StakeToken.connect(ctxFree.buyer).transfer(
      ctxFree.buyer2.address,
      '1000000000000000000000000'
    )
    ctxFree.PaymentToken.connect(ctxFree.buyer).transfer(
      ctxFree.buyer2.address,
      '1000000000000000000000000'
    )

    //Free deploy allocation master
    ctxFree.IFAllocationMaster = await IFAllocationMasterFactory.deploy(
      ethers.constants.AddressZero
    )

    mineNext()
    await ctxFree.IFAllocationMaster.addTrack(
      'IDIA track', // name
      ctxFree.StakeToken.address, // stake token
      10000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )
    mineNext()

    ctxFree.IFAllocationSale = await IFAllocationSaleFactory.deploy(
      ctxFree.salePrice,
      ctxFree.seller.address,
      ctxFree.PaymentToken.address,
      ctxFree.SaleToken.address,
      ctxFree.IFAllocationMaster.address,
      ctxFree.trackId,
      ctxFree.snapshotTimestamp,
      ctxFree.startTime,
      ctxFree.endTime,
      ctxFree.maxTotalDeposit
    )
    mineNext()

    // set the ctx.casher address
    await ctxFree.IFAllocationSale.setCasher(ctxFree.casher.address)
    mineNext()

    // fund sale
    await ctxFree.SaleToken.connect(ctxFree.seller).approve(
      ctxFree.IFAllocationSale.address,
      ctxFree.fundAmount
    )
    await ctxFree.IFAllocationSale.connect(ctxFree.seller).fund(ctxFree.fundAmount) // fund
    mineNext()

    currTime = await getBlockTime()
    ctxSale.snapshotTimestamp = currTime + 5000
    ctxSale.startTime = currTime + 10000
    ctxSale.endTime = currTime + 20000
    ctxSale.linearVestingEndTime = currTime + 50000

    //Free get test accounts
    ctxSale.owner = (await ethers.getSigners())[0]
    ctxSale.buyer = (await ethers.getSigners())[1]
    ctxSale.seller = (await ethers.getSigners())[2]
    ctxSale.casher = (await ethers.getSigners())[3]
    ctxSale.buyer2 = (await ethers.getSigners())[4]
    ctxSale.StakeToken = ctx.StakeToken
    ctxSale.PaymentToken = ctx.PaymentToken
    ctxSale.SaleToken = ctx.SaleToken

    //Free redistribute tokens
    mineNext()
    ctxSale.StakeToken.connect(ctxSale.buyer).transfer(
      ctxSale.buyer2.address,
      '1000000000000000000000000'
    )
    ctxFree.PaymentToken.connect(ctxSale.buyer).transfer(
      ctxSale.buyer2.address,
      '1000000000000000000000000'
    )

    ctxSale.IFAllocationSale = await IFAllocationSaleFactory.deploy(
      ctxSale.salePrice,
      ctxSale.seller.address,
      ctxSale.PaymentToken.address,
      ctxSale.SaleToken.address,
      ctx.IFAllocationMaster.address,
      ctxSale.trackId,
      ctxSale.snapshotTimestamp,
      ctxSale.startTime,
      ctxSale.endTime,
      ctxSale.maxTotalDeposit
    )
    mineNext()

    // set the ctx.casher address
    await ctxSale.IFAllocationSale.setCasher(ctxSale.casher.address)
    mineNext()

    // fund sale
    await ctxSale.SaleToken.connect(ctxSale.seller).approve(
      ctxSale.IFAllocationSale.address,
      ctxSale.fundAmount
    )
    await ctxSale.IFAllocationSale.connect(ctxSale.seller).fund(ctxSale.fundAmount) // fund
    mineNext()
  })

  it('can purchase, withdraw, and cash', async function () {
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    // Failover mechanism: Call emergencyTokenRetrieve while token is sale or payment token
    await expect(ctx.IFAllocationSale.connect(ctx.owner).emergencyTokenRetrieve(ctx.PaymentToken.address)).to.be.reverted

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()

    // expect balance to increase by fund amount
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')

    // test repeated withdraw (should fail)
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)
    mineNext()

    // expect balance to remain the same
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')

    // test cash
    await ctx.IFAllocationSale.connect(ctx.casher).cash()
    // access control: only ctx.casher can cash
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(ctx.IFAllocationSale.connect(ctx.seller).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(ctx.IFAllocationSale.connect(ctx.owner).cash()).to.be.revertedWith(ALREADY_CASHED)
    mineNext()

    // expect balance to increase by cash amount
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal(paymentAmount)

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(1)

    // Failover mechanism: Call emergencyTokenRetrieve while token is sale or payment token
    await expect(ctx.IFAllocationSale.connect(ctx.casher).emergencyTokenRetrieve(ctx.PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    await expect(ctx.IFAllocationSale.connect(ctx.seller).emergencyTokenRetrieve(ctx.PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).emergencyTokenRetrieve(ctx.PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    ctx.IFAllocationSale.connect(ctx.owner).emergencyTokenRetrieve(ctx.PaymentToken.address)
  })

  it('can cash when salePrice is 0', async function () {
    // fast forward from current time to after end time
    mineTimeDelta(ctxFree.endTime - (await getBlockTime()))

    // test cash
    await ctxFree.IFAllocationSale.connect(ctxFree.casher).cash()
    // access control: only ctx.casher can cash
    await expect(ctxFree.IFAllocationSale.connect(ctxFree.buyer).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(ctxFree.IFAllocationSale.connect(ctxFree.seller).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(ctxFree.IFAllocationSale.connect(ctxFree.owner).cash()).to.be.revertedWith(ALREADY_CASHED)
    mineNext()

    // expect balance to increase by cash amount
    expect(await ctx.SaleToken.balanceOf(ctx.casher.address)).to.equal(ctxFree.fundAmount)
  })

  it('can set funder', async function () {
    mineNext()

    await ctx.SaleToken.connect(ctx.seller).transfer(
      ctx.buyer.address,
      ctx.fundAmount,
    )

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    ctx.IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      ctx.seller.address,
      ctx.PaymentToken.address, // doesn't matter
      ctx.SaleToken.address,
      ctx.IFAllocationMaster.address, // doesn't matter
      ctx.trackId, // doesn't matter
      ctx.snapshotTimestamp, // doesn't matter
      ctx.startTime, // doesn't matter
      ctx.endTime, // doesn't matter
      ctx.maxTotalDeposit // doesn't matter
    )
    mineNext()

    // fund sale
    mineNext()

    await ctx.SaleToken.connect(ctx.seller).approve(
      ctx.IFAllocationSale.address,
      ctx.fundAmount
    ) // approve
    await ctx.IFAllocationSale.connect(ctx.seller).fund(ctx.fundAmount) // fund
    // access control: Address other than funder calls fund
    await expect(ctx.IFAllocationSale.connect(ctx.casher).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund

    await ctx.IFAllocationSale.connect(ctx.owner).setFunder(ctx.buyer.address) // set funer address
    await ctx.SaleToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      ctx.fundAmount
    ) // approve
    await ctx.IFAllocationSale.connect(ctx.buyer).fund(ctx.fundAmount) // fund
    mineNext()

    // send fund from funder address, expect error
    await expect(ctx.IFAllocationSale.connect(ctx.casher).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER)
    await expect(ctx.IFAllocationSale.connect(ctx.seller).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER)
  })

  it('cannot set 0x0 as funder', async function () {
    await expect(ctx.IFAllocationSale.connect(ctx.owner).setFunder(ethers.constants.AddressZero)).to.be.revertedWith(ADDRESS_ZERO_FUNDER)
  })

  it('can perform a zero price giveaway sale (unwhitelisted / first come first serve)', async function () {
    mineNext()

    // here set up a new ctx.IFAllocationSale with ctx.salePrice of 0, because
    // provided fixture sale does not have ctx.salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    ctx.IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      ctx.seller.address,
      ctx.PaymentToken.address, // doesn't matter
      ctx.SaleToken.address,
      ctx.IFAllocationMaster.address, // doesn't matter
      ctx.trackId, // doesn't matter
      ctx.snapshotTimestamp, // doesn't matter
      ctx.startTime, // doesn't matter
      ctx.endTime, // doesn't matter
      ctx.maxTotalDeposit // doesn't matter
    )
    mineNext()

    // fund sale
    mineNext()
    await ctx.SaleToken.connect(ctx.seller).approve(
      ctx.IFAllocationSale.address,
      ctx.fundAmount
    ) // approve
    await ctx.IFAllocationSale.connect(ctx.seller).fund(ctx.fundAmount) // fund
    // access control: Address other than funder calls fund
    await expect(ctx.IFAllocationSale.connect(ctx.casher).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).fund(ctx.fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test normal withdraw (should not go through, must go through withdrawGiveaway)
    // access control: Withdraw when sale price is 0
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer2).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()

    // expect balance to be 0 for both participants
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('0')

    // test withdrawGiveaway (should go through)
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer2).withdrawGiveaway([])
    mineNext()

    // expect both participants can claim
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.gt(0)
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.gt(0)

    // test purchaser counter (should be 0! nothing purchased in 0 price sales)
    // note: _this is the only scenario where _this is different from withdrawer counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(0)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  it('can set withdraw delay', async function () {
    mineNext()

    // delay of 10 blocks
    const delay = 10

    // add withdraw delay
    await ctx.IFAllocationSale.setWithdrawDelay(delay)
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw and cash (should fail because need 1 more block)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_BEFORE_CLAIM)
    // access control: Call cash before ctx.endTime + withdrawDelay
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cash())

    mineNext()

    // fails
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
    // fails
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal('0')

    // simulate `delay` time passing
    mineTimeDelta(delay)

    // test withdraw and cash (should work here after delay passed)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    await ctx.IFAllocationSale.connect(ctx.casher).cash()
    // access control: Call cash after ctx.endTime + withdrawDelay twice
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cash()).to.be.revertedWith(ALREADY_CASHED)

    mineNext()

    // expect balance to increase by fund amount
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')
    // expect balance to increase by cash amount
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal(paymentAmount)

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(1)
  })
  it('can set withdraw delay multiple times', async function () {
    mineNext()

    // delay of 10 blocks
    const delay = 10

    // add withdraw delay
    const withdrawTimeInitial = parseInt(await ctx.IFAllocationSale.withdrawTime())
    await ctx.IFAllocationSale.setWithdrawDelay(5)
    expect(await ctx.IFAllocationSale.withdrawTime()).to.equal(withdrawTimeInitial + 5)
    await ctx.IFAllocationSale.setWithdrawDelay(100)
    expect(await ctx.IFAllocationSale.withdrawTime()).to.equal(withdrawTimeInitial + 100)
    await ctx.IFAllocationSale.setWithdrawDelay(delay)
    expect(await ctx.IFAllocationSale.withdrawTime()).to.equal(withdrawTimeInitial + delay)
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw and cash (should fail because need 1 more block)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_BEFORE_CLAIM)
    // access control: Call cash before ctx.endTime + withdrawDelay
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cash())

    mineNext()

    // fails
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
    // fails
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal('0')

    // simulate `delay` time passing
    mineTimeDelta(delay)

    // test withdraw and cash (should work here after delay passed)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    await ctx.IFAllocationSale.connect(ctx.casher).cash()
    // access control: Call cash after ctx.endTime + withdrawDelay twice
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cash()).to.be.revertedWith(ALREADY_CASHED)

    mineNext()

    // expect balance to increase by fund amount
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')
    // expect balance to increase by cash amount
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal(paymentAmount)

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(1)
  })
  it('does not over cash', async function () {
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // cash first (testing that we do not over-remove sale token)
    await ctx.IFAllocationSale.connect(ctx.casher).cash()
    mineNext()

    // cash again (expect to revert)
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cash()).to.be.revertedWith(ALREADY_CASHED)
    mineNext()

    // withdraw
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()

    // expect balance to increase by purchased amount
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')
    // expect balance to increase by cash amount
    expect(await ctx.PaymentToken.balanceOf(ctx.casher.address)).to.equal(paymentAmount)
  })

  it('does not under cash (if accidental sale token direct transfer in)', async function () {
    mineNext()

    // ctx.seller accidentally transfers in token directly
    await ctx.SaleToken.connect(ctx.seller).transfer(
      ctx.IFAllocationSale.address,
      '1000000000000000000' // 1e18
    )
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // cash
    await ctx.IFAllocationSale.connect(ctx.casher).cash()
    mineNext()

    // withdraw
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()

    // expect balance to increase by purchased amount
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')
    // expect contract balance to be 0 (no coins locked)
    expect(await ctx.SaleToken.balanceOf(ctx.IFAllocationSale.address)).to.equal('0')
  })

  it('can set linear vesting', async function () {
    await ctx.IFAllocationSale.connect(ctx.owner).setLinearVestingEndTime(ctx.linearVestingEndTime)
    mineNext()

    // amount to pay
    const paymentAmount = 333330

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      ethers.constants.MaxUint256,
    )
    await ctx.PaymentToken.connect(ctx.buyer2).approve(
      ctx.IFAllocationSale.address,
      paymentAmount * 2,
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount / 2)
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount / 2)
    await ctx.IFAllocationSale.connect(ctx.buyer2)['purchase(uint256)'](paymentAmount * 2)

    mineNext()

    // linear vesting: User makes a purchase and claim before vesting starts
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.reverted
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))
    // linear vesting: User makes a purchase and claim the tokens during vesting period
    minePause()
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineStart()
    mineNext()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('1')

    mineTimeDelta((ctx.linearVestingEndTime - ctx.endTime) / 3)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('11113')

    mineTimeDelta((ctx.linearVestingEndTime - ctx.endTime) / 3 * 2)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')

    // linear vesting: User makes a purchase and claim the tokens after vesting period
    await ctx.IFAllocationSale.connect(ctx.buyer2).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('66666')
  })

  it('can set cliff vesting', async function () {
    // amount to pay
    const paymentAmount = 333330
    const withdrawDelay = 10000

    const cliffInterval = Math.floor((ctx.linearVestingEndTime - ctx.endTime) / 3)
    const cliffPeriod = [
      ctx.endTime + withdrawDelay + 1,
      ctx.endTime + withdrawDelay + cliffInterval * 1,
      ctx.endTime + withdrawDelay + cliffInterval * 2,
      ctx.endTime + withdrawDelay + cliffInterval * 3
    ]
    const cliffPct = [10, 20, 30, 40]
    await ctx.IFAllocationSale.connect(ctx.owner).setWithdrawDelay(withdrawDelay)
    await ctx.IFAllocationSale.connect(ctx.owner).setCliffPeriod(cliffPeriod, cliffPct)

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))
    // purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)
    // cliff vesting: User makes a purchase and claim before cliff vesting starts
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_BEFORE_CLAIM)

    mineTimeDelta(ctx.endTime + withdrawDelay - (await getBlockTime()) + 1)

    // test withdraw
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('3333')

    // just before the second cliff time
    mineNext()
    mineTimeDelta((ctx.endTime + withdrawDelay + cliffInterval * 1) - (await getBlockTime()) - 2)
    // cliff vesting: User makes a purchase. Time pasts cliff 1. He makes claims.
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)

    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('9999')


    mineTimeDelta(cliffPeriod[3] - (await getBlockTime()))
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33333')
  })
  it('can limit access', async function () {
    const notOwner = [ctx.casher, ctx.seller, ctx.buyer, ctx.buyer2]
    const withdrawDelay = 10000
    const cliffInterval = Math.floor(ctx.linearVestingEndTime / 3)

    for (const user of notOwner) {
      await expect(ctx.IFAllocationSale.connect(user).setMinTotalPayment(0)).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).setCasher(ctx.owner.address)).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).setWhitelistSetter(ctx.owner.address)).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).setWithdrawDelay(3600)).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).setLinearVestingEndTime(ctx.linearVestingEndTime)).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).setCliffPeriod(
        [
          ctx.endTime + withdrawDelay + 1,
          ctx.endTime + withdrawDelay + cliffInterval * 1,
          ctx.endTime + withdrawDelay + cliffInterval * 2,
          ctx.endTime + withdrawDelay + cliffInterval * 3
        ],
        [10, 20, 30, 40]
      )).to.be.revertedWith(NOT_OWNER)
      await expect(ctx.IFAllocationSale.connect(user).emergencyTokenRetrieve(ctx.PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
      // IFMerkleAllocationSale doesn't have this function
      if (typeof ctx.IFAllocationSale.setSaleTokenAllocationOverride === 'function') {
        await expect(ctx.IFAllocationSale.connect(user).setSaleTokenAllocationOverride(0)).to.be.revertedWith(NOT_OWNER)
      }
    }
  })
  it("allows authorized users to cash payment tokens multiple times", async function () {
    // Sending some payment tokens to the IFAllocationSale contract to simulate earnings
    await ctx.PaymentToken.transfer(ctx.IFAllocationSale.address, ethers.utils.parseEther("100"));

    // Assuming ctx.IFAllocationSale has a function to cash payment tokens and ctx.casher is authorized
    let initialCasherBalance = await ctx.PaymentToken.balanceOf(ctx.casher.address);

    // Define an excessive cash amount that exceeds the contract's balance
    const excessiveCashAmount = ethers.utils.parseEther("150");

    // Attempting to cash more than the available balance by an authorized user should revert
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cashPaymentToken(excessiveCashAmount))
      .to.be.revertedWith(NOT_ENOUGH_PAYMENT_TOKEN_TO_CASH);

    // Authorized casher cashes payment tokens
    const validCashAmount = ethers.utils.parseEther("50");
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cashPaymentToken(validCashAmount))
      .to.emit(ctx.IFAllocationSale, "Cash")
      .withArgs(ctx.casher.address, validCashAmount, 0);

    // Validate casher's new balance
    let newCasherBalance = await ctx.PaymentToken.balanceOf(ctx.casher.address);
    // expect(newCasherBalance.sub(initialCasherBalance)).to.equal(validCashAmount);

    // Attempt to cash payment tokens by an unauthorized user (e.g., ctx.buyer) should fail
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).cashPaymentToken(validCashAmount))
      .to.be.revertedWith(NOT_CASHER_OR_OWNER); // Adjust the error message based on your contract's requirements

    // Second cashing operation
    initialCasherBalance = newCasherBalance; // Update the initial balance to the new balance for the next comparison
    await expect(ctx.IFAllocationSale.connect(ctx.casher).cashPaymentToken(validCashAmount))
      .to.emit(ctx.IFAllocationSale, "Cash")
      .withArgs(ctx.casher.address, validCashAmount, 0);

    // Check balance after second cashing
    newCasherBalance = await ctx.PaymentToken.balanceOf(ctx.casher.address);
    // expect(newCasherBalance.sub(initialCasherBalance)).to.equal(validCashAmount);

    // Attempt to cash payment tokens by an unauthorized user (e.g., ctx.buyer) should still fail
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).cashPaymentToken(validCashAmount))
      .to.be.revertedWith(NOT_CASHER_OR_OWNER); // Adjust the error message based on your contract's requirements

    // mineTimeDelta((ctx.endTime + ctx.withdrawDelay) - (await getBlockTime() + 1))
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // check the balance of payment token is 0
    expect(await ctx.PaymentToken.balanceOf(ctx.IFAllocationSale.address)).to.equal(0)

    expect(await ctx.SaleToken.balanceOf(ctx.casher.address)).to.equal(0)
    // can call cash()
    await ctx.IFAllocationSale.connect(ctx.casher).cash() 
    expect(await ctx.SaleToken.balanceOf(ctx.casher.address)).to.equal(ctx.fundAmount)
  });
  it("can pause purchase", async function () {
    mineNext()

    // amount to pay
    const paymentAmount = '333330'
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // pause purchase
    await ctx.IFAllocationSale.connect(ctx.owner).setIsPurchaseHalted(true)

    // test purchase (expect to revert)
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount))
      .to.be.revertedWith(PURCHASE_IS_HALTED)

    // unpause purchase
    await ctx.IFAllocationSale.connect(ctx.owner).setIsPurchaseHalted(false)
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)
  });
  it('can set integer purchase', async function () {
    const integerPaymentAmount = '5920'
    const nonIntegerPaymentAmount = '333331'
    await ctxSale.IFAllocationSale.connect(ctxSale.owner).setIsIntegerSale(true)
    // fast forward from current time to start time
    mineTimeDelta(ctxSale.startTime - (await getBlockTime()))
    await ctxSale.PaymentToken.connect(ctxSale.buyer).approve(
      ctxSale.IFAllocationSale.address,
      nonIntegerPaymentAmount
    )
    await ctxSale.IFAllocationSale.connect(ctxSale.buyer)['purchase(uint256)'](integerPaymentAmount)
    await expect(ctxSale.IFAllocationSale.connect(ctxSale.buyer)['purchase(uint256)'](nonIntegerPaymentAmount))
      .to.be.revertedWith(CAN_ONLY_BUY_INTEGER_AMOUNT)

    // fast forward from current time to after end time
    mineTimeDelta(ctxSale.endTime - (await getBlockTime()))
    // test withdraw
    mineNext()
    await ctxSale.IFAllocationSale.connect(ctxSale.buyer).withdraw()
    mineNext()

    const expectedBalance = (parseFloat(integerPaymentAmount) / ctxSale.paymentTokenPerSaleToken).toString()
    expect(await ctxSale.SaleToken.balanceOf(ctxSale.buyer.address)).to.equal(expectedBalance)
  })
  it('cannot renounce ownership', async function () {
    await expect(ctxSale.IFAllocationSale.connect(ctxSale.owner).renounceOwnership())
      .to.be.revertedWith('ownership renunciation is disabled')
  })
}
