import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { Contract } from '@ethersproject/contracts'
import { mineNext, getBlockTime, mineTimeDelta } from './helpers'
import { first } from 'lodash'

const MaxUint256 = ethers.constants.MaxUint256
const WeiPerEth = ethers.constants.WeiPerEther
const _1 = ethers.constants.One
const _10000 = ethers.BigNumber.from(10000)
const TWO_WEEKS = 14 * 86400

const convToBN = (num: number) => {
  return ethers.BigNumber.from(num).mul(WeiPerEth)
}

export default describe('vIDIA', function () {
  // unset timeout from the test
  this.timeout(0)

  let vIDIA: Contract
  let underlying: Contract
  let owner: SignerWithAddress
  let vester: SignerWithAddress

  beforeEach(async function () {
    // Get the ContractFactory and Signers here.
    // Token = await ethers.getContractFactory("Token");
    [owner, vester] = await ethers.getSigners()

    // To deploy our contract, we just have to call Token.deploy() and await
    // for it to be deployed(), which happens once its transaction has been
    // mined.
    const vIDIAFactory = await ethers.getContractFactory('vIDIA')

    owner = (await ethers.getSigners())[0]
    vester = (await ethers.getSigners())[1]
    const TestTokenFactory = await ethers.getContractFactory('GenericToken')
    underlying = await TestTokenFactory.connect(owner).deploy(
      'Test Vest Token',
      'Vest',
      MaxUint256
    )
    vIDIA = await vIDIAFactory.deploy(
      'vIDIA contract',
      'VIDIA',
      owner.address,
      underlying.address
    )
  })

  it('test setters', async function () {
    const value = [0, 100, 200, 300]
    const fns = [
      { set: vIDIA.updateSkipDelayFee, get: vIDIA.skipDelayFee},
      { set: vIDIA.updateCancelUnstakeFee, get: vIDIA.cancelUnstakeFee},
      { set: vIDIA.updateUnstakingDelay, get: vIDIA.unstakingDelay}
    ]

    for (let i = 0; i < fns.length; i++) {
      for (let j = 0; j < value.length; j++) {
        await fns[i].set(value[j])
        expect(await fns[i].get()).to.eq(value[j])
      }

      if (i !== 2) // fee setters should throw when setting >100%
        await expect(fns[i].set(10001)).to.be.revertedWith('Fee must be less than 100%')
    }
  })

  it('test stake tokens', async function () {
    const transferAmt = 10000000
    await underlying.transfer(vester.address, transferAmt) 
    await underlying.connect(vester).approve(vIDIA.address, ethers.constants.MaxUint256)
    const stakeAmt = [100, 250]

    for (let i = 0; i < stakeAmt.length; i++) {
      await vIDIA.connect(vester).stake(stakeAmt[i])
      expect((await vIDIA.totalStakedAmount()).toNumber())
        .to.eq(stakeAmt.reduce((prev, curr, idx) => idx <= i ? prev+curr : prev))
    }
  })

  it('test stake/unstake', async function () {
    const transferAmt = 10000000

    await underlying.transfer(vester.address, transferAmt) 
    await underlying.connect(vester).approve(vIDIA.address, ethers.constants.MaxUint256)
    const firstStakeAmt = 100
    const secondStakeAmt = 250
    await vIDIA.connect(vester).stake(firstStakeAmt)
    let totalStaked = (await vIDIA.totalStakedAmount()).toNumber()
    expect(totalStaked).to.eq(firstStakeAmt)

    await vIDIA.connect(vester).stake(secondStakeAmt)
    totalStaked = (await vIDIA.totalStakedAmount()).toNumber()
    expect(totalStaked).to.eq(firstStakeAmt + secondStakeAmt)
    await vIDIA.connect(vester).unstake(secondStakeAmt)
    const userData = await vIDIA.userInfo(vester.address)
    expect(userData.unstakedAmount).to.eq(secondStakeAmt)
    const unstakeTime =
      (await getBlockTime()) + (await vIDIA.unstakingDelay()).toNumber()
    expect(userData.unstakeAt).to.eq(unstakeTime)
    await expect(vIDIA.connect(vester).unstake(firstStakeAmt))
      .to.be.revertedWith('User has pending tokens unstaking')

  })

  it('test whitelist feature', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    await vIDIA.stake(WeiPerEth)
    await vIDIA.approve(vester.address, MaxUint256)

    const checkFailure = async () => {
      await expect(vIDIA.transfer(vester.address, _1)).to.be.revertedWith(
        'Origin and dest address not in whitelist'
      )
      await expect(
        vIDIA.connect(vester).transferFrom(owner.address, vester.address, _1)
      ).to.be.revertedWith('Origin and dest address not in whitelist')
    }

    const checkSuccess = async () => {
      await expect(vIDIA.transfer(vester.address, _1))
        .to.emit(vIDIA, 'Transfer')
        .withArgs(owner.address, vester.address, _1)
      await expect(
        vIDIA.connect(vester).transferFrom(owner.address, vester.address, _1)
      )
        .to.emit(vIDIA, 'Transfer')
        .withArgs(owner.address, vester.address, _1)
    }

    const checkWhitelist = async (addrArr: string[]) => {
      expect(JSON.stringify(await vIDIA.getAllWhitelistedAddrs())).to.eq(
        JSON.stringify(addrArr)
      )
    }

    // case 1: no whitelist, should fail transfer
    await checkWhitelist([])
    await checkFailure()

    // case 2: source addr in whitelist, should not fail xfer
    await vIDIA.addToWhitelist(owner.address)
    await checkWhitelist([owner.address])
    await checkSuccess()

    // case 3: source addr and dest addr in whitelist, should not fail xfer
    await vIDIA.addToWhitelist(vester.address)
    await checkWhitelist([owner.address, vester.address])
    await checkSuccess()

    // case 4: dest addr in whitelist, should not fail xfer
    await vIDIA.removeFromWhitelist(owner.address)
    await checkWhitelist([vester.address])
    await checkSuccess()

    // case 5: remove all addr from whitelist, should fail xfer
    await vIDIA.removeFromWhitelist(vester.address)
    await checkWhitelist([])
    await checkFailure()
  })

  it('test claimstaked', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    await vIDIA.stake(convToBN(200))

    const withdrawAmt = [convToBN(1), convToBN(12), convToBN(0), convToBN(123)]

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    const feePercentBasisPts = await vIDIA.skipDelayFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      
      const fee = feePercentBasisPts.mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      const reward = await vIDIA.calculateUserReward()

      expect(await vIDIA.claimStaked(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimReward')
        .withArgs(owner.address, reward)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, reward)
        .to.emit(vIDIA, 'ClaimStaked')
        .withArgs(owner.address, fee, receiveAmt)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, receiveAmt)

      expect(await vIDIA.balanceOf(owner.address)).to.equal(userVidiaBalance.sub(withdrawAmt[i]))
      expect(await underlying.balanceOf(owner.address)).to.equal(userUnderlying.add(receiveAmt).add(reward))
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees.add(fee))
      expect(await underlying.balanceOf(vIDIA.address)).to.equal(contractUnderlying.sub(receiveAmt).sub(reward))

      userVidiaBalance = userVidiaBalance.sub(withdrawAmt[i])
      userUnderlying = userUnderlying.add(receiveAmt).add(reward)
      contractUnderlying = contractUnderlying.sub(receiveAmt).sub(reward)
      sumFees = sumFees.add(fee)
    }
  })

  it('test claimpendingunstake with pending unstake', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    const stakeAmt = 200
    await vIDIA.stake(convToBN(stakeAmt))
    await vIDIA.unstake(convToBN(stakeAmt))

    const withdrawAmt = [convToBN(1), convToBN(8), convToBN(0), convToBN(102)]

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let userUnstakingAmt = (await vIDIA.userInfo(owner.address)).unstakedAmount
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    const feePercentBasisPts = await vIDIA.skipDelayFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      const fee = feePercentBasisPts.mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      expect(await vIDIA.claimPendingUnstake(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimPendingUnstake')
        .withArgs(owner.address, fee, receiveAmt)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, receiveAmt)

      expect(await vIDIA.balanceOf(owner.address)).to.equal(userVidiaBalance)
      expect(await underlying.balanceOf(owner.address)).to.equal(userUnderlying.add(receiveAmt))
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees.add(fee))
      expect((await vIDIA.userInfo(owner.address)).unstakedAmount).to.equal(userUnstakingAmt.sub(withdrawAmt[i]))
      expect(await underlying.balanceOf(vIDIA.address)).to.equal(contractUnderlying.sub(receiveAmt))

      userUnderlying = userUnderlying.add(receiveAmt)
      userUnstakingAmt = userUnstakingAmt.sub(withdrawAmt[i])
      contractUnderlying = contractUnderlying.sub(receiveAmt)
      sumFees = sumFees.add(fee)
    }

    // test failure mode
    await mineTimeDelta(TWO_WEEKS)
    await expect(vIDIA.claimPendingUnstake(0))
      .to.be.revertedWith('Can unstake without paying fee')
  })

  it('test cancelpendingunstake with pending unstake', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    const stakeAmt = 200
    await vIDIA.stake(convToBN(stakeAmt))
    await vIDIA.unstake(convToBN(stakeAmt))

    const withdrawAmt = [convToBN(1), convToBN(6), convToBN(0), convToBN(99)]

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let userUnstakingAmt = (await vIDIA.userInfo(owner.address)).unstakedAmount
    let userStakedAmt = (await vIDIA.userInfo(owner.address)).stakedAmount
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    const feePercentBasisPts = await vIDIA.cancelUnstakeFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      const fee = feePercentBasisPts.mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      const reward = await vIDIA.calculateUserReward()

      expect(await vIDIA.cancelPendingUnstake(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimReward')
        .withArgs(owner.address, reward)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, reward)
        .to.emit(vIDIA, 'CancelPendingUnstake')
        .withArgs(owner.address, fee, receiveAmt)

      expect(await vIDIA.balanceOf(owner.address)).to.equal(userVidiaBalance.add(receiveAmt))
      expect(await underlying.balanceOf(owner.address)).to.equal(userUnderlying.add(reward)) // receive reward
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees.add(fee))
      expect((await vIDIA.userInfo(owner.address)).unstakedAmount).to.equal(userUnstakingAmt.sub(withdrawAmt[i]))
      expect((await vIDIA.userInfo(owner.address)).stakedAmount).to.equal(userStakedAmt.add(receiveAmt))
      expect(await underlying.balanceOf(vIDIA.address)).to.equal(contractUnderlying)
  
      userVidiaBalance = userVidiaBalance.add(receiveAmt)
      userUnderlying = userUnderlying.add(reward)
      userUnstakingAmt = userUnstakingAmt.sub(withdrawAmt[i])
      contractUnderlying = contractUnderlying.sub(reward)
      userStakedAmt = userStakedAmt.add(receiveAmt)
      sumFees = sumFees.add(fee)
    }

    // test failure mode
    await mineTimeDelta(TWO_WEEKS)
    await expect(vIDIA.claimPendingUnstake(0))
      .to.be.revertedWith('Can unstake without paying fee')
  })
})
