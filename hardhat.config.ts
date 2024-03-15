import '@nomiclabs/hardhat-waffle'
import '@nomiclabs/hardhat-etherscan'
import '@nomiclabs/hardhat-web3'
import 'hardhat-tracer'
import 'hardhat-abi-exporter'
import '@matterlabs/hardhat-zksync-deploy'
import '@matterlabs/hardhat-zksync-solc'
import '@matterlabs/hardhat-zksync-verify'
import 'solidity-coverage'

// import 'hardhat-gas-reporter'
// import 'hardhat-ethernal'
import dotenv from 'dotenv'

dotenv.config()

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

module.exports = {
  solidity: {
    version: '0.8.9',
    settings: {
      optimizer: { enabled: true },
    },
  },
  etherscan: {
    apiKey: {
      goerli: process.env.ETHERSCAN_API_KEY,
      kovan: process.env.ETHERSCAN_API_KEY,
      bscTestnet: process.env.BSCSCAN_API_KEY,
      arbitrumOne: process.env.ARBISCAN_API_KEY,
    },
  },
  networks: {
    zksolc: {
      version: 'latest',
      compilerSource: 'binary',
      settings: {
        optimizer: {
          enabled: true,
        },
      },
      url: '',
    },
    zkSyncMainnet: {
      url: 'https://zksync2-mainnet.zksync.io',
      ethNetwork: 'mainnet', // Can also be the RPC URL of the network (e.g. `https://goerli.infura.io/v3/<API_KEY>`)
      zksync: true,
      verifyURL:
        'https://zksync2-mainnet-explorer.zksync.io/contract_verification',
    },
    // defaultNetwork: 'zkSyncTestnet',
    zkSyncTestnet: {
      url: 'https://testnet.era.zksync.dev',
      ethNetwork: 'goerli', // or a Goerli RPC endpoint from Infura/Alchemy/Chainstack etc.
      zksync: true,
      verifyURL:
        'https://zksync2-testnet-explorer.zksync.dev/contract_verification',
    },
    hardhat: {
      forking: {
        url: 'https://bsc-dataseed.binance.org/',
        accounts: {
          mnemonic: process.env.MAINNET_MNEMONIC || '',
        },
      },
    },
    arbitrumOne: {
      url: 'https://arb1.arbitrum.io/rpc',
      chainId: 42161,
      gasPrice: 11000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    bsc_test: {
      url: 'https://data-seed-prebsc-1-s3.binance.org:8545',
      chainId: 97,
      gasPrice: 11000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    bsc_main: {
      url: 'https://bsc-dataseed.binance.org/',
      chainId: 56,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    eth_goerli: {
      url: 'https://rpc.goerli.mudit.blog/',
      chainId: 5,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    eth_ropsten: {
      url: 'https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
      accounts: {
        // first address: 0x99cb319980e55f4737c848e01BB74b8DE7863683
        mnemonic:
          'option skill video cause achieve joy section refuse infant goose any check',
      },
    },
    eth_main: {
      url: 'https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    polygon_main: {
      url: 'https://polygon-rpc.com',
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    avax_main: {
      url: 'https://api.avax.network/ext/bc/C/rpc',
      chainId: 43114,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    moonriver_main: {
      url: 'https://rpc.moonriver.moonbeam.network',
      chainId: 1285,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    kovan: {
      url: 'https://kovan.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
      chainId: 42,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    kava_test: {
      url: 'https://evm.testnet.kava.io',
      chainId: 2221,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    omni_test: {
      url: 'https://testnet.omni.network',
      chainId: 165,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    arbitrum_goerili: {
      url: 'https://goerli-rollup.arbitrum.io/rpc',
      chainId: 421613,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },
    linea_goerili: {
      url: 'https://rpc.goerli.linea.build',
      chainId: 59140,
      gasPrice: 5000000000,
      accounts: {
        mnemonic: process.env.MAINNET_MNEMONIC || '',
      },
    },

  },
}
