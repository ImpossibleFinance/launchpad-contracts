import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { ethers } from 'hardhat'
import { Contract } from 'ethers'
import { expect } from 'chai'

export default describe('Loyalty Card Master contract', function () {
  let LoyaltyCardMaster
  let loyaltyCardMaster: Contract
  let owner: SignerWithAddress
  let attacker: SignerWithAddress
  let user: SignerWithAddress
  let user2: SignerWithAddress
  let user3: SignerWithAddress
  let user4: SignerWithAddress
  let user5: SignerWithAddress
  let user6: SignerWithAddress

  let operator1: SignerWithAddress
  // let operator2: SignerWithAddress
  // let operator3: SignerWithAddress
  let destination1: SignerWithAddress
  // let destination2: SignerWithAddress
  // let destination3: SignerWithAddress

  beforeEach(async function () {
    owner = (await ethers.getSigners())[0]
    operator1 = (await ethers.getSigners())[1]
    // operator2 = (await ethers.getSigners())[2]
    // operator3 = (await ethers.getSigners())[3]
    destination1 = (await ethers.getSigners())[4]
    // destination2 = (await ethers.getSigners())[5]
    // destination3 = (await ethers.getSigners())[6]
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
  })

  // ============= MINTING & BURNING

  it('Only allows the owner to set a minter', async function () {
    expect(
      loyaltyCardMaster.connect(attacker).setMinter(attacker.address)
    ).to.be.revertedWith('Ownable: caller is not the owner')
  })

  it('Only allows the owner to set a burner', async function () {
    expect(
      loyaltyCardMaster.connect(attacker).setBurner(attacker.address)
    ).to.be.revertedWith('Ownable: caller is not the owner')
  })

  it('Only allows specified minter to mint', async function () {
    expect(
      loyaltyCardMaster.connect(attacker).mint(attacker.address)
    ).to.be.revertedWith('NotAllowedToMint')
  })

  it('Allows token owner to burn', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(1)
    const tokenId = 1
    expect(await loyaltyCardMaster.connect(user).burn(tokenId))
      .to.emit(loyaltyCardMaster, '')
      .withArgs(user.address, ethers.constants.AddressZero, tokenId)
  })

  it('Allows approved burner to burn', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(1)
    const tokenId = 1
    await loyaltyCardMaster.setBurner(owner.address)
    await loyaltyCardMaster.connect(user).approve(owner.address, tokenId)
    expect(await loyaltyCardMaster.burn(tokenId))
      .to.emit(loyaltyCardMaster, '')
      .withArgs(user.address, ethers.constants.AddressZero, tokenId)
  })

  it('Prevents burning if not by token owner or approved burner', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(1)
    const tokenId = 1

    // without approval
    expect(
      loyaltyCardMaster.connect(attacker).burn(tokenId)
    ).to.be.revertedWith('NotAllowedToBurn')

    // with approval
    await loyaltyCardMaster.connect(user).approve(attacker.address, tokenId)
    expect(
      loyaltyCardMaster.connect(attacker).burn(attacker.address)
    ).to.be.revertedWith('NotAllowedToBurn')
  })

  it('Should correctly keep track of mints and burns', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(0)
    expect(await loyaltyCardMaster.burnCounter()).to.equal(0)
    await loyaltyCardMaster.mint(user.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(1)
    await loyaltyCardMaster.mint(user2.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(2)
    await loyaltyCardMaster.setBurner(owner.address)
    const tokenId = 1
    await loyaltyCardMaster.connect(user).approve(owner.address, tokenId)
    expect(await loyaltyCardMaster.burn(tokenId))
    expect(await loyaltyCardMaster.burnCounter()).to.equal(1)
  })

  it('Should prevent a user from having multiple loyalty cards', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const tokenId1 = 1
    expect(await loyaltyCardMaster.balanceOf(user.address)).to.equal(1)
    expect(loyaltyCardMaster.mint(user.address)).to.be.revertedWith(
      'AlreadyOwnsCard'
    )
    await loyaltyCardMaster.setBurner(owner.address)
    await loyaltyCardMaster.connect(user).approve(owner.address, tokenId1)
    expect(await loyaltyCardMaster.burn(tokenId1))

    // Once an existing loyalty card is burned, minting is possible again
    await loyaltyCardMaster.mint(user.address)
    const tokenId2 = 2
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user.address)
    ).to.equal(tokenId2)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(2)
    expect(await loyaltyCardMaster.burnCounter()).to.equal(1)
  })

  it('Allows for batch minting', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(0)
    const users = [user, user2, user3, user4, user5].map((u) => u.address)
    await loyaltyCardMaster.mintForNonOwners(users)
    expect(await loyaltyCardMaster.mintCounter()).to.equal(users.length)
  })

  // ============= OWNED TOKEN LOOKUP

  it('Should provide the tokenId of the loyalty card of a given user', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    await loyaltyCardMaster.mint(user2.address)
    const tokenId1 = 1
    const tokenId2 = 2
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user.address)
    ).to.equal(tokenId1)
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user2.address)
    ).to.equal(tokenId2)
  })

  it('Should return tokenId 0 for users who do not own a loyalty card', async function () {
    expect(await loyaltyCardMaster.balanceOf(user.address)).to.equal(0)
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user.address)
    ).to.equal(0)
  })

  it('Should track the original owner of a card even when staked', async function () {
    // "staked" means transferred to a whitelisted destination
    // (typically an IF contract that accepts loyalty cards to be staked)
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)

    const mintedTokenId = await loyaltyCardMaster.originalOwnerToTokenId(
      user.address
    )

    await loyaltyCardMaster.addOperator(operator1.address)
    await loyaltyCardMaster.addDestination(destination1.address)
    await loyaltyCardMaster
      .connect(user)
      .transferFrom(user.address, destination1.address, mintedTokenId)

    const newTokenOwner = await loyaltyCardMaster.ownerOf(mintedTokenId)
    expect(newTokenOwner).to.equal(destination1.address)
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user.address)
    ).to.equal(mintedTokenId)
  })

  it('Should not treat staking contracts as loyalty card owners', async function () {
    // "staked" means transferred to a whitelisted destination
    // (typically an IF contract that accepts loyalty cards to be staked)
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.originalOwnerToTokenId(
      user.address
    )
    await loyaltyCardMaster.addOperator(operator1.address)
    await loyaltyCardMaster.addDestination(destination1.address)

    // staking destination not an "original owner"
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(destination1.address)
    ).to.equal(0)

    // STAKE
    await loyaltyCardMaster
      .connect(user)
      .transferFrom(user.address, destination1.address, mintedTokenId)
    const newTokenOwner = await loyaltyCardMaster.ownerOf(mintedTokenId)
    // nft owned by staking contract
    expect(newTokenOwner).to.equal(destination1.address)
    // original owner still the IF user
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(user.address)
    ).to.equal(mintedTokenId)
    // staking destination still not an "original owner"
    expect(
      await loyaltyCardMaster.originalOwnerToTokenId(destination1.address)
    ).to.equal(0)
  })

  // ============= OPERATORS

  it('Can add and remove operators', async function () {
    expect(await loyaltyCardMaster.isOperator(operator1.address)).to.equal(
      false
    )
    await expect(loyaltyCardMaster.addOperator(operator1.address))
      .to.emit(loyaltyCardMaster, 'AddedOperator')
      .withArgs(operator1.address)
    expect(await loyaltyCardMaster.isOperator(operator1.address)).to.equal(true)

    await expect(loyaltyCardMaster.removeOperator(operator1.address))
      .to.emit(loyaltyCardMaster, 'RemovedOperator')
      .withArgs(operator1.address)
    expect(await loyaltyCardMaster.isOperator(operator1.address)).to.equal(
      false
    )
  })

  // ============= POINTS

  it('Newly minted card should have no total and no current points', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.mintCounter()
    expect(await loyaltyCardMaster.totalPointsCard(mintedTokenId)).to.equal(0)
    expect(await loyaltyCardMaster.currentPointsCard(mintedTokenId)).to.equal(0)
  })

  it('Should only allow operator to add and redeem points', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.mintCounter()
    await expect(
      loyaltyCardMaster.connect(attacker).addPointsCard(mintedTokenId, 1)
    ).to.be.revertedWith('NotAllowedToOperate')
    await expect(
      loyaltyCardMaster.connect(attacker).redeemPointsCard(mintedTokenId, 1)
    ).to.be.revertedWith('NotAllowedToOperate')
    await loyaltyCardMaster.addOperator(operator1.address)
    await expect(
      loyaltyCardMaster.connect(operator1).addPointsCard(mintedTokenId, 1)
    ).to.emit(loyaltyCardMaster, 'AddedPoints')
    await expect(
      loyaltyCardMaster.connect(operator1).redeemPointsCard(mintedTokenId, 1)
    ).to.emit(loyaltyCardMaster, 'RedeemedPoints')
  })

  it('Should correctly account, add & redeem points per card (NFT)', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.mintCounter()
    await loyaltyCardMaster.addOperator(operator1.address)

    await loyaltyCardMaster.connect(operator1).addPointsCard(mintedTokenId, 1)
    expect(await loyaltyCardMaster.currentPointsCard(mintedTokenId)).to.equal(1)
    expect(await loyaltyCardMaster.totalPointsCard(mintedTokenId)).to.equal(1)
    await loyaltyCardMaster.connect(operator1).addPointsCard(mintedTokenId, 10)
    expect(await loyaltyCardMaster.currentPointsCard(mintedTokenId)).to.equal(
      11
    )
    expect(await loyaltyCardMaster.totalPointsCard(mintedTokenId)).to.equal(11)
    await loyaltyCardMaster
      .connect(operator1)
      .redeemPointsCard(mintedTokenId, 1)
    expect(await loyaltyCardMaster.currentPointsCard(mintedTokenId)).to.equal(
      10
    )
    expect(await loyaltyCardMaster.totalPointsCard(mintedTokenId)).to.equal(11)
  })

  it('Should not be able to redeem more points than available', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.mintCounter()
    await loyaltyCardMaster.addOperator(operator1.address)
    await loyaltyCardMaster.connect(operator1).addPointsCard(mintedTokenId, 10)
    await expect(
      loyaltyCardMaster.connect(operator1).redeemPointsCard(mintedTokenId, 11)
    ).to.be.revertedWith('InsufficientPoints')
  })

  it('Should correctly account, add & redeem points based on user account', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    await loyaltyCardMaster.addOperator(operator1.address)
    await loyaltyCardMaster
      .connect(operator1)
      .addPointsAccount(user.address, 10)
    expect(await loyaltyCardMaster.currentPointsAccount(user.address)).to.equal(
      10
    )
    await loyaltyCardMaster
      .connect(operator1)
      .redeemPointsAccount(user.address, 10)
    expect(await loyaltyCardMaster.currentPointsAccount(user.address)).to.equal(
      0
    )
  })

  it('Should be able to add the same amount of points to a batch of user accounts', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    await loyaltyCardMaster.mint(user2.address)
    await loyaltyCardMaster.mint(user3.address)
    await loyaltyCardMaster.mint(user4.address)
    await loyaltyCardMaster.mint(user5.address)
    await loyaltyCardMaster.mint(user6.address)
    await loyaltyCardMaster.addOperator(operator1.address)
    const userAccs = [user, user2, user3, user4, user5, user6].map(
      (u) => u.address
    )
    const pointsAmount = 11
    const multipliers = [11, 12, 13, 14, 15, 16]
    await loyaltyCardMaster
      .connect(operator1)
      .addPointsBatchAccSingleValue(userAccs, pointsAmount, multipliers)

    let idx = 0
    for (const acc of userAccs) {
      expect(await loyaltyCardMaster.currentPointsAccount(acc)).to.equal(
        pointsAmount * multipliers[idx]
      )
      idx++
    }
  })

  it('Should be able to add the different rewards to a batch of user accounts', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    await loyaltyCardMaster.mint(user2.address)
    await loyaltyCardMaster.mint(user3.address)
    await loyaltyCardMaster.mint(user4.address)
    await loyaltyCardMaster.mint(user5.address)
    await loyaltyCardMaster.mint(user6.address)
    await loyaltyCardMaster.addOperator(operator1.address)
    const userAccs = [user, user2, user3, user4, user5, user6].map(
      (u) => u.address
    )
    const pointsAmounts = [11, 12, 13, 14, 15, 16]
    await loyaltyCardMaster
      .connect(operator1)
      .addPointsBatchAccMultiValues(userAccs, pointsAmounts)

    let i = 0
    for (const acc of userAccs) {
      expect(await loyaltyCardMaster.currentPointsAccount(acc)).to.equal(
        pointsAmounts[i++]
      )
    }
  })

  it('Should revert on data length mismatch when adding different rewards to a batch of user accounts', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    await loyaltyCardMaster.mint(user2.address)
    const userAccs = [user, user2].map((u) => u.address)

    const pointsAmountsTooShort = [11]
    const pointsAmountsTooLong = [11, 12, 13]
    const pointsAmounts = [11, 12]

    await expect(
      loyaltyCardMaster
        .connect(operator1)
        .addPointsBatchAccMultiValues(userAccs, pointsAmountsTooShort)
    ).to.be.revertedWith('BatchRewardLengthsMismatch')

    await expect(
      loyaltyCardMaster
        .connect(operator1)
        .addPointsBatchAccMultiValues(userAccs, pointsAmountsTooLong)
    ).to.be.revertedWith('BatchRewardLengthsMismatch')

    const multipliersTooShort = [11]
    const multipliersTooLong = [11, 12, 13]
    await expect(
      loyaltyCardMaster
        .connect(operator1)
        .addPointsBatchAccSingleValue(
          userAccs,
          pointsAmounts,
          multipliersTooShort
        )
    ).to.be.revertedWith('BatchRewardLengthsMismatch')

    await expect(
      loyaltyCardMaster
        .connect(operator1)
        .addPointsBatchAccSingleValue(
          userAccs,
          pointsAmounts,
          multipliersTooLong
        )
    ).to.be.revertedWith('BatchRewardLengthsMismatch')
  })

  // ============= TRANSFERS

  it('Can add, remove & confirm destinations', async function () {
    expect(
      await loyaltyCardMaster.isDestination(destination1.address)
    ).to.equal(false)

    loyaltyCardMaster.addDestination(destination1.address)
    await expect(loyaltyCardMaster.addDestination(destination1.address))
      .to.emit(loyaltyCardMaster, 'AddedDestination')
      .withArgs(destination1.address)
    expect(
      await loyaltyCardMaster.isDestination(destination1.address)
    ).to.equal(true)

    await expect(loyaltyCardMaster.removeDestination(destination1.address))
      .to.emit(loyaltyCardMaster, 'RemovedDestination')
      .withArgs(destination1.address)
    expect(
      await loyaltyCardMaster.isDestination(destination1.address)
    ).to.equal(false)
  })

  it('Should not allow transfer to a destination that is not whitelisted', async function () {
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.mintCounter()
    let mintedTokenOwner = await loyaltyCardMaster.ownerOf(mintedTokenId)

    expect(mintedTokenOwner).to.equal(user.address)

    await loyaltyCardMaster.addOperator(operator1.address)

    expect(
      await loyaltyCardMaster.isDestination(destination1.address)
    ).to.equal(false)

    await expect(
      loyaltyCardMaster
        .connect(user)
        .transferFrom(user.address, destination1.address, mintedTokenId)
    ).to.be.revertedWith('NotAllowedAsDestination')

    await loyaltyCardMaster.addDestination(destination1.address)
    await loyaltyCardMaster
      .connect(user)
      .transferFrom(user.address, destination1.address, mintedTokenId)
    mintedTokenOwner = await loyaltyCardMaster.ownerOf(mintedTokenId)

    expect(mintedTokenOwner).to.equal(destination1.address)
  })

  it('Should recognize transferred tokens as staked', async function () {
    // "staked" means transferred to a whitelisted destination
    await loyaltyCardMaster.setMinter(owner.address)
    await loyaltyCardMaster.mint(user.address)
    const mintedTokenId = await loyaltyCardMaster.originalOwnerToTokenId(
      user.address
    )
    await loyaltyCardMaster.addDestination(destination1.address)

    // does not count as STAKED
    expect(await loyaltyCardMaster.isStaked(mintedTokenId)).to.equal(false)

    // STAKE
    await loyaltyCardMaster
      .connect(user)
      .transferFrom(user.address, destination1.address, mintedTokenId)
    const newTokenOwner = await loyaltyCardMaster.ownerOf(mintedTokenId)
    // nft owned by staking contract
    expect(newTokenOwner).to.equal(destination1.address)

    // now it counts as STAKED
    expect(await loyaltyCardMaster.isStaked(mintedTokenId)).to.equal(true)
  })
})
