const { deployContract, contractAt , sendTxn } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const { toUsd } = require("../../test/shared/units")
const BigNumber = require('bignumber.js');

const { errors } = require("../../test/core/Vault/helpers")

const network = (process.env.HARDHAT_NETWORK || 'mainnet');

async function main() {

  const vault = await deployContract("Vault", [])
  const usdg = await deployContract("USDG", [vault.address])
  const weth = await deployContract("WETH", ["WETH", "WETH", 18])
  const usdt = await deployContract("USDT", ["USDT", "USDT"])
  const btc = await deployContract("BTC", ["BTC", "BTC"])
  const router = await deployContract("Router", [vault.address, usdg.address, weth.address])
  const usdtPriceFeed = await deployContract("USDTPriceFeed", [])
  const btcPriceFeed = await deployContract("BTCPriceFeed", [])
  const vaultPriceFeed = await deployContract("VaultPriceFeed", [])
  await sendTxn(vaultPriceFeed.setMaxStrictPriceDeviation(expandDecimals(1, 28)), "vaultPriceFeed.setMaxStrictPriceDeviation") // 0.05 USD
  await sendTxn(vaultPriceFeed.setPriceSampleSpace(1), "vaultPriceFeed.setPriceSampleSpace")
  await sendTxn(vaultPriceFeed.setIsAmmEnabled(false), "vaultPriceFeed.setIsAmmEnabled")
  await sendTxn(vaultPriceFeed.setTokenConfig(usdt.address, usdtPriceFeed.address, usdt.decimals(), true), "vaultPriceFeed.setTokenConfig")
  await sendTxn(vaultPriceFeed.setTokenConfig(btc.address, btcPriceFeed.address, btc.decimals(), false), "vaultPriceFeed.setTokenConfig")
  const _maxTokenNum = new BigNumber("1e70");
  const maxTokenNum = ethers.BigNumber.from(_maxTokenNum.toFixed());
  await sendTxn(vault.setTokenConfig(usdt.address, usdt.decimals(), 1000, 15, maxTokenNum, true, false))
  await sendTxn(vault.setTokenConfig(btc.address, btc.decimals(), 1000, 15, maxTokenNum, false, true))

  const glp = await deployContract("GLP", [])
  await sendTxn(glp.setInPrivateTransferMode(true), "glp.setInPrivateTransferMode")
  const glpManager = await deployContract("GlpManager", [vault.address, usdg.address, glp.address, glp.address, 15 * 60])
  await sendTxn(glpManager.setInPrivateMode(true), "glpManager.setInPrivateMode")

  await sendTxn(glp.setMinter(glpManager.address, true), "glp.setMinter")
  await sendTxn(usdg.addVault(glpManager.address), "usdg.addVault(glpManager)")

  await sendTxn(vault.initialize(
    router.address, // router
    usdg.address, // usdg
    vaultPriceFeed.address, // priceFeed
    toUsd(2), // liquidationFeeUsd
    100, // fundingRateFactor
    100 // stableFundingRateFactor
  ), "vault.initialize")

  await sendTxn(vault.setFundingRate(60 * 60, 100, 100), "vault.setFundingRate")

  await sendTxn(vault.setInManagerMode(true), "vault.setInManagerMode")
  await sendTxn(vault.setManager(glpManager.address, true), "vault.setManager")

  await sendTxn(vault.setFees(
    10, // _taxBasisPoints
    5, // _stableTaxBasisPoints
    20, // _mintBurnFeeBasisPoints
    20, // _swapFeeBasisPoints
    1, // _stableSwapFeeBasisPoints
    10, // _marginFeeBasisPoints
    toUsd(2), // _liquidationFeeUsd
    24 * 60 * 60, // _minProfitTime
    true // _hasDynamicFees
  ), "vault.setFees")

  const vaultErrorController = await deployContract("VaultErrorController", [])
  await sendTxn(vault.setErrorController(vaultErrorController.address), "vault.setErrorController")
  await sendTxn(vaultErrorController.setErrors(vault.address, errors), "vaultErrorController.setErrors")

  const vaultUtils = await deployContract("VaultUtils", [vault.address])
  await sendTxn(vault.setVaultUtils(vaultUtils.address), "vault.setVaultUtils")
  process.exit(0);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
