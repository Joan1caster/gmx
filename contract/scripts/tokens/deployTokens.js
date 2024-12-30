const { deployContract, sendTxn, writeTmpAddresses, callWithRetries } = require("../shared/helpers")
const { expandDecimals } = require("../../test/shared/utilities")

async function main() {
  const addresses = {}
  addresses.BTC = (await callWithRetries(deployContract, ["BTC", ["Bitcoin", "BTC", 18, expandDecimals(1000, 18)]])).address
  addresses.USDT = (await callWithRetries(deployContract, ["USDT", ["Tether", "USDT", 18, expandDecimals(1000, 18)]])).address
  await callWithRetries(deployContract, ["WETH", ["WETH", "WETH", 18, expandDecimals(1000, 18)]])

  writeTmpAddresses(addresses)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
