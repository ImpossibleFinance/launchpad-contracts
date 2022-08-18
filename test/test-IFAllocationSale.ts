import '@nomiclabs/hardhat-ethers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import {
  getBlockTime,
  getGasUsed,
  mineNext,
  mineTimeDelta,
  minePause,
  mineStart,
  setAutomine,
} from './helpers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { Contract } from '@ethersproject/contracts'
import {
  computeMerkleRoot,
  computeMerkleProof,
  getAddressIndex,
} from '../library/merkleWhitelist'
import { BigNumber } from 'ethers'
import { ALREADY_CASHED, NO_TOKEN_TO_BE_WITHDRAWN, CANNOT_WITHDRAW_YET, EXCEED_MAX_PAYMENT, NOT_CASHER_OR_OWNER, NOT_OWNER, NOT_WHITELIST_SETTER_OR_OWNER, NOT_A_GIVEAWAY, NOT_FUNDER, USE_WITHDRAWGIVEAWAY, PROOF_INVALID, SALE_IS_STARTED } from './reverts/msg-IFAllocationSale'

export default describe('IF Allocation Sale', function () {
  // unset timeout from the test
  this.timeout(0)

  // deployer address
  let owner: SignerWithAddress
  let buyer: SignerWithAddress
  let buyer2: SignerWithAddress
  let seller: SignerWithAddress
  let casher: SignerWithAddress

  // contract vars
  let StakeToken: Contract
  let PaymentToken: Contract
  let SaleToken: Contract
  let IFAllocationMaster: Contract
  let IFAllocationSale: Contract

  // allocation master vars
  let trackId: number

  // sale contract vars
  let snapshotTimestamp: number // block at which to take allocation snapshot
  let startTime: number // start timestamp of sale (inclusive)
  let endTime: number // end timestamp of sale (inclusive)
  let vestingEndTime: number // end timestamp of vesting
  const salePrice = '10000000000000000000' // 10 PAY per SALE
  const maxTotalDeposit = '25000000000000000000000000' // max deposit
  // other vars
  // const fundAmount = '33333'
  const fundAmount = '1000000000'

  this.beforeAll(async () => {
    await setAutomine(false)
  })

  this.afterAll(async () => {
    await setAutomine(true)
  })

  // setup for each test
  beforeEach(async () => {
    // set launchpad blocks in future
    mineNext()
    const currTime = await getBlockTime()
    mineNext()
    snapshotTimestamp = currTime + 5000
    startTime = currTime + 10000
    endTime = currTime + 20000
    vestingEndTime = currTime + 50000

    // get test accounts
    owner = (await ethers.getSigners())[0]
    buyer = (await ethers.getSigners())[1]
    seller = (await ethers.getSigners())[2]
    casher = (await ethers.getSigners())[3]
    buyer2 = (await ethers.getSigners())[4]

    // deploy test tokens
    const TestTokenFactory = await ethers.getContractFactory('GenericToken')
    StakeToken = await TestTokenFactory.connect(buyer).deploy(
      'Test Stake Token',
      'STAKE',
      '21000000000000000000000000' // 21 million * 10**18
    )
    PaymentToken = await TestTokenFactory.connect(buyer).deploy(
      'Test Payment Token',
      'PAY',
      '21000000000000000000000000' // 21 million * 10**18
    )
    SaleToken = await TestTokenFactory.connect(seller).deploy(
      'Test Sale Token',
      'SALE',
      '21000000000000000000000000' // 21 million * 10**18
    )

    // redistribute tokens
    mineNext()
    StakeToken.connect(buyer).transfer(
      buyer2.address,
      '1000000000000000000000000'
    )
    PaymentToken.connect(buyer).transfer(
      buyer2.address,
      '1000000000000000000000000'
    )

    // deploy allocation master
    const IFAllocationMasterFactory = await ethers.getContractFactory(
      'IFAllocationMaster'
    )
    IFAllocationMaster = await IFAllocationMasterFactory.deploy(
      ethers.constants.AddressZero
    )

    // add track on allocation master
    mineNext()
    mineNext()
    await IFAllocationMaster.addTrack(
      'IDIA track', // name
      StakeToken.address, // stake token
      10000, // weight accrual rate
      '100000000000000000', // passive rollover rate (10%)
      '200000000000000000', // active rollover rate (20%)
      '1000000000000000000000000000000' // max total stake (1 trillion)
    )

    // get new track id
    mineNext()
    trackId = (await IFAllocationMaster.trackCount()) - 1

    // deploy sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      salePrice,
      seller.address,
      PaymentToken.address,
      SaleToken.address,
      IFAllocationMaster.address,
      trackId,
      snapshotTimestamp,
      startTime,
      endTime,
      maxTotalDeposit
    )
    mineNext()

    // set the casher address
    await IFAllocationSale.setCasher(casher.address)
    mineNext()

    // fund sale
    mineNext()
    await SaleToken.connect(seller).approve(
      IFAllocationSale.address,
      fundAmount
    ) // approve
    await IFAllocationSale.connect(seller).fund(fundAmount) // fund
    //
    // stake and accrue stake weight
    mineNext()
    const stakeAmount = 100000000000000
    minePause()
    // buyer 1
    await StakeToken.connect(buyer).approve(
      IFAllocationMaster.address,
      3 * stakeAmount
    ) // approve
    await IFAllocationMaster.connect(buyer).stake(trackId, 3 * stakeAmount) // stake
    // buyer 2
    await StakeToken.connect(buyer2).approve(
      IFAllocationMaster.address,
      stakeAmount
    ) // approve
    await IFAllocationMaster.connect(buyer2).stake(trackId, stakeAmount) // stake
    mineStart()

    // expect staked amount to match
    mineNext()
    expect(
      (await StakeToken.balanceOf(IFAllocationMaster.address)).toString()
    ).to.equal((stakeAmount * 4).toString())

    //fastforward from current block to after snapshot block
    mineTimeDelta(snapshotTimestamp - (await getBlockTime()))
  })

  it('can purchase, withdraw, and cash', async function () {
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)

    // Failover mechanism: Call emergencyTokenRetrieve while token is sale or payment token
    await expect(IFAllocationSale.connect(owner).emergencyTokenRetrieve(PaymentToken.address)).to.be.reverted

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await IFAllocationSale.connect(buyer).withdraw()
    mineNext()

    // expect balance to increase by fund amount
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')

    // test repeated withdraw (should fail)
    mineNext()
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)
    mineNext()

    // expect balance to remain the same
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')

    // test cash
    await IFAllocationSale.connect(casher).cash()
    // access control: only casher can cash
    await expect(IFAllocationSale.connect(buyer).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(IFAllocationSale.connect(seller).cash()).to.be.revertedWith(NOT_CASHER_OR_OWNER)
    await expect(IFAllocationSale.connect(owner).cash()).to.be.revertedWith(ALREADY_CASHED)
    mineNext()

    // expect balance to increase by cash amount
    expect(await PaymentToken.balanceOf(casher.address)).to.equal(paymentAmount)

    // test purchaser counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(1)

    // Failover mechanism: Call emergencyTokenRetrieve while token is sale or payment token
    await expect(IFAllocationSale.connect(casher).emergencyTokenRetrieve(PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    await expect(IFAllocationSale.connect(seller).emergencyTokenRetrieve(PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    await expect(IFAllocationSale.connect(buyer).emergencyTokenRetrieve(PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    IFAllocationSale.connect(owner).emergencyTokenRetrieve(PaymentToken.address)
  })

  it('can whitelist purchase', async function () {
    mineNext()

    // whitelisted addresses (sorted)
    const addresses = (await ethers.getSigners())
      .map((s) => s.address.toLowerCase())
      .sort()

    // get merkle root
    const merkleRoot = computeMerkleRoot(addresses)

    // add whitelist merkleroot to sale
    await IFAllocationSale.setWhitelist(merkleRoot)
    // access control: only owner or whitelistSetter can set white list
    await IFAllocationSale.setWhitelistSetter(seller.address)
    await IFAllocationSale.connect(seller).setWhitelist(merkleRoot)
    await expect(IFAllocationSale.connect(buyer).setWhitelist(merkleRoot)).to.be.revertedWith(NOT_WHITELIST_SETTER_OR_OWNER)
    await expect(IFAllocationSale.connect(casher).setWhitelist(merkleRoot)).to.be.revertedWith(NOT_WHITELIST_SETTER_OR_OWNER)
    mineNext()

    // test checking all whitelist accounts
    for (let i = 0; i < addresses.length; i++) {
      const tempAcct = (await ethers.getSigners())[i]
      const tempAcctIdx = getAddressIndex(addresses, tempAcct.address)
      try {
        expect(
          await IFAllocationSale.connect(tempAcct).checkWhitelist(
            tempAcct.address,
            computeMerkleProof(addresses, tempAcctIdx)
          )
        ).to.equal(true)
      } catch (e) {
        console.log(e)
      }
    }

    const account = buyer
    const acctIdx = getAddressIndex(addresses, account.address)

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test whitelist purchase
    mineNext()
    await PaymentToken.connect(account).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(account).whitelistedPurchase(
      paymentAmount,
      computeMerkleProof(addresses, acctIdx)
    )

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await IFAllocationSale.connect(account).withdraw()
    mineNext()

    // expect balance to increase by fund amount
    expect(await SaleToken.balanceOf(account.address)).to.equal('33333')

    // test purchaser counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(1)
  })

  it('can override sale token allocations (test preventing exceeding allocation)', async function () {
    mineNext()

    // amount to pay (should fail, because this is 1 over allocation)
    const paymentAmount = '100001'

    // set sale token allocation override
    await IFAllocationSale.setSaleTokenAllocationOverride(10000)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await expect(IFAllocationSale.connect(buyer).purchase(paymentAmount)).to.be.revertedWith(EXCEED_MAX_PAYMENT)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)
    mineNext()

    // expect balance to be 0
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')
  })

  it('can override sale token allocations (test multiple buyers)', async function () {
    mineNext()

    // amount to pay for each claimer (should go through since this is exactly how much allocation they have)
    const paymentAmount = '50000'

    // set sale token allocation override
    await IFAllocationSale.setSaleTokenAllocationOverride(5000)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase for buyers 1 and 2
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)

    mineNext()
    await PaymentToken.connect(buyer2).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer2).purchase(paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    await IFAllocationSale.connect(buyer).withdraw()
    mineNext()
    await IFAllocationSale.connect(buyer2).withdraw()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    mineNext()

    // expect balance to be 5000 for both buyers
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('5000')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('5000')

    // test purchaser counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(2)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  it('can perform a zero price giveaway sale (unwhitelisted / first come first serve)', async function () {
    mineNext()

    // here set up a new IFAllocationSale with salePrice of 0, because
    // provided fixture sale does not have salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      seller.address,
      PaymentToken.address, // doesn't matter
      SaleToken.address,
      IFAllocationMaster.address, // doesn't matter
      trackId, // doesn't matter
      snapshotTimestamp, // doesn't matter
      startTime, // doesn't matter
      endTime, // doesn't matter
      maxTotalDeposit // doesn't matter
    )
    mineNext()

    // fund sale
    mineNext()
    await SaleToken.connect(seller).approve(
      IFAllocationSale.address,
      fundAmount
    ) // approve
    await IFAllocationSale.connect(seller).fund(fundAmount) // fund
    // access control: Address other than funder calls fund
    await expect(IFAllocationSale.connect(casher).fund(fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund
    await expect(IFAllocationSale.connect(buyer).fund(fundAmount)).to.be.revertedWith(NOT_FUNDER) // fund

    // set sale token allocation override (flat amount every participant receives)
    await IFAllocationSale.setSaleTokenAllocationOverride(5000)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test normal withdraw (should not go through, must go through withdrawGiveaway)
    // access control: Withdraw when sale price is 0
    mineNext()
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()
    await expect(IFAllocationSale.connect(buyer2).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()

    // expect balance to be 0 for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('0')

    // test withdrawGiveaway (should go through)
    mineNext()
    await IFAllocationSale.connect(buyer).withdrawGiveaway([])
    mineNext()
    await IFAllocationSale.connect(buyer2).withdrawGiveaway([])
    mineNext()

    // expect balance to be 5000 for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('5000')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('5000')

    // test purchaser counter (should be 0! nothing purchased in 0 price sales)
    // note: this is the only scenario where this is different from withdrawer counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(0)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  it('can perform a zero price giveaway sale (whitelisted)', async function () {
    mineNext()

    // here set up a new IFAllocationSale with salePrice of 0, because
    // provided fixture sale does not have salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      seller.address,
      PaymentToken.address, // doesn't matter
      SaleToken.address,
      IFAllocationMaster.address, // doesn't matter
      trackId, // doesn't matter
      snapshotTimestamp, // doesn't matter
      startTime, // doesn't matter
      endTime, // doesn't matter
      maxTotalDeposit // doesn't matter
    )
    mineNext()

    // fund sale
    mineNext()
    await SaleToken.connect(seller).approve(
      IFAllocationSale.address,
      fundAmount
    ) // approve
    await IFAllocationSale.connect(seller).fund(fundAmount) // fund

    // set sale token allocation override (flat amount every participant receives)
    await IFAllocationSale.setSaleTokenAllocationOverride(5000)
    mineNext()

    // whitelisted addresses (sorted)
    const addresses = (await ethers.getSigners())
      .map((s) => s.address.toLowerCase())
      .sort()

    // get merkle root
    const merkleRoot = computeMerkleRoot(addresses)

    // add whitelist merkleroot to sale
    await IFAllocationSale.setWhitelist(merkleRoot)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdrawGiveaway without proof (should not go through)
    mineNext()
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([])).to.be.revertedWith(PROOF_INVALID)
    mineNext()
    await expect(IFAllocationSale.connect(buyer2).withdrawGiveaway([])).to.be.revertedWith(PROOF_INVALID)
    mineNext()

    // expect balance to be 0 for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('0')

    // test withdrawGiveaway with proof (should go through)
    mineNext()

    await IFAllocationSale.connect(buyer).withdrawGiveaway(
      computeMerkleProof(addresses, getAddressIndex(addresses, buyer.address))
    )
    mineNext()
    await IFAllocationSale.connect(buyer2).withdrawGiveaway(
      computeMerkleProof(addresses, getAddressIndex(addresses, buyer2.address))
    )
    mineNext()

    // expect balance to be 5000 for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('5000')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('5000')
  })

  it('can perform a zero price giveaway sale unwhitelisted with staked amount', async function () {
    mineNext()

    // here set up a new IFAllocationSale with salePrice of 0, because
    // provided fixture sale does not have salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      seller.address,
      PaymentToken.address, // doesn't matter
      SaleToken.address,
      IFAllocationMaster.address, // doesn't matter
      trackId, // doesn't matter
      snapshotTimestamp, // doesn't matter
      startTime, // doesn't matter
      endTime, // doesn't matter
      maxTotalDeposit // doesn't matter
    )
    mineNext()

    // fund sale
    mineNext()
    await SaleToken.connect(seller).approve(
      IFAllocationSale.address,
      fundAmount
    ) // approve
    await IFAllocationSale.connect(seller).fund(fundAmount) // fund

    // no need to set override, because skaked drop
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    mineNext()

    // test normal withdraw (should not go through, must go through withdrawGiveaway)
    mineNext()
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()
    await expect(IFAllocationSale.connect(buyer2).withdraw()).to.be.revertedWith(USE_WITHDRAWGIVEAWAY)
    mineNext()

    // expect balance to be 0 for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('0')

    // test withdrawGiveaway (should go through)
    mineNext()
    await IFAllocationSale.connect(buyer).withdrawGiveaway([])
    mineNext()
    await IFAllocationSale.connect(buyer2).withdrawGiveaway([])
    mineNext()

    // expect balance to be 1:3 ratio for both participants
    expect(await SaleToken.balanceOf(buyer.address)).to.equal(
      (Number(fundAmount) * 3) / 4
    )
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal(
      (Number(fundAmount) * 1) / 4
    )

    // test purchaser counter (should be 0! nothing purchased in 0 price sales)
    // note: this is the only scenario where this is different from withdrawer counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(0)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  it('can set withdraw delay', async function () {
    mineNext()

    // delay of 10 blocks
    const delay = 10

    // add withdraw delay
    await IFAllocationSale.setWithdrawDelay(delay)
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw and cash (should fail because need 1 more block)
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_YET)
    // access control: Call cash before endTime + withdrawDelay
    await expect(IFAllocationSale.connect(casher).cash())

    mineNext()

    // fails
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')
    // fails
    expect(await PaymentToken.balanceOf(casher.address)).to.equal('0')

    // simulate `delay` time passing
    mineTimeDelta(delay)

    // test withdraw and cash (should work here after delay passed)
    await IFAllocationSale.connect(buyer).withdraw()
    await IFAllocationSale.connect(casher).cash()
    // access control: Call cash after endTime + withdrawDelay twice
    await expect(IFAllocationSale.connect(casher).cash()).to.be.revertedWith(ALREADY_CASHED)

    mineNext()

    // expect balance to increase by fund amount
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')
    // expect balance to increase by cash amount
    expect(await PaymentToken.balanceOf(casher.address)).to.equal(paymentAmount)

    // test purchaser counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(1)
  })

  it('does not over cash', async function () {
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // cash first (testing that we do not over-remove sale token)
    await IFAllocationSale.connect(casher).cash()
    mineNext()

    // cash again (expect to revert)
    await expect(IFAllocationSale.connect(casher).cash()).to.be.revertedWith(ALREADY_CASHED)
    mineNext()

    // withdraw
    await IFAllocationSale.connect(buyer).withdraw()
    mineNext()

    // expect balance to increase by purchased amount
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')
    // expect balance to increase by cash amount
    expect(await PaymentToken.balanceOf(casher.address)).to.equal(paymentAmount)
  })

  it('does not under cash (if accidental sale token direct transfer in)', async function () {
    mineNext()

    // seller accidentally transfers in token directly
    await SaleToken.connect(seller).transfer(
      IFAllocationSale.address,
      '1000000000000000000' // 1e18
    )
    mineNext()

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // cash
    await IFAllocationSale.connect(casher).cash()
    mineNext()

    // withdraw
    await IFAllocationSale.connect(buyer).withdraw()
    mineNext()

    // expect balance to increase by purchased amount
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')
    // expect contract balance to be 0 (no coins locked)
    expect(await SaleToken.balanceOf(IFAllocationSale.address)).to.equal('0')
  })

  it('can set linear vesting', async function () {
    await IFAllocationSale.connect(owner).setVestingEndTime(vestingEndTime)
    mineNext()

    // amount to pay
    const paymentAmount = 333330

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      ethers.constants.MaxUint256,
    )
    await PaymentToken.connect(buyer2).approve(
      IFAllocationSale.address,
      paymentAmount * 2,
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount / 2)
    await IFAllocationSale.connect(buyer).purchase(paymentAmount / 2)
    await IFAllocationSale.connect(buyer2).purchase(paymentAmount * 2)
    const maxPayment: BigNumber = await IFAllocationSale.getMaxPayment(buyer.address)
    // linear vesting: User has a purchase cap of x. He tried to buy x +1.
    await expect(IFAllocationSale.connect(buyer).purchase(maxPayment.add(1))).to.be.revertedWith(EXCEED_MAX_PAYMENT)

    mineNext()

    // linear vesting: User makes a purchase and claim before vesting starts
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_YET)
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('0')

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))
    // linear vesting: User makes a purchase and claim the tokens during vesting period
    minePause()
    await IFAllocationSale.connect(buyer).withdraw()
    await IFAllocationSale.connect(buyer).withdraw()
    mineStart()
    mineNext()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('1')

    mineTimeDelta((vestingEndTime - endTime) / 3)
    await IFAllocationSale.connect(buyer).withdraw()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('11113')

    mineTimeDelta((vestingEndTime - endTime) / 3 * 2)
    await IFAllocationSale.connect(buyer).withdraw()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')

    // linear vesting: User makes a purchase and claim the tokens after vesting period
    await IFAllocationSale.connect(buyer2).withdraw()
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('66666')
  })

  it('can vest with withdrawGiveaway', async function () {
    const withdrawDelay = 10000
    mineNext()
    // here set up a new IFAllocationSale with salePrice of 0, because
    // provided fixture sale does not have salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      'IFAllocationSale'
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      seller.address,
      PaymentToken.address, // doesn't matter
      SaleToken.address,
      IFAllocationMaster.address, // doesn't matter
      trackId, // doesn't matter
      snapshotTimestamp, // doesn't matter
      startTime, // doesn't matter
      endTime, // doesn't matter
      maxTotalDeposit // doesn't matter
    )
    // fund sale
    await SaleToken.connect(seller).approve(
      IFAllocationSale.address,
      fundAmount
    ) // approve
    await IFAllocationSale.connect(seller).fund(fundAmount) // fund
    await IFAllocationSale.connect(owner).setWithdrawDelay(withdrawDelay)

    // set sale token allocation override (flat amount every participant receives)
    await IFAllocationSale.setSaleTokenAllocationOverride(33330)

    IFAllocationSale.connect(owner).setVestingEndTime(vestingEndTime)

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // set withdrawal delay
    await expect(IFAllocationSale.connect(owner).setWithdrawDelay(withdrawDelay)).to.be.revertedWith(SALE_IS_STARTED)
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([])).to.be.revertedWith(CANNOT_WITHDRAW_YET)

    mineTimeDelta(endTime + withdrawDelay - (await getBlockTime()))
    await IFAllocationSale.connect(buyer).withdrawGiveaway([])
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('1')

    mineTimeDelta(vestingEndTime - endTime)
    await IFAllocationSale.connect(buyer).withdrawGiveaway([])
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33330')
  })

  it('can set cliff vesting', async function () {
    // amount to pay
    const paymentAmount = 333330
    const withdrawDelay = 10000

    const cliffInterval = Math.floor((vestingEndTime - endTime) / 3)
    const cliffPeriod = [
      endTime + withdrawDelay + 1,
      endTime + withdrawDelay + cliffInterval * 1,
      endTime + withdrawDelay + cliffInterval * 2,
      endTime + withdrawDelay + cliffInterval * 3
    ]
    const cliffPct = [10, 20, 30, 40]
    await IFAllocationSale.connect(owner).setWithdrawDelay(withdrawDelay)
    await IFAllocationSale.connect(owner).setCliffPeriod(cliffPeriod, cliffPct)

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))
    // purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    await IFAllocationSale.connect(buyer).purchase(paymentAmount)
    // cliff vesting: User makes a purchase and claim before cliff vesting starts
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(CANNOT_WITHDRAW_YET)

    mineTimeDelta(endTime + withdrawDelay - (await getBlockTime()) + 1)

    // test withdraw
    await IFAllocationSale.connect(buyer).withdraw()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('3333')

    // just before the second cliff time
    mineNext()
    mineTimeDelta((endTime + withdrawDelay + cliffInterval * 1) - (await getBlockTime()) - 2)
    // cliff vesting: User makes a purchase. Time pasts cliff 1. He makes claims.
    await expect(IFAllocationSale.connect(buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)

    mineNext()
    await IFAllocationSale.connect(buyer).withdraw()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('9999')


    mineTimeDelta(cliffPeriod[3] - (await getBlockTime()))
    await IFAllocationSale.connect(buyer).withdraw()
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('33333')
  })
  it('can limit access', async function () {
    const notOwner = [casher, seller, buyer, buyer2]
    const withdrawDelay = 10000
    const cliffInterval = Math.floor(vestingEndTime / 3)

    for (const user of notOwner) {
      await expect(IFAllocationSale.connect(user).setMinTotalPayment(0)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setSaleTokenAllocationOverride(0)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setCasher(owner.address)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setWhitelistSetter(owner.address)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setWithdrawDelay(3600)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setVestingEndTime(vestingEndTime)).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).setCliffPeriod(
        [
          endTime + withdrawDelay + 1,
          endTime + withdrawDelay + cliffInterval * 1,
          endTime + withdrawDelay + cliffInterval * 2,
          endTime + withdrawDelay + cliffInterval * 3
        ],
        [10, 20, 30, 40]
      )).to.be.revertedWith(NOT_OWNER)
      await expect(IFAllocationSale.connect(user).emergencyTokenRetrieve(PaymentToken.address)).to.be.revertedWith(NOT_OWNER)
    }
  })
})
