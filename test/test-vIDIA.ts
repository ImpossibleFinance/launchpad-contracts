import '@nomiclabs/hardhat-ethers'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { Contract } from '@ethersproject/contracts'
import { mineNext, getBlockTime, mineTimeDelta } from './helpers'
import { first } from 'lodash'
import { BigNumber } from 'ethers'

const MaxUint256 = ethers.constants.MaxUint256
const WeiPerEth = ethers.constants.WeiPerEther
const _0 = ethers.constants.Zero
const _1 = ethers.constants.One
const _10 = BigNumber.from(10)
const _10000 = BigNumber.from(10000)
const FACTOR = BigNumber.from(_10.pow(BigNumber.from(30)))

const TWO_WEEKS = 14 * 86400

const convToBN = (num: number) => {
  return BigNumber.from(num).mul(WeiPerEth)
}

const checkWithinTolerance = (test: BigNumber, target: BigNumber, tolerance = _10) => {
  expect(test.add(tolerance).gte(target)).to.eq(true, `failed gte tolerance, ${test.toString()} and ${target.toString()}`)
  expect(test.sub(tolerance).lte(target)).to.eq(true, `failed lte tolerance, ${test.toString()} and ${target.toString()}`)
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

    await underlying.transfer(vester.address, convToBN(1000))
    await underlying.approve(vIDIA.address, MaxUint256)
    await underlying.connect(vester).approve(vIDIA.address, MaxUint256)
  })

  it('test static funcs', async function () {
    expect(await vIDIA.supportsInterface('0xb0202a11')).to.eq(true)
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
    let userData = await vIDIA.userInfo(vester.address)
    expect(userData.unstakedAmount).to.eq(secondStakeAmt)
    const unstakeTime =
      (await getBlockTime()) + (await vIDIA.unstakingDelay()).toNumber()
    expect(userData.unstakeAt).to.eq(unstakeTime)
    await expect(
      vIDIA.connect(vester).unstake(firstStakeAmt)
    ).to.be.revertedWith('User has pending tokens unstaking')

    // test claimUnstaked
    await mineTimeDelta((await vIDIA.unstakingDelay()).toNumber())
    const preUnstake = await underlying.balanceOf(vester.address)
    await vIDIA.connect(vester).claimUnstaked()
    expect((await underlying.balanceOf(vester.address)).toString())
      .to.eq(preUnstake.add(BigNumber.from(secondStakeAmt)))
    userData = await vIDIA.userInfo(vester.address)
    expect(userData.unstakeAt).to.eq(0)
    expect(userData.unstakedAmount).to.eq(0)
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
    const ownerStakeAmt = convToBN(200)
    const rewarderStakeAmt = convToBN(1)
    await vIDIA.stake(ownerStakeAmt)
    await vIDIA.connect(vester).stake(rewarderStakeAmt)

    const withdrawAmt = [convToBN(1), convToBN(12), convToBN(0), convToBN(123)]

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      
      const fee = (await vIDIA.skipDelayFee()).mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      const newRewardSum = 
        (await vIDIA.rewardSum())
          .add(fee.mul(FACTOR)
            .div(rewarderStakeAmt))

      const reward = 
        rewarderStakeAmt.mul(
          newRewardSum.sub((await vIDIA.userInfo(vester.address)).lastRewardSum))
            .div(FACTOR)

      expect(await vIDIA.claimStaked(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimReward')
        .withArgs(owner.address, _0)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, _0)
        .to.emit(vIDIA, 'ClaimStaked')
        .withArgs(owner.address, fee, receiveAmt)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, receiveAmt)

      expect(await vIDIA.calculateUserReward()).to.eq(_0) // fees should never accrue to for fee payer

      // these are the state changes every loop
      sumFees = sumFees.add(fee)
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees)
      expect(await vIDIA.connect(vester).calculateUserReward()).to.eq(sumFees) // all fees to vester
      checkWithinTolerance(reward, sumFees) // default tolerance = 10wei

      userVidiaBalance = userVidiaBalance.sub(withdrawAmt[i])
      expect(await vIDIA.balanceOf(owner.address)).to.equal(userVidiaBalance)

      userUnderlying = userUnderlying.add(receiveAmt)
      expect(await underlying.balanceOf(owner.address)).to.equal(userUnderlying)

      contractUnderlying = contractUnderlying.sub(receiveAmt)
      expect(await underlying.balanceOf(vIDIA.address)).to.equal(contractUnderlying)
    }
  })

  it('test claimpendingunstake with pending unstake', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    const ownerStakeAmt = convToBN(200)
    const rewarderStakeAmt = convToBN(1)
    await vIDIA.stake(ownerStakeAmt)
    await vIDIA.connect(vester).stake(rewarderStakeAmt)
    await vIDIA.unstake(ownerStakeAmt.sub(rewarderStakeAmt))

    // sums up to stakeAmt-1 for LOC coverage
    const withdrawAmt = [convToBN(1), convToBN(6), convToBN(0), convToBN(99), convToBN(93)] 

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let userStakedAmt = (await vIDIA.userInfo(owner.address)).stakedAmount
    let userUnstakingAmt = (await vIDIA.userInfo(owner.address)).unstakedAmount
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      const fee = (await vIDIA.skipDelayFee()).mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      const newRewardSum = 
        (await vIDIA.rewardSum())
          .add(fee.mul(FACTOR)
            .div(rewarderStakeAmt))

      const reward = 
        rewarderStakeAmt.mul(
          newRewardSum.sub((await vIDIA.userInfo(vester.address)).lastRewardSum))
            .div(FACTOR)

      expect(await vIDIA.claimPendingUnstake(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimReward')
        .withArgs(owner.address, _0)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, _0)
        .to.emit(vIDIA, 'ClaimPendingUnstake')
        .withArgs(owner.address, fee, receiveAmt)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, receiveAmt)

      // no change
      expect(await vIDIA.calculateUserReward()).to.eq(_0) // fees should never accrue to for fee payer
      expect(await vIDIA.balanceOf(owner.address)).to.equal(userVidiaBalance)

      sumFees = sumFees.add(fee)
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees)
      expect(await vIDIA.connect(vester).calculateUserReward()).to.eq(sumFees) // all fees to vester
      checkWithinTolerance(reward, sumFees) // default tolerance = 10wei

      userUnderlying = userUnderlying.add(receiveAmt)
      expect(await underlying.balanceOf(owner.address)).to.equal(userUnderlying)

      userUnstakingAmt = userUnstakingAmt.sub(withdrawAmt[i])
      expect((await vIDIA.userInfo(owner.address)).unstakedAmount).to.equal(userUnstakingAmt)

      contractUnderlying = contractUnderlying.sub(receiveAmt)
      expect(await underlying.balanceOf(vIDIA.address)).to.equal(contractUnderlying)
    }

    // test failure mode
    await mineTimeDelta((await vIDIA.unstakingDelay()).toNumber())
    await expect(vIDIA.claimPendingUnstake(0))
      .to.be.revertedWith('Can unstake without paying fee')
  })

  it('test cancelpendingunstake with pending unstake', async () => {
    await underlying.approve(vIDIA.address, MaxUint256)
    const ownerStakeAmt = convToBN(200)
    const rewarderStakeAmt = convToBN(1)
    await vIDIA.stake(ownerStakeAmt)
    await vIDIA.connect(vester).stake(rewarderStakeAmt)
    await vIDIA.unstake(ownerStakeAmt.sub(rewarderStakeAmt))

    // sums up to stakeAmt-1 for LOC coverage
    const withdrawAmt = [convToBN(1), convToBN(6), convToBN(0), convToBN(99), convToBN(93)] 

    let userVidiaBalance = await vIDIA.balanceOf(owner.address)
    let userUnderlying = await underlying.balanceOf(owner.address)
    let userUnstakingAmt = (await vIDIA.userInfo(owner.address)).unstakedAmount
    let userStakedAmt = (await vIDIA.userInfo(owner.address)).stakedAmount
    let contractUnderlying = await underlying.balanceOf(vIDIA.address)
    let sumFees = await vIDIA.accumulatedFee()

    for (let i = 0; i < withdrawAmt.length; i++) {
      const fee = (await vIDIA.cancelUnstakeFee()).mul(withdrawAmt[i]).div(_10000) // 10000 basis pts = 100%
      const receiveAmt = withdrawAmt[i].sub(fee)

      const newRewardSum = 
        (await vIDIA.rewardSum())
          .add(fee.mul(FACTOR)
            .div(rewarderStakeAmt))

      const reward = 
        rewarderStakeAmt.mul(
          newRewardSum.sub((await vIDIA.userInfo(vester.address)).lastRewardSum))
            .div(FACTOR)
      
      expect(await vIDIA.cancelPendingUnstake(withdrawAmt[i]))
        .to.emit(vIDIA, 'ClaimReward')
        .withArgs(owner.address, _0)
        .to.emit(underlying, 'Transfer')
        .withArgs(vIDIA.address, owner.address, _0)
        .to.emit(vIDIA, 'CancelPendingUnstake')
        .withArgs(owner.address, fee, receiveAmt)

      expect(await vIDIA.calculateUserReward()).to.eq(_0) // fees should never accrue to for fee payer

      sumFees = sumFees.add(fee)
      expect(await vIDIA.accumulatedFee()).to.equal(sumFees)
      expect(await vIDIA.connect(vester).calculateUserReward()).to.eq(sumFees) // all fees to vester
      checkWithinTolerance(reward, sumFees) // default tolerance = 10wei

      // no change
      expect(await underlying.balanceOf(owner.address)) // full fee since owner owns 100% of totalstaked
        .to.equal(userUnderlying) 
      expect(await underlying.balanceOf(vIDIA.address)) // full fee sent out since its all owner
        .to.equal(contractUnderlying)

      userVidiaBalance = userVidiaBalance.add(receiveAmt)
      expect(await vIDIA.balanceOf(owner.address)) // receives receiveAmt which is amt - fee
        .to.equal(userVidiaBalance)

      userUnstakingAmt = userUnstakingAmt.sub(withdrawAmt[i])
      expect((await vIDIA.userInfo(owner.address)).unstakedAmount) // reduce unstakedAmt by amt
        .to.equal(userUnstakingAmt)

      userStakedAmt = userStakedAmt.add(receiveAmt)
      expect((await vIDIA.userInfo(owner.address)).stakedAmount) // inc stakedAmt by receiveAmt
        .to.equal(userStakedAmt)
    }

    // test failure mode
    await mineTimeDelta((await vIDIA.unstakingDelay()).toNumber())
    await expect(vIDIA.cancelPendingUnstake(0))
      .to.be.revertedWith('Can restake without paying fee')
  })

  it('test claimunstaked', async () => {
    await vIDIA.claimUnstaked()
  })
})
