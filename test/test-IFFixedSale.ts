import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { expect } from 'chai'
import { Contract } from 'ethers'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex, pad } from '../library/merkleWhitelist'
import { getBlockTime, mineNext, minePause, mineStart, mineTimeDelta, setAutomine } from './helpers'
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { _ctx } from './IFAllocationSaleGeneralTest'
import { EXCEED_MAX_PAYMENT, NO_TOKEN_TO_BE_WITHDRAWN, NOT_A_GIVEAWAY, ALREADY_CASHED, CANNOT_WITHDRAW_YET, NOT_CASHER_OR_OWNER, NOT_FUNDER, NOT_OWNER, USE_WITHDRAWGIVEAWAY } from './reverts/msg-IFAllocationSale'

function computeMerkleRootWithAllocation(signers: SignerWithAddress[], allocations: number[]): [string[], Map<string, string>]{
    const leaves: string[] = []
    const addressValMap = new Map()
    signers.forEach((s: SignerWithAddress, i: number) => {
        const amount = allocations[i].toString()
        const packed = ethers.utils.solidityPack(
          ['address', 'uint256'],
          [s.address.toLowerCase(), amount],
        )
        leaves.push(packed)
        addressValMap.set(s.address.toLowerCase(), packed)
      }
    )
    leaves.sort()
    return [leaves, addressValMap]
}

export default describe('IF Fixed Sale', function () {
  const contractName = 'MockIFFixedSale'
  // unset timeout from the test
  this.timeout(800000)

  // deployer address
  let owner: SignerWithAddress
  let buyer: SignerWithAddress
  let buyer2: SignerWithAddress
  let seller: SignerWithAddress
  let casher: SignerWithAddress

  // contract vars
  let PaymentToken: Contract
  let SaleToken: Contract
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
    const StakeToken = await TestTokenFactory.connect(buyer).deploy(
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
    const IFAllocationMaster = await IFAllocationMasterFactory.deploy(
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
      contractName
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      salePrice,
      seller.address,
      PaymentToken.address,
      SaleToken.address,
      trackId,
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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount)

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

  it('can perform a zero price giveaway sale (unwhitelisted / first come first serve)', async function () {
    mineNext()

    // here set up a new IFAllocationSale with salePrice of 0, because
    // provided fixture sale does not have salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      contractName
    )
    IFAllocationSale = await IFAllocationSaleFactory.deploy(
      0, // sale price
      seller.address,
      PaymentToken.address, // doesn't matter
      SaleToken.address,
      trackId, // doesn't matter
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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount)

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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount)

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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount)

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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount / 2)
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount / 2)
    await IFAllocationSale.connect(buyer2)['purchase(uint256)'](paymentAmount * 2)

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
    await IFAllocationSale.connect(buyer)['purchase(uint256)'](paymentAmount)
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
      // IFMerkleAllocationSale doesn't have this function
      if (typeof IFAllocationSale.setSaleTokenAllocationOverride === 'function') {
        await expect(IFAllocationSale.connect(user).setSaleTokenAllocationOverride(0)).to.be.revertedWith(NOT_OWNER)
      }
    }
  })


  it('can save allocation amount in merkle tree', async function () {
    const signers = await ethers.getSigners()
    const allocations = Array(signers.length).fill(1)
    const [leaves, addressValMap] = computeMerkleRootWithAllocation(signers, allocations)
    const merkleRoot = computeMerkleRoot(leaves)
    await IFAllocationSale.connect(owner).setWhitelist(merkleRoot)
    mineNext()

    const tempAcct = (await ethers.getSigners())[0]
    const packed = addressValMap.get(tempAcct.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    expect(
      await IFAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        1,
      )
    ).to.equal(true)
    expect(
      await IFAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        200,
      )
    ).to.equal(false)
  })

  it('can override sale token allocations (test preventing exceeding allocation)', async function () {

    const allocationAmount = 10000
    const paymentAmount = 100001
    const [leaves, addressValMap] = computeMerkleRootWithAllocation([buyer], [allocationAmount])
    // amount to pay (should fail, because this is 1 over allocation)

    // set sale token allocation override
    await IFAllocationSale.connect(owner).setWhitelist(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )

    const packed = addressValMap.get(buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await expect(IFAllocationSale.connect(buyer)['purchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      allocationAmount,
    )).to.be.revertedWith(EXCEED_MAX_PAYMENT)

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

    const allocationAmount = 5000
    // amount to pay for each claimer (should go through since this is exactly how much allocation they have)
    const paymentAmount = 50000

    const [leaves, addressValMap] = computeMerkleRootWithAllocation([buyer, buyer2], [allocationAmount, allocationAmount])
    await IFAllocationSale.connect(owner).setWhitelist(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(startTime - (await getBlockTime()))

    // test purchase for buyers 1 and 2
    mineNext()
    await PaymentToken.connect(buyer).approve(
      IFAllocationSale.address,
      paymentAmount
    )
    const packed = addressValMap.get(buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await IFAllocationSale.connect(buyer)['purchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      allocationAmount,
    )

    mineNext()
    await PaymentToken.connect(buyer2).approve(
      IFAllocationSale.address,
      paymentAmount
    )

    const packed2 = addressValMap.get(buyer2.address.toLowerCase()) || ''
    const tempAcctIdx2 = getAddressIndex(leaves, packed2)
    await IFAllocationSale.connect(buyer2)['purchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx2),
      allocationAmount,
    )

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([], allocationAmount)).to.be.revertedWith(NOT_A_GIVEAWAY)
    await IFAllocationSale.connect(buyer).withdraw()
    mineNext()
    await IFAllocationSale.connect(buyer2).withdraw()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFAllocationSale.connect(buyer).withdrawGiveaway([], allocationAmount)).to.be.revertedWith(NOT_A_GIVEAWAY)
    mineNext()

    // expect balance to be 5000 for both buyers
    expect(await SaleToken.balanceOf(buyer.address)).to.equal('5000')
    expect(await SaleToken.balanceOf(buyer2.address)).to.equal('5000')

    // test purchaser counter
    expect(await IFAllocationSale.purchaserCount()).to.equal(2)

    // test withdrawer counter
    expect(await IFAllocationSale.withdrawerCount()).to.equal(2)
  })
})