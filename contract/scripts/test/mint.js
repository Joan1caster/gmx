const { readTmpAddresses, contractAt, callWithRetries } = require("../shared/helpers.js")
const { expandDecimals } = require("../../test/shared/utilities.js")
const utils = require("../../utils.js");

async function main() {
	const TestUser1Address = utils.getYamlValue("TestUser1Address")
	const TestUser2Address = utils.getYamlValue("TestUser2Address")
	const USDT = utils.getYamlValue("USDT")
	const BTC = utils.getYamlValue("BTC")
	const amount = expandDecimals(100000, 18)
	for (const user of [TestUser1Address, TestUser2Address]) {
		const usdt = await contractAt("USDT", USDT)
		await callWithRetries(usdt.mint.bind(usdt), [user, amount])
		const btc = await contractAt("BTC", BTC)
		await callWithRetries(btc.mint.bind(btc), [user, amount])

		const balanceUSDT = await usdt.balanceOf(user);
    	console.log("Balance of USDT ", user, ":", ethers.utils.formatUnits(balanceUSDT, 18));
		const balanceBTC = await usdt.balanceOf(user);
    	console.log("Balance of BTC ", user, ":", ethers.utils.formatUnits(balanceUSDT, 18));
	}
}

main()