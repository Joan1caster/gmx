const { getFrameSigner, deployContract, contractAt , sendTxn } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const { toUsd } = require("../../test/shared/units")
const { errors } = require("../../test/core/Vault/helpers")
const utils = require("../../utils")
const network = (process.env.HARDHAT_NETWORK || 'mainnet');

const depositFee = 30 // 0.3%

async function getAnvilValues(signer) {
  const vault = await contractAt("Vault", utils.getYamlValue("Vault"), signer)
  // const timelock = await contractAt("Timelock", await vault.gov(), signer)
  const router = await contractAt("Router", await vault.router(), signer)
  // const shortsTracker = await contractAt("ShortsTracker", "0xf58eEc83Ba28ddd79390B9e90C4d3EbfF1d434da", signer)
  const weth = await contractAt("WETH", utils.getYamlValue("WETH"))
  const orderBook = await contractAt("OrderBook", utils.getYamlValue("OrderBook"))
  // const referralStorage = await contractAt("ReferralStorage", "0xe6fab3f0c7199b0d34d7fbe83394fc0e0d06e99d")

  const orderKeepers = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" }
  ]
  const liquidators = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" }
  ]

  const partnerContracts = [
    utils.getYamlValue("WETH"),
    utils.getYamlValue("Vault"),
    utils.getYamlValue("Vault"),
    utils.getYamlValue("Vault"),
  ]

  return { vault, router, weth, depositFee, orderBook, orderKeepers, liquidators, partnerContracts }
}

async function getAvaxValues(signer) {
  const vault = await contractAt("Vault", "0x9ab2De34A33fB459b538c43f251eB825645e8595")
  // const timelock = await contractAt("Timelock", await vault.gov(), signer)
  const router = await contractAt("Router", await vault.router(), signer)
  // const shortsTracker = await contractAt("ShortsTracker", "0x9234252975484D75Fd05f3e4f7BdbEc61956D73a", signer)
  const weth = await contractAt("WETH", tokens.nativeToken.address)
  const orderBook = await contractAt("OrderBook", "0x4296e307f108B2f583FF2F7B7270ee7831574Ae5")
  // const referralStorage = await contractAt("ReferralStorage", "0x827ed045002ecdabeb6e2b0d1604cf5fc3d322f8")
  // order agent account
  const orderKeepers = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" },
  ]
  // liquidation agent account
  const liquidators = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" }
  ]

  const partnerContracts = []

  return { vault, timelock, router,  weth, depositFee, orderBook, referralStorage, orderKeepers, liquidators, partnerContracts }
}

async function getValues(signer) {
  if (network === "anvil") {
    return getAnvilValues(signer)
  }

  if (network === "avax") {
    return getAvaxValues(signer)
  }
}

async function main() {
  const signer = await getFrameSigner()

  const {
    positionManagerAddress,
    vault,
    timelock,
    router,
    shortsTracker,
    weth,
    depositFee,
    orderBook,
    referralStorage,
    orderKeepers,
    liquidators,
    partnerContracts
  } = await getValues(signer)

  let positionManager
  if (positionManagerAddress) {
    console.log("Using position manager at", positionManagerAddress)
    positionManager = await contractAt("PositionManager", positionManagerAddress)
  } else {
    console.log("Deploying new position manager")
    const positionUtils = await deployContract("PositionUtils", [])
    const PositionManager = await ethers.getContractFactory("PositionManager", {
      libraries: {
          PositionUtils: positionUtils.address
      }
    })
    const positionManagerArgs = [vault.address, router.address, router.address, weth.address, depositFee, orderBook.address]
    positionManager = await PositionManager.deploy(...positionManagerArgs)
    await positionManager.deployTransaction.wait()
    console.info(`Deployed PositionManager at ${positionManager.address}`)
    utils.saveData("PositionManager", positionManager.address)
  }

  // positionManager only reads from referralStorage so it does not need to be set as a handler of referralStorage
  // if ((await positionManager.referralStorage()).toLowerCase() != referralStorage.address.toLowerCase()) {
  //   await sendTxn(positionManager.setReferralStorage(referralStorage.address), "positionManager.setReferralStorage")
  // }
  if (await positionManager.shouldValidateIncreaseOrder()) {
    await sendTxn(positionManager.setShouldValidateIncreaseOrder(false), "positionManager.setShouldValidateIncreaseOrder(false)")
  }

  for (let i = 0; i < orderKeepers.length; i++) {
    const orderKeeper = orderKeepers[i]
    if (!(await positionManager.isOrderKeeper(orderKeeper.address))) {
      await sendTxn(positionManager.setOrderKeeper(orderKeeper.address, true), "positionManager.setOrderKeeper(orderKeeper)")
    }
  }

  for (let i = 0; i < liquidators.length; i++) {
    const liquidator = liquidators[i]
    if (!(await positionManager.isLiquidator(liquidator.address))) {
      await sendTxn(positionManager.setLiquidator(liquidator.address, true), "positionManager.setLiquidator(liquidator)")
    }
  }

  // if (!(await timelock.isHandler(positionManager.address))) {
  //   await sendTxn(timelock.setContractHandler(positionManager.address, true), "timelock.setContractHandler(positionManager)")
  // }
  // if (!(await vault.isLiquidator(positionManager.address))) {
  //   await sendTxn(timelock.setLiquidator(vault.address, positionManager.address, true), "timelock.setLiquidator(vault, positionManager, true)")
  // }
  // if (!(await shortsTracker.isHandler(positionManager.address))) {
  //   await sendTxn(shortsTracker.setHandler(positionManager.address, true), "shortsTracker.setContractHandler(positionManager.address, true)")
  // }
  if (!(await router.plugins(positionManager.address))) {
    await sendTxn(router.addPlugin(positionManager.address), "router.addPlugin(positionManager)")
  }

  for (let i = 0; i < partnerContracts.length; i++) {
    const partnerContract = partnerContracts[i]
    if (!(await positionManager.isPartner(partnerContract))) {
      await sendTxn(positionManager.setPartner(partnerContract, true), "positionManager.setPartner(partnerContract)")
    }
  }

  if ((await positionManager.gov()) != (await vault.gov())) {
    await sendTxn(positionManager.setGov(await vault.gov()), "positionManager.setGov")
  }

  console.log("done.")
  process.exit(0);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
