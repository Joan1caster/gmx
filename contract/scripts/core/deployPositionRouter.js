const { getFrameSigner, deployContract, contractAt , sendTxn, readTmpAddresses, writeTmpAddresses } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const { toUsd } = require("../../test/shared/units")
const utils  = require("../../utils.js")

const network = (process.env.HARDHAT_NETWORK || 'mainnet');
const tokens = require('./tokens')[network];

const {
  ANVIL_URL,
  ANVIL_DEPLOY_KEY,
  ARBITRUM_URL,
  ARBITRUM_CAP_KEEPER_KEY,
  AVAX_URL,
  AVAX_CAP_KEEPER_KEY,
} = require("../../env.json")

async function getAnvilValues(signer) {
  const provider = new ethers.providers.JsonRpcProvider(ANVIL_URL)
  const capKeeperWallet = new ethers.Wallet(ANVIL_DEPLOY_KEY).connect(provider)
  const positionUtils = await contractAt("PositionUtils", utils.getYamlValue("PositionUtils"))
  const vault = await contractAt("Vault", utils.getYamlValue("Vault"))
  const router = await contractAt("Router", utils.getYamlValue("Router"))
  const weth = await contractAt("WETH", utils.getYamlValue("WETH"))
  const referralStorage = await contractAt("ReferralStorage", "0xe6fab3F0c7199b0d34d7FbE83394fc0e0D06e99d")
  const shortsTracker = await contractAt("ShortsTracker", utils.getYamlValue("ShortsTracker"), signer)
  const shortsTrackerTimelock = await contractAt("ShortsTrackerTimelock", "0xf58eEc83Ba28ddd79390B9e90C4d3EbfF1d434da", signer)
  const depositFee = "30" // 0.3%
  const minExecutionFee = "100000000000000" // 0.0001 ETH

  return {
    positionUtils,
    capKeeperWallet,
    vault,
    // timelock,
    router,
    weth,
    referralStorage,
    shortsTracker,
    shortsTrackerTimelock,
    depositFee,
    minExecutionFee,
    // positionKeepers
  }
}

async function getAvaxValues(signer) {
  const provider = new ethers.providers.JsonRpcProvider(AVAX_URL)
  const capKeeperWallet = new ethers.Wallet(AVAX_CAP_KEEPER_KEY).connect(provider)

  const vault = await contractAt("Vault", "0x9ab2De34A33fB459b538c43f251eB825645e8595")
  const router = await contractAt("Router", await vault.router(), signer)
  const weth = await contractAt("WETH", tokens.nativeToken.address)
  const referralStorage = await contractAt("ReferralStorage", "0x827ED045002eCdAbEb6e2b0d1604cf5fC3d322F8")
  const shortsTracker = await contractAt("ShortsTracker", "0x9234252975484D75Fd05f3e4f7BdbEc61956D73a", signer)
  const shortsTrackerTimelock = await contractAt("ShortsTrackerTimelock", "0xf58eEc83Ba28ddd79390B9e90C4d3EbfF1d434da", signer)
  const depositFee = "30" // 0.3%
  const minExecutionFee = "20000000000000000" // 0.02 AVAX

  return {
    capKeeperWallet,
    vault,
    timelock,
    router,
    weth,
    referralStorage,
    shortsTracker,
    shortsTrackerTimelock,
    depositFee,
    minExecutionFee
  }
}

async function getValues(signer) {
  if (network === "arbitrum") {
    return getArbValues(signer)
  }

  if (network === "avax") {
    return getAvaxValues(signer)
  }

  if (network === "anvil") {
    return getAnvilValues(signer)
  }
}

async function main() {
  const signer = await getFrameSigner()

  const {
    positionUtils,
    capKeeperWallet,
    vault,
    timelock,
    router,
    weth,
    shortsTracker,
    shortsTrackerTimelock,
    depositFee,
    minExecutionFee,
    referralStorage
  } = await getValues(signer)
  const positionRouterArgs = [vault.address, router.address, weth.address, shortsTracker.address, depositFee, minExecutionFee]
  const positionRouterFactor = await ethers.getContractFactory("PositionRouter",{
    libraries: {
      PositionUtils: positionUtils.address
    }
})
  const positionRouter = positionRouterFactor.deploy(...positionRouterArgs)
  setTimeout(function() {
    console.log("1秒钟已到!");
  }, 1000);

  // await sendTxn(positionRouter.setReferralStorage(referralStorage.address), "positionRouter.setReferralStorage")

  await sendTxn(router.addPlugin(positionRouter.address), "router.addPlugin")

  await sendTxn(positionRouter.setDelayValues(0, 180, 30 * 60), "positionRouter.setDelayValues")
  await sendTxn(timelock.setContractHandler(positionRouter.address, true), "timelock.setContractHandler(positionRouter)")

  await sendTxn(positionRouter.setGov(await vault.gov()), "positionRouter.setGov")

  await sendTxn(positionRouter.setAdmin( vault.gov()), "positionRouter.setAdmin")
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
