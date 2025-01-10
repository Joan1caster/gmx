const { deployContract, contractAt , sendTxn, writeTmpAddresses } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")
const utils = retuire("../../utils.js")
const network = (process.env.HARDHAT_NETWORK || 'mainnet');
// const tokens = require('./tokens')[network];

async function main() {
  // const { nativeToken } = tokens

  const orderBook = await deployContract("OrderBook", []);

  // Arbitrum mainnet addresses
  await sendTxn(orderBook.initialize(
    utils.getYamlValue("Router"), // router
    utils.getYamlValue("Vault"), // vault
    utils.getYamlValue("WETH"), // weth
    utils.getYamlValue("USDG"), // usdg
    "10000000000000000", // 0.01 ETH
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
