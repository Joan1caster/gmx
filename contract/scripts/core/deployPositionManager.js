const { getFrameSigner, deployContract, contractAt , sendTxn } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const { toUsd } = require("../../test/shared/units")
const { errors } = require("../../test/core/Vault/helpers")
const utils = require("../../utils.js");
const network = (process.env.HARDHAT_NETWORK || 'anvil');
const tokens = require('./tokens')[network];

const depositFee = 30 // 0.3%

async function getAnvilValues(signer) {
  const vault = await contractAt("Vault", utils.getYamlValue("Vault"), signer)
  console.log("Vault contract:", vault.address);
  const timelock = await contractAt("Timelock", await vault.gov(), signer)
  const router = await contractAt("Router", await vault.router(), signer)
  const shortsTracker = await contractAt("ShortsTracker", utils.getYamlValue("ShortsTracker"), signer)
  const weth = await contractAt("WETH", tokens.nativeToken.address)
  const orderBook = await contractAt("OrderBook", utils.getYamlValue("OrderBook"))
  const referralStorage = await contractAt("ReferralStorage", utils.getYamlValue("Vault"))
  const options = {
    network: "http://127.0.0.1:8545",
  }

  const orderKeepers = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" },
  ]
  const liquidators = [
    { address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" }
  ]

  const partnerContracts = [
  ]

  return { vault, timelock, router, shortsTracker, weth, depositFee, orderBook, referralStorage, orderKeepers, liquidators, partnerContracts }
}

async function getAvaxValues(signer) {
  const vault = await contractAt("Vault", "0x9ab2De34A33fB459b538c43f251eB825645e8595")
  const timelock = await contractAt("Timelock", await vault.gov(), signer)
  const router = await contractAt("Router", await vault.router(), signer)
  const shortsTracker = await contractAt("ShortsTracker", "0x9234252975484D75Fd05f3e4f7BdbEc61956D73a", signer)
  const weth = await contractAt("WETH", tokens.nativeToken.address)
  const orderBook = await contractAt("OrderBook", "0x4296e307f108B2f583FF2F7B7270ee7831574Ae5")
  const referralStorage = await contractAt("ReferralStorage", "0x827ed045002ecdabeb6e2b0d1604cf5fc3d322f8")

  const orderKeepers = [
    { address: "0x06f34388A7CFDcC68aC9167C5f1C23DD39783179" },
    { address: "0xf26f52d5985F6391E541A8d638e1EDaa522Ae56C" }
  ]
  const liquidators = [
    { address: "0x7858A4C42C619a68df6E95DF7235a9Ec6F0308b9" }
  ]

  const partnerContracts = []

  return { vault, timelock, router, shortsTracker, weth, depositFee, orderBook, referralStorage, orderKeepers, liquidators, partnerContracts }
}


async function getValues(signer) {

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

  const positionUtils = await deployContract("PositionUtils", [])
  let positionManager
  if (positionManagerAddress) {
    console.log("Using position manager at", positionManagerAddress)
    positionManager = await contractAt("PositionManager", positionManagerAddress)
  } else {
    console.log("Deploying new position manager")
    const positionManagerArgs = [vault.address, router.address, shortsTracker.address, weth.address, depositFee, orderBook.address]
    
    // 获取合约工厂时需要包含库链接
    const PositionManager = await ethers.getContractFactory("PositionManager", {
      libraries: {
          PositionUtils: positionUtils.address
      }
  })
  
  // 部署合约
  positionManager = await PositionManager.deploy(...positionManagerArgs)
  await positionManager.deployTransaction.wait()
  console.info(`Deployed PositionManager at ${positionManager.address}`)
    utils.saveData("PositionManager", positionManager.address)
  }

  // positionManager only reads from referralStorage so it does not need to be set as a handler of referralStorage
  // if ((await positionManager.referralStorage()).toLowerCase() != referralStorage.address.toLowerCase()) {
  //   await sendTxn(positionManager.setReferralStorage(referralStorage.address), "positionManager.setReferralStorage")
  // }
  // if (await positionManager.shouldValidateIncreaseOrder()) {
  //   await sendTxn(positionManager.setShouldValidateIncreaseOrder(false), "positionManager.setShouldValidateIncreaseOrder(false)")
  // }

  for (let i = 0; i < orderKeepers.length; i++) {
    const orderKeeper = orderKeepers[i]
    if (!(await positionManager.isOrderKeeper(orderKeeper.address))) {
      await sendTxn(positionManager.setOrderKeeper(orderKeeper.address, true), "positionManager.setOrderKeeper(orderKeeper)")
    }
  }
  console.log("positionManager: setOrderKeeper succeed.")
  for (let i = 0; i < liquidators.length; i++) {
    const liquidator = liquidators[i]
    if (!(await positionManager.isLiquidator(liquidator.address))) {
      await sendTxn(positionManager.setLiquidator(liquidator.address, true), "positionManager.setLiquidator(liquidator)")
    }
  }
  console.log("positionManager: setLiquidator succeed.")
  // if (!(await timelock.isHandler(positionManager.address))) {
  //   await sendTxn(timelock.setContractHandler(positionManager.address, true), "timelock.setContractHandler(positionManager)")
  // }
  // console.log("timelock: setContractHandler succeed.")
  if (!(await vault.isLiquidator(positionManager.address))) {
    await sendTxn(timelock.setLiquidator(vault.address, positionManager.address, true), "timelock.setLiquidator(vault, positionManager, true)")
  }
  console.log("timelock: setLiquidator succeed.")
  if (!(await shortsTracker.isHandler(positionManager.address))) {
    await sendTxn(shortsTracker.setHandler(positionManager.address, true), "shortsTracker.setContractHandler(positionManager.address, true)")
  }
  console.log("shortsTracker: setHandler succeed.")
  if (!(await router.plugins(positionManager.address))) {
    await sendTxn(router.addPlugin(positionManager.address), "router.addPlugin(positionManager)")
  }
  console.log("router: addPlugin succeed.")
  for (let i = 0; i < partnerContracts.length; i++) {
    const partnerContract = partnerContracts[i]
    if (!(await positionManager.isPartner(partnerContract))) {
      await sendTxn(positionManager.setPartner(partnerContract, true), "positionManager.setPartner(partnerContract)")
    }
    console.log("positionManager: setPartner succeed.")
  }

  if ((await positionManager.gov()) != (await vault.gov())) {
    await sendTxn(positionManager.setGov(await vault.gov()), "positionManager.setGov")
    console.log("positionManager: setGov succeed.")
  }

  console.log("done.")
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
