import { Wallet } from 'zksync-web3'
import * as ethers from 'ethers'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { Deployer } from '@matterlabs/hardhat-zksync-deploy'

// load env file
import dotenv from 'dotenv'
dotenv.config()

// load wallet private key from env file
const PRIVATE_KEY = process.env.WALLET_PRIVATE_KEY || ''

if (!PRIVATE_KEY) throw '⛔️ Private key not detected! Add it to the .env file!'

// An example of a deploy script that will deploy and call a simple contract.
export default async function (hre: HardhatRuntimeEnvironment) {
  console.log('Running deploy script for the zk contract')

  // Initialize the wallet.
  const wallet = new Wallet(PRIVATE_KEY)
  const messageBus = '0x0000000000000000000000000000000000000000' // celer's message bus

  // Create deployer object and load the artifact of the contract you want to deploy.
  const deployer = new Deployer(hre, wallet)
  const artifact = await deployer.loadArtifact('IFAllocationMaster')

  // Estimate contract deployment fee
  const deploymentFee = await deployer.estimateDeployFee(artifact, [messageBus])

  const parsedFee = ethers.utils.formatEther(deploymentFee.toString())
  console.log(`The deployment is estimated to cost ${parsedFee}`)

  const masterContract = await deployer.deploy(artifact, [messageBus])

  //obtain the Constructor Arguments
  console.log(
    'constructor args:' + masterContract.interface.encodeDeploy([messageBus])
  )

  // Show the contract info.
  const contractAddress = masterContract.address
  console.log(`${artifact.contractName} was deployed to ${contractAddress}`)
}
