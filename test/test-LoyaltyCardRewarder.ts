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
  let user3: SignerWithAddress
  let user4: SignerWithAddress
  let user5: SignerWithAddress
  let user6: SignerWithAddress
  const KYC_CREDENTIAL_CODE = 0
  const KYC_CREDENTIAL_NAME = 'KYC'
  const KYC_CREDENTIAL_POINTS = 11
  const IDO_CREDENTIAL_CODE = 1
  const IDO_CREDENTIAL_NAME = 'IDO'
  const IDO_CREDENTIAL_POINTS = 22
  const DEX_CREDENTIAL_CODE = 2
  const DEX_CREDENTIAL_NAME = 'DEX'
  const DEX_CREDENTIAL_POINTS = 33

  beforeEach(async function () {
    owner = (await ethers.getSigners())[0]

    attacker = (await ethers.getSigners())[9]
    user = (await ethers.getSigners())[10]
    user2 = (await ethers.getSigners())[11]
    user3 = (await ethers.getSigners())[12]
    user4 = (await ethers.getSigners())[13]
    user5 = (await ethers.getSigners())[14]
    user6 = (await ethers.getSigners())[15]
    LoyaltyCardMaster = await ethers.getContractFactory('LoyaltyCardMaster')
    loyaltyCardMaster = await LoyaltyCardMaster.deploy(
      'ImpossibleLoyaltyCard',
      'ILC'
    )
    LoyaltyRewardsLookup = await ethers.getContractFactory(
      'LoyaltyRewardsLookup'
    )
    loyaltyRewardsLookup = await LoyaltyRewardsLookup.deploy()
    await loyaltyRewardsLookup.setCredential(
      KYC_CREDENTIAL_CODE,
      KYC_CREDENTIAL_POINTS,
      KYC_CREDENTIAL_NAME
    )
    await loyaltyRewardsLookup.setCredential(
      IDO_CREDENTIAL_CODE,
      IDO_CREDENTIAL_POINTS,
      IDO_CREDENTIAL_NAME
    )
    await loyaltyRewardsLookup.setCredential(
      DEX_CREDENTIAL_CODE,
      DEX_CREDENTIAL_POINTS,
      DEX_CREDENTIAL_NAME
    )

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
        .rewardAccount(user.address, KYC_CREDENTIAL_CODE, KYC_CREDENTIAL_NAME)
    ).to.be.revertedWith('Ownable: caller is not the owner')
  })

  it('Gives single account rewards based on account and the info provided by lookup contract', async function () {
    const points = await loyaltyRewardsLookup.getPoints(
      KYC_CREDENTIAL_CODE,
      KYC_CREDENTIAL_NAME
    )
    await loyaltyCardRewarder.rewardAccount(
      user.address,
      KYC_CREDENTIAL_CODE,
      KYC_CREDENTIAL_NAME
    )
    expect(
      await loyaltyCardMaster.currentPointsAccount(user.address)
    ).to.be.equal(points)
  })

  it('Gives single value rewards to a batch of accounts', async function () {
    await loyaltyCardMaster.mint(user2.address)
    await loyaltyCardMaster.mint(user3.address)
    await loyaltyCardMaster.mint(user4.address)
    await loyaltyCardMaster.mint(user5.address)
    await loyaltyCardMaster.mint(user6.address)
    const userAccs = [user, user2, user3, user4, user5, user6].map(
      (u) => u.address
    )
    const dexPoints = await loyaltyRewardsLookup.getPoints(
      DEX_CREDENTIAL_CODE,
      DEX_CREDENTIAL_NAME
    )
    await loyaltyCardRewarder.rewardBatchSingleCredential(
      userAccs,
      DEX_CREDENTIAL_CODE,
      DEX_CREDENTIAL_NAME
    )
    for (const acc of userAccs) {
      expect(await loyaltyCardMaster.currentPointsAccount(acc)).to.be.equal(
        dexPoints
      )
    }
  })

  it('Gives mixed rewards to a batch of accounts', async function () {
    await loyaltyCardMaster.mint(user2.address)
    await loyaltyCardMaster.mint(user3.address)
    await loyaltyCardMaster.mint(user4.address)
    await loyaltyCardMaster.mint(user5.address)
    await loyaltyCardMaster.mint(user6.address)

    const userAccs = [user, user2, user3, user4, user5, user6].map(
      (u) => u.address
    )
    const credCodes = [
      [KYC_CREDENTIAL_CODE, IDO_CREDENTIAL_CODE],
      [KYC_CREDENTIAL_CODE, DEX_CREDENTIAL_CODE],
      [IDO_CREDENTIAL_CODE, DEX_CREDENTIAL_CODE],
      [DEX_CREDENTIAL_CODE, KYC_CREDENTIAL_CODE],
      [IDO_CREDENTIAL_CODE, KYC_CREDENTIAL_CODE],
      [DEX_CREDENTIAL_CODE, IDO_CREDENTIAL_CODE],
    ]
    const credNames = [
      [KYC_CREDENTIAL_NAME, IDO_CREDENTIAL_NAME],
      [KYC_CREDENTIAL_NAME, DEX_CREDENTIAL_NAME],
      [IDO_CREDENTIAL_NAME, DEX_CREDENTIAL_NAME],
      [DEX_CREDENTIAL_NAME, KYC_CREDENTIAL_NAME],
      [IDO_CREDENTIAL_NAME, KYC_CREDENTIAL_NAME],
      [DEX_CREDENTIAL_NAME, IDO_CREDENTIAL_NAME],
    ]

    const kycPoints = +(await loyaltyRewardsLookup.getPoints(
      KYC_CREDENTIAL_CODE,
      KYC_CREDENTIAL_NAME
    ))
    const idoPoints = +(await loyaltyRewardsLookup.getPoints(
      IDO_CREDENTIAL_CODE,
      IDO_CREDENTIAL_NAME
    ))
    const dexPoints = +(await loyaltyRewardsLookup.getPoints(
      DEX_CREDENTIAL_CODE,
      DEX_CREDENTIAL_NAME
    ))
    const pointsAmounts = [
      kycPoints + idoPoints,
      kycPoints + dexPoints,
      idoPoints + dexPoints,
      dexPoints + kycPoints,
      idoPoints + kycPoints,
      dexPoints + idoPoints,
    ]

    await loyaltyCardRewarder.rewardBatchMultiCredentials(
      userAccs,
      credCodes,
      credNames
    )
    let i = 0
    for (const acc of userAccs) {
      expect(await loyaltyCardMaster.currentPointsAccount(acc)).to.be.equal(
        pointsAmounts[i++]
      )
    }
  })

  it('Should revert on data length mismatch when giving mixed rewards to a batch of user accounts', async function () {
    await loyaltyCardMaster.mint(user2.address)

    const userAccs = [user, user2].map((u) => u.address)
    const credCodesShort = [[KYC_CREDENTIAL_CODE, IDO_CREDENTIAL_CODE]]
    const credCodes = [
      [KYC_CREDENTIAL_CODE, IDO_CREDENTIAL_CODE],
      [KYC_CREDENTIAL_CODE, DEX_CREDENTIAL_CODE],
    ]
    const credCodesLong = [
      [KYC_CREDENTIAL_CODE, IDO_CREDENTIAL_CODE],
      [KYC_CREDENTIAL_CODE, DEX_CREDENTIAL_CODE],
      [IDO_CREDENTIAL_CODE, DEX_CREDENTIAL_CODE],
    ]
    const credNamesShort = [[KYC_CREDENTIAL_NAME, IDO_CREDENTIAL_NAME]]
    const credNames = [
      [KYC_CREDENTIAL_NAME, IDO_CREDENTIAL_NAME],
      [KYC_CREDENTIAL_NAME, DEX_CREDENTIAL_NAME],
    ]
    const credNamesLong = [
      [KYC_CREDENTIAL_NAME, IDO_CREDENTIAL_NAME],
      [KYC_CREDENTIAL_NAME, DEX_CREDENTIAL_NAME],
      [IDO_CREDENTIAL_NAME, DEX_CREDENTIAL_NAME],
    ]

    // any combination except correct length for both
    const argsCombinations = [
      [credCodesShort, credNamesShort],
      [credCodesShort, credNames],
      [credCodesShort, credNamesLong],
      [credCodes, credNamesShort],
      [credCodes, credNamesLong],
      [credCodesLong, credNamesShort],
      [credCodesLong, credNames],
      [credCodesLong, credNamesLong],
    ]

    for (const comb of argsCombinations) {
      await expect(
        loyaltyCardRewarder.rewardBatchMultiCredentials(
          userAccs,
          comb[0],
          comb[1]
        )
      ).to.be.revertedWith('BatchRewardLengthsMismatch')
    }
  })

  it('Cannot reward a user that has no card', async function () {
    expect(await loyaltyCardMaster.balanceOf(user2.address)).to.be.equal(0)
    expect(
      loyaltyCardRewarder.rewardAccount(
        user2.address,
        KYC_CREDENTIAL_CODE,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('TokenDoesntExist')
  })
})
