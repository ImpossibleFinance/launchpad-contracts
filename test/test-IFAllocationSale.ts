import '@nomiclabs/hardhat-ethers'
import { expect } from 'chai'
import { BigNumber } from 'ethers'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex } from '../library/merkleWhitelist'
import { getBlockTime, mineNext, mineTimeDelta } from './helpers'
import IFAllocationSaleGeneralTest, { _ctx, _ctxFree, _ctxSale } from './IFAllocationSaleGeneralTest'
import { EXCEED_MAX_PAYMENT, NOT_WHITELIST_SETTER_OR_OWNER, PROOF_INVALID, SALE_ALREADY_STARTED, USE_WITHDRAWGIVEAWAY } from './reverts/msg-IFAllocationSale'

export default describe('IF Allocation Sale', function () {
  this.timeout(0)
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const ctx: any = _ctx
  const ctxFree: any = _ctxFree
  const ctxSale: any = _ctxSale

  const contractName = 'IFAllocationSale'

  const generalTest = IFAllocationSaleGeneralTest
  generalTest(this, contractName, ctx, ctxFree, _ctxSale)

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

    ctx.IFAllocationSale.connect(ctx.owner).setLinearVestingEndTime(ctx.linearVestingEndTime)

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // set withdrawal delay
    await expect(ctx.IFAllocationSale.connect(ctx.owner).setWithdrawDelay(withdrawDelay)).to.be.revertedWith(SALE_ALREADY_STARTED)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.reverted

    mineTimeDelta(ctx.endTime + withdrawDelay - (await getBlockTime()))
    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])
    const value: number = await ctx.IFAllocationSale.getUserStakeValue(ctx.buyer.address)
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal(value / 20000)

    mineTimeDelta(ctx.linearVestingEndTime - ctx.endTime)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdrawGiveaway([])
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal(value)
  })
  it('can get cliff period', async function () {
    const pct = [10, 20, 30, 40]
    await ctx.IFAllocationSale.connect(ctx.owner).setCliffPeriod(
      [
        ctx.endTime + 10,
        ctx.endTime + 20,
        ctx.endTime + 30,
        ctx.endTime + 40,
      ],
      pct
    )
    const cliffPeriod = await ctx.IFAllocationSale.getCliffPeriod()
    expect(cliffPeriod[0].pct).eq(pct[0])
  })
})
