import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { ethers } from 'hardhat'
import { Contract } from 'ethers'
import { expect } from 'chai'

export default describe('Loyalty Card Rewarder contract', function () {
  let LoyaltyCardMaster
  let loyaltyCardMaster: Contract
  let LoyaltyRewardsLookup
  let loyaltyRewardsLookup: Contract
  let LoyaltyCardRewarder
  let loyaltyCardRewarder: Contract
  let owner: SignerWithAddress
  let attacker: SignerWithAddress
  let user: SignerWithAddress
  let user2: SignerWithAddress
  const KYC_CREDENTIAL = 0

  beforeEach(async function () {
    owner = (await ethers.getSigners())[0]

    attacker = (await ethers.getSigners())[9]
    user = (await ethers.getSigners())[10]
    user2 = (await ethers.getSigners())[11]
    LoyaltyCardMaster = await ethers.getContractFactory('LoyaltyCardMaster')
    loyaltyCardMaster = await LoyaltyCardMaster.deploy(
      'ImpossibleLoyaltyCard',
      'ILC'
    )
    LoyaltyRewardsLookup = await ethers.getContractFactory(
      'LoyaltyRewardsLookup'
    )
    loyaltyRewardsLookup = await LoyaltyRewardsLookup.deploy()

    LoyaltyCardRewarder = await ethers.getContractFactory('LoyaltyCardRewarder')
    loyaltyCardRewarder = await LoyaltyCardRewarder.deploy(
      loyaltyCardMaster.address,
      loyaltyRewardsLookup.address
    )

    await loyaltyCardMaster.addOperator(loyaltyCardRewarder.address)

    await loyaltyCardMaster.setMinter(owner.address)
    loyaltyCardMaster.mint(user.address)
  })

  it('Only gives rewards if called by owner', async function () {
    expect(
      loyaltyCardRewarder
        .connect(attacker)
        .rewardAccount(user.address, KYC_CREDENTIAL)
    ).to.be.revertedWith('Ownable: caller is not the owner')
  })

  it('Gives rewards based on account and the info provided by lookup contract', async function () {
    const points = await loyaltyRewardsLookup.pointsForCredential(
      KYC_CREDENTIAL
    )
    await loyaltyCardRewarder.rewardAccount(user.address, KYC_CREDENTIAL)
    expect(
      await loyaltyCardMaster.currentPointsAccount(user.address)
    ).to.be.equal(points)
  })

  it('Cannot reward a user that has no card', async function () {
    expect(await loyaltyCardMaster.balanceOf(user2.address)).to.be.equal(0)
    expect(
      loyaltyCardRewarder.rewardAccount(user2.address, KYC_CREDENTIAL)
    ).to.be.revertedWith('TokenDoesntExist')
  })
})
