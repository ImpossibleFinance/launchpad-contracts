import '@nomiclabs/hardhat-ethers'
import { expect } from 'chai'
import { BigNumber } from 'ethers'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex } from '../library/merkleWhitelist'
import { getBlockTime, mineNext, mineTimeDelta } from './helpers'
import IFAllocationSaleGeneralTest, { _ctx } from './IFAllocationSaleGeneralTest'
import { CANNOT_WITHDRAW_YET, EXCEED_MAX_PAYMENT, NOT_A_GIVEAWAY, NOT_WHITELIST_SETTER_OR_OWNER, NO_TOKEN_TO_BE_WITHDRAWN, PROOF_INVALID, SALE_IS_STARTED, USE_WITHDRAWGIVEAWAY } from './reverts/msg-IFAllocationSale'

export default describe('IF Allocation Sale', function () {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const ctx: any = _ctx

  const contractName = 'IFAllocationSale'

  const generalTest = IFAllocationSaleGeneralTest
  generalTest(this, contractName, ctx)

  generalTest.prototype.it = it('can whitelist purchase', async function () {
    mineNext()

    // whitelisted addresses (sorted)
    const addresses = (await ethers.getSigners())
      .map((s) => s.address.toLowerCase())
      .sort()

    // get merkle root
    const merkleRoot = computeMerkleRoot(addresses)

    // add whitelist merkleroot to sale
    await ctx.IFAllocationSale.setWhitelist(merkleRoot)
    // access control: only ctx.owner or whitelistSetter can set white list
    await ctx.IFAllocationSale.setWhitelistSetter(ctx.seller.address)
    await ctx.IFAllocationSale.connect(ctx.seller).setWhitelist(merkleRoot)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).setWhitelist(merkleRoot)).to.be.revertedWith(NOT_WHITELIST_SETTER_OR_OWNER)
    await expect(ctx.IFAllocationSale.connect(ctx.casher).setWhitelist(merkleRoot)).to.be.revertedWith(NOT_WHITELIST_SETTER_OR_OWNER)
    mineNext()

    // test checking all whitelist accounts
    for (let i = 0; i < addresses.length; i++) {
      const tempAcct = (await ethers.getSigners())[i]
      const tempAcctIdx = getAddressIndex(addresses, tempAcct.address)
      try {
        expect(
          await ctx.IFAllocationSale.connect(tempAcct).checkWhitelist(
            tempAcct.address,
            computeMerkleProof(addresses, tempAcctIdx)
          )
        ).to.equal(true)
      } catch (e) {
        console.log(e)
      }
    }

    const account = ctx.buyer
    const acctIdx = getAddressIndex(addresses, account.address)

    // amount to pay
    const paymentAmount = '333330'

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test whitelist purchase
    mineNext()
    await ctx.PaymentToken.connect(account).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(account).whitelistedPurchase(
      paymentAmount,
      computeMerkleProof(addresses, acctIdx)
    )

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await ctx.IFAllocationSale.connect(account).withdraw()
    mineNext()

    // expect balance to increase by fund amount
    expect(await ctx.SaleToken.balanceOf(account.address)).to.equal('33333')

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(1)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(1)
  })

  generalTest.prototype.it = it('can perform a zero price giveaway sale unwhitelisted with staked amount', async function () {
    mineNext()

    // here set up a new ctx.IFAllocationSale with ctx.salePrice of 0, because
    // provided fixture sale does not have ctx.salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      contractName
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

    // no need to set override, because skaked drop
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    mineNext()

    // test normal withdraw (should not go through, must go through withdrawGiveaway)
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

    // expect balance to be 1:3 ratio for both participants
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal(
      (Number(ctx.fundAmount) * 3) / 4
    )
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal(
      (Number(ctx.fundAmount) * 1) / 4
    )

    // test purchaser counter (should be 0! nothing purchased in 0 price sales)
    // note: _this is the only scenario where _this is different from withdrawer counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(0)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  generalTest.prototype.it = it('can override sale token allocations (test preventing exceeding allocation)', async function () {
    mineNext()

    // amount to pay (should fail, because _this is 1 over allocation)
    const paymentAmount = '100001'

    // set sale token allocation override
    await ctx.IFAllocationSale.setSaleTokenAllocationOverride(10000)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)).to.be.revertedWith(EXCEED_MAX_PAYMENT)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)
    mineNext()

    // expect balance to be 0
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
  })

  generalTest.prototype.it = it('can override sale token allocations (test multiple buyers)', async function () {
    mineNext()

    // amount to pay for each claimer (should go through since _this is exactly how much allocation they have)
    const paymentAmount = '50000'

    // set sale token allocation override
    await ctx.IFAllocationSale.setSaleTokenAllocationOverride(5000)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase for buyers 1 and 2
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](paymentAmount)

    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer2).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    await ctx.IFAllocationSale.connect(ctx.buyer2)['purchase(uint256)'](paymentAmount)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer2).withdraw()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    mineNext()

    // expect balance to be 5000 for both buyers
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('5000')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('5000')

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(2)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(2)
  })

  generalTest.prototype.it = it('can perform a zero price giveaway sale (whitelisted)', async function () {
    mineNext()

    // here set up a new ctx.IFAllocationSale with ctx.salePrice of 0, because
    // provided fixture sale does not have ctx.salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      contractName
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

    // set sale token allocation override (flat amount every participant receives)
    await ctx.IFAllocationSale.setSaleTokenAllocationOverride(5000)
    mineNext()

    // whitelisted addresses (sorted)
    const addresses = (await ethers.getSigners())
      .map((s) => s.address.toLowerCase())
      .sort()

    // get merkle root
    const merkleRoot = computeMerkleRoot(addresses)

    // add whitelist merkleroot to sale
    await ctx.IFAllocationSale.setWhitelist(merkleRoot)
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // nothing to do here

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdrawGiveaway without proof (should not go through)
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(PROOF_INVALID)
    mineNext()
    await expect(ctx.IFAllocationSale.connect(ctx.buyer2).withdrawGiveaway([])).to.be.revertedWith(PROOF_INVALID)
    mineNext()

    // expect balance to be 0 for both participants
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('0')

    // test withdrawGiveaway with proof (should go through)
    mineNext()

    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway(
      computeMerkleProof(addresses, getAddressIndex(addresses, ctx.buyer.address))
    )
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer2).withdrawGiveaway(
      computeMerkleProof(addresses, getAddressIndex(addresses, ctx.buyer2.address))
    )
    mineNext()

    // expect balance to be 5000 for both participants
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('5000')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('5000')
  })

  generalTest.prototype.it = it('can limit payment amount', async function () {
    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))
    const maxPayment: BigNumber = await ctx.IFAllocationSale.connect(ctx.buyer).getMaxPayment(ctx.buyer.address)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['purchase(uint256)'](maxPayment.add(1))).to.be.revertedWith(EXCEED_MAX_PAYMENT)
  })

  generalTest.prototype.it = it('can vest with withdrawGiveaway', async function () {
    const withdrawDelay = 10000
    mineNext()
    // here set up a new ctx.IFAllocationSale with ctx.salePrice of 0, because
    // provided fixture sale does not have ctx.salePrice set to 0

    // deploy 0 price allocation sale
    const IFAllocationSaleFactory = await ethers.getContractFactory(
      contractName
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
    // fund sale
    await ctx.SaleToken.connect(ctx.seller).approve(
      ctx.IFAllocationSale.address,
      ctx.fundAmount
    ) // approve
    await ctx.IFAllocationSale.connect(ctx.seller).fund(ctx.fundAmount) // fund
    await ctx.IFAllocationSale.connect(ctx.owner).setWithdrawDelay(withdrawDelay)

    // set sale token allocation override (flat amount every participant receives)
    await ctx.IFAllocationSale.setSaleTokenAllocationOverride(33330)

    ctx.IFAllocationSale.connect(ctx.owner).setVestingEndTime(ctx.vestingEndTime)

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // set withdrawal delay
    await expect(ctx.IFAllocationSale.connect(ctx.owner).setWithdrawDelay(withdrawDelay)).to.be.revertedWith(SALE_IS_STARTED)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(CANNOT_WITHDRAW_YET)

    mineTimeDelta(ctx.endTime + withdrawDelay - (await getBlockTime()))
    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('1')

    mineTimeDelta(ctx.vestingEndTime - ctx.endTime)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('33330')
  })
})
