const { deployContract, contractAt, sendTxn, callWithRetries } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const { toChainlinkPrice } = require("../../test/shared/chainlink")
const utils = require("../../utils.js");

const network = (process.env.HARDHAT_NETWORK || 'mainnet');
const tokens = require('./tokens')[network];

async function main() {
  const vault = await contractAt("Vault",  utils.getYamlValue("Vault"))
  const vaultPriceFeed = await contractAt("VaultPriceFeed", utils.getYamlValue("VaultPriceFeed"))
  console.log("vault", vault.address)
  console.log("vaultPriceFeed", vaultPriceFeed.address)
  
  const WETH = utils.getYamlValue("WETH")
  const BTC = utils.getYamlValue("BTC")
  const USDT = utils.getYamlValue("USDT")
  const tokenArr = [WETH, BTC, USDT]

  for (const token of tokenArr) {
    await sendTxn(vaultPriceFeed.setTokenConfig(
      token.address, // _token
      token.priceFeed, // _priceFeed
      token.priceDecimals, // _priceDecimals
      token.isStrictStable // _isStrictStable
    ), `vaultPriceFeed.setTokenConfig(${token.name}) ${token.address} ${token.priceFeed}`)

    await sendTxn(vault.setTokenConfig(
      token.address, // _token
      token.decimals, // _tokenDecimals
      token.tokenWeight, // _tokenWeight
      token.minProfitBps, // _minProfitBps
      expandDecimals(token.maxUsdgAmount, 18), // _maxUsdgAmount
      token.isStable, // _isStable
      token.isShortable // _isShortable
    ), `vault.setTokenConfig(${token.name}) ${token.address}`)
  }
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
