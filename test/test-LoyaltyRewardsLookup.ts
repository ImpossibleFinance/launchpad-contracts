import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { ethers } from 'hardhat'
import { Contract } from 'ethers'
import { expect } from 'chai'
import { BigNumber } from '@ethersproject/bignumber'

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
  const KYC_CREDENTIAL_CODE = 0
  const OTHER_CREDENTIAL_CODE = 1
  const KYC_CREDENTIAL_NAME = 'KYC'
  const KYC_CREDENTIAL_POINTS = 10
  const OTHER_CREDENTIAL_NAME = 'ABC'

  beforeEach(async function () {
    owner = (await ethers.getSigners())[0]
    user = (await ethers.getSigners())[10]
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

    LoyaltyCardRewarder = await ethers.getContractFactory('LoyaltyCardRewarder')
    loyaltyCardRewarder = await LoyaltyCardRewarder.deploy(
      loyaltyCardMaster.address,
      loyaltyRewardsLookup.address
    )

    await loyaltyCardMaster.addOperator(loyaltyCardRewarder.address)

    await loyaltyCardMaster.setMinter(owner.address)
    loyaltyCardMaster.mint(user.address)
  })

  it('Provides points for credential', async function () {
    expect(
      await loyaltyRewardsLookup.getPoints(
        KYC_CREDENTIAL_CODE,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.equal(KYC_CREDENTIAL_POINTS)
  })

  it('Refuses to provide credential points if given wrong credential name', async function () {
    expect(
      loyaltyRewardsLookup.getPoints(KYC_CREDENTIAL_CODE, OTHER_CREDENTIAL_NAME)
    ).to.be.revertedWith('CredentialMismatch')
  })

  it('Reverts on retrieval of credential points if credential is not set', async function () {
    // only KYC credential is set so far
    expect(
      loyaltyRewardsLookup.getPoints(
        OTHER_CREDENTIAL_CODE,
        OTHER_CREDENTIAL_NAME
      )
    ).to.be.revertedWith(`CredentialNotSet(${OTHER_CREDENTIAL_CODE})`)
  })

  it('Allows updating a credential name', async function () {
    const newName = 'New Name'
    await loyaltyRewardsLookup.updateCredentialName(
      KYC_CREDENTIAL_CODE,
      newName
    )
    expect(
      loyaltyRewardsLookup.getPoints(KYC_CREDENTIAL_CODE, KYC_CREDENTIAL_NAME)
    ).to.be.revertedWith('CredentialMismatch')
    expect(await loyaltyRewardsLookup.getName(KYC_CREDENTIAL_CODE)).to.be.equal(
      newName
    )
    expect(
      await loyaltyRewardsLookup.getPoints(KYC_CREDENTIAL_CODE, newName)
    ).to.but.equal(KYC_CREDENTIAL_POINTS)
  })

  it('Allows updating credential points', async function () {
    const newPoints = 123
    expect(newPoints).to.not.be.equal(KYC_CREDENTIAL_POINTS)
    await loyaltyRewardsLookup.updateCredentialPoints(
      KYC_CREDENTIAL_CODE,
      newPoints,
      KYC_CREDENTIAL_NAME
    )
    expect(
      await loyaltyRewardsLookup.getPoints(
        KYC_CREDENTIAL_CODE,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.equal(newPoints)
  })

  it('Refuses to update credential points if given wrong credential name', async function () {
    const newPoints = 123
    expect(
      loyaltyRewardsLookup.updateCredentialPoints(
        KYC_CREDENTIAL_CODE,
        newPoints,
        OTHER_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('CredentialMismatch')
  })

  it('Refuses to set credential if code is already in use', async function () {
    expect(
      loyaltyRewardsLookup.setCredential(
        KYC_CREDENTIAL_CODE,
        KYC_CREDENTIAL_POINTS,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('CredentialCodeAlreadyInUse')
  })

  it('Refuses an empty name when setting credential', async function () {
    const credPoints = 123
    expect(
      loyaltyRewardsLookup.setCredential(OTHER_CREDENTIAL_CODE, credPoints, '')
    ).to.be.revertedWith('EmptyCredentialName')
  })

  it('Refuses an empty name when updating credential', async function () {
    expect(
      loyaltyRewardsLookup.updateCredentialName(OTHER_CREDENTIAL_CODE, '')
    ).to.be.revertedWith('EmptyCredentialName')
  })

  it('Refuses a duplicate name when setting credential', async function () {
    const credPoints = 123
    expect(
      loyaltyRewardsLookup.setCredential(
        OTHER_CREDENTIAL_CODE,
        credPoints,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('DuplicateCredentialName')
  })

  it('Refuses a duplicate name when updating a credential name', async function () {
    const credPoints = 123
    await loyaltyRewardsLookup.setCredential(
      OTHER_CREDENTIAL_CODE,
      credPoints,
      OTHER_CREDENTIAL_NAME
    )
    expect(
      loyaltyRewardsLookup.updateCredentialName(
        OTHER_CREDENTIAL_CODE,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('DuplicateCredentialName')
  })

  it('Refuses to set a zero points credential', async function () {
    expect(
      loyaltyRewardsLookup.setCredential(
        OTHER_CREDENTIAL_CODE,
        OTHER_CREDENTIAL_NAME,
        0
      )
    ).to.be.revertedWith('ZeroCredentialPoints')
  })

  it('Refuses to update credential points with 0 value', async function () {
    expect(
      loyaltyRewardsLookup.updateCredentialPoints(
        KYC_CREDENTIAL_CODE,
        0,
        KYC_CREDENTIAL_NAME
      )
    ).to.be.revertedWith('ZeroCredentialPoints')
  })

  it('Returns all credential info for given array of credential codes', async function () {
    // kyc code (0) already set
    await loyaltyRewardsLookup.setCredential(1, 10, 'name0')
    await loyaltyRewardsLookup.setCredential(2, 11, 'name1')
    await loyaltyRewardsLookup.setCredential(3, 12, 'name2')
    const retrieved = await loyaltyRewardsLookup.getNamesAndPointsForAll([
      0, 1, 2, 3,
    ])
    const converted = retrieved.map((pair: [BigNumber, string]) => [
      pair[0].toNumber(),
      pair[1],
    ])
    expect(converted).to.eql([
      [KYC_CREDENTIAL_POINTS, KYC_CREDENTIAL_NAME],
      [10, 'name0'],
      [11, 'name1'],
      [12, 'name2'],
    ])
  })

  it('Reverts if batch credential info is queried with unknown (not set) credential codes', async function () {
    // so far only KYC code is set (0)
    expect(
      loyaltyRewardsLookup.getNamesAndPointsForAll([KYC_CREDENTIAL_CODE, 1])
    ).to.be.revertedWith('CredentialNotSet')
  })
})
