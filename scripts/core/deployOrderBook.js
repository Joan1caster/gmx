const { deployContract, contractAt , sendTxn, writeTmpAddresses } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const utils = require("../../utils.js");
const network = (process.env.HARDHAT_NETWORK || 'mainnet');
const tokens = require('./tokens')[network];

async function main() {
  const { nativeToken } = tokens

  const orderBook = await deployContract("OrderBook", []);

  const router = utils.getYamlValue("Router")
  const vault = utils.getYamlValue("Vault")
  const usdg = utils.getYamlValue("USDG")
  // Arbitrum mainnet addresses
  await sendTxn(orderBook.initialize(
    router, // router
    vault, // vault
    nativeToken.address, // weth
    usdg, // usdg
    "10000000000000000", // 0.01 AVAX
    expandDecimals(10, 30) // min purchase token amount usd
  ), "orderBook.initialize");

  writeTmpAddresses({
    orderBook: orderBook.address
  })
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
