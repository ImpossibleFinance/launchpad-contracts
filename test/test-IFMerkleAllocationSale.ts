import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { expect } from 'chai'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex, pad } from '../library/merkleWhitelist'
import { getBlockTime, mineNext, mineTimeDelta } from './helpers'
import IFAllocationSaleGeneralTest, { _ctx } from './IFAllocationSaleGeneralTest'
import { EXCEED_MAX_PAYMENT, NO_TOKEN_TO_BE_WITHDRAWN, NOT_A_GIVEAWAY } from './reverts/msg-IFAllocationSale'

function computeMerkleRootWithAllocation(signers: SignerWithAddress[], allocations: number[]): [string[], Map<string, string>]{
    const leaves: string[] = []
    const addressValMap = new Map()
    signers.forEach((s: SignerWithAddress, i: number) => {
        const amount = ethers.utils.parseUnits(allocations[i].toString(), 'ether')
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

export default describe('IF Allocation Sale Fixed', async function () {
  this.timeout(0)
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const ctx: any = _ctx

  const generalTest = IFAllocationSaleGeneralTest
  generalTest(this, 'MockIFMerkleAllocationSale', ctx)

  const IFAllocationSaleFactory = await ethers.getContractFactory(
    'IFMerkleAllocationSale'
  )

  const IFMerkleAllocationSale = await IFAllocationSaleFactory.deploy(
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

  generalTest.prototype.it = it('can save allocation amount in merkle tree', async function () {
    const signers = await ethers.getSigners()
    const allocations = Array(signers.length).fill(1)
    const [leaves, addressValMap] = computeMerkleRootWithAllocation(signers, allocations)
    const merkleRoot = computeMerkleRoot(leaves)
    await IFMerkleAllocationSale.connect(ctx.owner).setWhitelistAllocation(merkleRoot)
    mineNext()

    const tempAcct = (await ethers.getSigners())[0]
    const packed = addressValMap.get(tempAcct.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    expect(
      await IFMerkleAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        ethers.utils.parseUnits('1', 'ether'),
      )
    ).to.equal(true)
    expect(
      await IFMerkleAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        ethers.utils.parseUnits('200', 'ether'),
      )
    ).to.equal(false)
  })

  generalTest.prototype.it = it('can override sale token allocations (test preventing exceeding allocation)', async function () {
    const allocationAmount = '100000'
    const paymentAmount = '100001'
    const [leaves, addressValMap] = computeMerkleRootWithAllocation([ctx.buyer], [parseInt(allocationAmount)])
    // amount to pay (should fail, because _this is 1 over allocation)

    // set sale token allocation override
    await IFMerkleAllocationSale.connect(ctx.owner).setWhitelistAllocation(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      IFMerkleAllocationSale.address,
      paymentAmount
    )

    const packed = addressValMap.get(ctx.buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await expect(IFMerkleAllocationSale.connect(ctx.buyer)['purchase(uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      ethers.utils.parseUnits(paymentAmount, 'ether'),
    )).to.be.revertedWith(EXCEED_MAX_PAYMENT)

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    await expect(IFMerkleAllocationSale.connect(ctx.buyer).withdraw()).to.be.revertedWith(NO_TOKEN_TO_BE_WITHDRAWN)
    mineNext()

    // expect balance to be 0
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('0')
  })

  generalTest.prototype.it = it('can override sale token allocations (test multiple buyers)', async function () {
    mineNext()

    // amount to pay for each claimer (should go through since _this is exactly how much allocation they have)
    const paymentAmount = '50000'
    const allocationAmount = '5000'

    const [leaves, addressValMap] = computeMerkleRootWithAllocation([ctx.buyer, ctx.buyer2], [parseInt(allocationAmount), parseInt(allocationAmount)])
    await IFMerkleAllocationSale.connect(ctx.owner).setWhitelistAllocation(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase for buyers 1 and 2
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      IFMerkleAllocationSale.address,
      paymentAmount
    )
    const packed = addressValMap.get(ctx.buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await IFMerkleAllocationSale.connect(ctx.buyer)['purchase(uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      ethers.utils.parseUnits(allocationAmount, 'ether'),
    )

    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer2).approve(
      IFMerkleAllocationSale.address,
      paymentAmount
    )

    const packed2 = addressValMap.get(ctx.buyer.address.toLowerCase()) || ''
    const tempAcctIdx2 = getAddressIndex(leaves, packed2)
    await IFMerkleAllocationSale.connect(ctx.buyer2)['purchase(uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx2),
      ethers.utils.parseUnits(allocationAmount, 'ether'),
    )

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFMerkleAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    await IFMerkleAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()
    await IFMerkleAllocationSale.connect(ctx.buyer2).withdraw()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(IFMerkleAllocationSale.connect(ctx.buyer).withdrawGiveaway([])).to.be.revertedWith(NOT_A_GIVEAWAY)
    mineNext()

    // expect balance to be 5000 for both buyers
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('5000')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('5000')

    // test purchaser counter
    expect(await IFMerkleAllocationSale.purchaserCount()).to.equal(2)

    // test withdrawer counter
    expect(await IFMerkleAllocationSale.withdrawerCount()).to.equal(2)
  })
})