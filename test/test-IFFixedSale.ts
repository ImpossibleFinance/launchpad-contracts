import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { expect } from 'chai'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex } from '../library/merkleWhitelist'
import IFAllocationSaleGeneralTest, { _ctx } from './IFAllocationSaleGeneralTest'
import { getBlockTime, mineNext, mineTimeDelta } from './helpers'
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { EXCEED_MAX_PAYMENT, NO_TOKEN_TO_BE_WITHDRAWN, NOT_A_GIVEAWAY } from './reverts/msg-IFAllocationSale'

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
  // unset timeout from the test
  this.timeout(0)

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const ctx: any = _ctx

  const contractName = 'MockIFFixedSale'

  const generalTest = IFAllocationSaleGeneralTest
  generalTest(this, contractName, ctx)


  generalTest.prototype.it = it('can save allocation amount in merkle tree', async function () {
    const signers = await ethers.getSigners()
    const allocations = Array(signers.length).fill(1)
    const [leaves, addressValMap] = computeMerkleRootWithAllocation(signers, allocations)
    const merkleRoot = computeMerkleRoot(leaves)
    await ctx.IFAllocationSale.connect(ctx.owner).setWhitelist(merkleRoot)
    mineNext()

    const tempAcct = (await ethers.getSigners())[0]
    const packed = addressValMap.get(tempAcct.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    expect(
      await ctx.IFAllocationSale.connect(tempAcct)['checkWhitelist(address,bytes32[],uint256)'](
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        1,
      )
    ).to.equal(true)
    expect(
      await ctx.IFAllocationSale.connect(tempAcct)['checkWhitelist(address,bytes32[],uint256)'](
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        200,
      )
    ).to.equal(false)
  })

  generalTest.prototype.it = it('can override sale token allocations (test preventing exceeding allocation)', async function () {

    const allocationAmount = 10000
    const paymentAmount = 100001
    const [leaves, addressValMap] = computeMerkleRootWithAllocation([ctx.buyer], [allocationAmount])
    // amount to pay (should fail, because this is 1 over allocation)

    // set sale token allocation override
    await ctx.IFAllocationSale.connect(ctx.owner).setWhitelist(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )

    const packed = addressValMap.get(ctx.buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['whitelistedPurchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      allocationAmount,
    )).to.be.revertedWith(EXCEED_MAX_PAYMENT)

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

    const allocationAmount = 5000
    // amount to pay for each claimer (should go through since this is exactly how much allocation they have)
    const paymentAmount = 50000

    const [leaves, addressValMap] = computeMerkleRootWithAllocation([ctx.buyer, ctx.buyer2], [allocationAmount, allocationAmount])
    await ctx.IFAllocationSale.connect(ctx.owner).setWhitelist(computeMerkleRoot(leaves))
    mineNext()

    // fast forward from current time to start time
    mineTimeDelta(ctx.startTime - (await getBlockTime()))

    // test purchase for buyers 1 and 2
    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )
    const packed = addressValMap.get(ctx.buyer.address.toLowerCase()) || ''
    const tempAcctIdx = getAddressIndex(leaves, packed)
    await ctx.IFAllocationSale.connect(ctx.buyer)['whitelistedPurchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx),
      allocationAmount,
    )

    mineNext()
    await ctx.PaymentToken.connect(ctx.buyer2).approve(
      ctx.IFAllocationSale.address,
      paymentAmount
    )

    const packed2 = addressValMap.get(ctx.buyer2.address.toLowerCase()) || ''
    const tempAcctIdx2 = getAddressIndex(leaves, packed2)
    await ctx.IFAllocationSale.connect(ctx.buyer2)['whitelistedPurchase(uint256,bytes32[],uint256)'](
      paymentAmount,
      computeMerkleProof(leaves, tempAcctIdx2),
      allocationAmount,
    )

    mineNext()

    // fast forward from current time to after end time
    mineTimeDelta(ctx.endTime - (await getBlockTime()))

    // test withdraw
    mineNext()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['withdrawGiveaway(bytes32[],uint256)']([], allocationAmount)).to.be.revertedWith(NOT_A_GIVEAWAY)
    await ctx.IFAllocationSale.connect(ctx.buyer).withdraw()
    mineNext()
    await ctx.IFAllocationSale.connect(ctx.buyer2).withdraw()
    // access control: Withdraw giveaway when sale price is not 0
    await expect(ctx.IFAllocationSale.connect(ctx.buyer)['withdrawGiveaway(bytes32[],uint256)']([], allocationAmount)).to.be.revertedWith(NOT_A_GIVEAWAY)
    mineNext()

    // expect balance to be 5000 for both buyers
    expect(await ctx.SaleToken.balanceOf(ctx.buyer.address)).to.equal('5000')
    expect(await ctx.SaleToken.balanceOf(ctx.buyer2.address)).to.equal('5000')

    // test purchaser counter
    expect(await ctx.IFAllocationSale.purchaserCount()).to.equal(2)

    // test withdrawer counter
    expect(await ctx.IFAllocationSale.withdrawerCount()).to.equal(2)
  })
})