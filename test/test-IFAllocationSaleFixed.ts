import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { expect } from 'chai'
import { ethers } from 'hardhat'
import { computeMerkleProof, computeMerkleRoot, getAddressIndex, pad } from '../library/merkleWhitelist'
import { mineNext } from './helpers'
import IFAllocationSaleGeneralTest, { _ctx } from './IFAllocationSaleGeneralTest'

export default describe('IF Allocation Sale Fixed', function () {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const ctx: any = _ctx

  const contract = 'MockIFAllocationSaleFixed'

  const generalTest = IFAllocationSaleGeneralTest
  generalTest(this, contract, ctx)

  generalTest.prototype.it = it('can save allocation amount in merkle tree', async function () {
    const leaves: string[] = []
    const addressValMap = new Map();
    (await ethers.getSigners())
      .forEach((s: SignerWithAddress, i: number) => {
        const amount = ethers.utils.parseUnits((i + 1).toString(), 'ether')
        const packed = ethers.utils.solidityPack(
          ['address', 'uint256'],
          [s.address.toLowerCase(), amount],
        )
        leaves.push(packed)
        addressValMap.set(s.address.toLowerCase(), [packed, amount])
      }
    )
    leaves.sort()

    const merkleRoot = computeMerkleRoot(leaves)
    await ctx.IFAllocationSale.connect(ctx.owner).setWhitelistAllocation(merkleRoot)
    mineNext()

    const tempAcct = (await ethers.getSigners())[0]
    const [packed, amount] = addressValMap.get(tempAcct.address.toLowerCase())
    const tempAcctIdx = getAddressIndex(leaves, packed)
    expect(
      await ctx.IFAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        amount,

      )
    ).to.equal(true)
    const wrongAmount = '0x' + pad(ethers.constants.One.mul(100).toString().toLowerCase().replace('0x', ''))
    expect(
      await ctx.IFAllocationSale.connect(tempAcct).checkWhitelistAllocation(
        tempAcct.address,
        computeMerkleProof(leaves, tempAcctIdx),
        wrongAmount,
      )
    ).to.equal(false)
  })
})