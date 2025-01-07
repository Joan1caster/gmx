const { readTmpAddresses, contractAt, callWithRetries, sendTxn} = require("../shared/helpers.js")
const { expandDecimals } = require("../../test/shared/utilities.js")
const utils = require("../../utils.js");
const BigNumber = require('bignumber.js');
const { log } = require("easy-table");
const { ethers } = require("hardhat");

const BTCaddress = utils.getYamlValue("BTC")
const USDTaddress = utils.getYamlValue("USDT")
const GLPaddress = utils.getYamlValue("GLP")
const network = hre.network;
const provider = new ethers.providers.JsonRpcProvider(network.config.url)
const user1key = utils.getYamlValue("Usr1Key")
const user2key = utils.getYamlValue("Usr2Key")
const GOVKey = utils.getYamlValue("GOVKey")
const Vaultaddress = utils.getYamlValue("Vault")
const Routeraddress = utils.getYamlValue("Router")
const TestUser1Address = utils.getYamlValue("Usr1Address")
const TestUser2Address = utils.getYamlValue("Usr2Address")
const GOVAddress = utils.getYamlValue("GOVAddress")
const _GLPManager = utils.getYamlValue("GlpManager")

const BTCPrice = ethers.BigNumber.from(96000);
const USDTPrice = ethers.BigNumber.from(1);
const USDGDecimals = 18;

async function approve(allowanceNum) {
	const amount = ethers.BigNumber.from(allowanceNum);
    for (const userkey of [user1key, user2key, GOVKey]) {
        for (const [tokenName, tokenAddress] of Object.entries({
            "BTC": BTCaddress,
            "USDT": USDTaddress
          })) {
            const signer = new ethers.Wallet(userkey).connect(provider)            
            const Contract = await contractAt(tokenName, tokenAddress, signer)
			const TokenDecimal = await Contract.decimals()
			const AllowanceNum = await Contract.allowance(signer.address, _GLPManager);
            console.log(`succeed ${signer.address} allowance ${tokenName} :`, ethers.utils.formatUnits(AllowanceNum, TokenDecimal));

            await callWithRetries(Contract.approve.bind(Contract), [_GLPManager, amount.mul(ethers.BigNumber.from(10).pow(TokenDecimal))])
			}
    }
}

async function mint(_amount) {
	const amount = ethers.BigNumber.from(_amount);
	for (const user of [TestUser1Address, TestUser2Address, GOVAddress]) {
		const usdt = await contractAt("USDT", USDTaddress)
		const usdtDecimals = await usdt.decimals()
		await callWithRetries(usdt.mint.bind(usdt), [user, amount.mul(ethers.BigNumber.from(10).pow(usdtDecimals))])
		const btc = await contractAt("BTC", BTCaddress)
		const btcDecimals = await btc.decimals()
		await callWithRetries(btc.mint.bind(btc), [user, amount.mul(ethers.BigNumber.from(10).pow(btcDecimals))])

		const balanceUSDT = await usdt.balanceOf(user);
    	console.log("Balance of USDT ", user, ":", ethers.utils.formatUnits(balanceUSDT, usdtDecimals));
		const balanceBTC = await usdt.balanceOf(user);
    	console.log("Balance of BTC ", user, ":", ethers.utils.formatUnits(balanceBTC, btcDecimals));
	}
}


async function updatePrice() {
	const _BTCPriceFeed = utils.getYamlValue("BTCPriceFeed")
	const _USDTPriceFeed = utils.getYamlValue("USDTPriceFeed")
	const BTCPriceFeed = await contractAt("BTCPriceFeed",_BTCPriceFeed)
	const USDTPriceFeed = await contractAt("USDTPriceFeed",_USDTPriceFeed)
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()
	const btc = await contractAt("BTC", BTCaddress)
	const btcDecimals = await btc.decimals()
	// 将数字乘以原本的decimal
	await sendTxn(BTCPriceFeed.setLatestAnswer(BTCPrice.mul(ethers.BigNumber.from(10).pow(btcDecimals))), "BTCPriceFeed.setLatestAnswer")
	console.log(`set btc price: ${BTCPrice} USD`)

	await sendTxn(USDTPriceFeed.setLatestAnswer(USDTPrice.mul(ethers.BigNumber.from(10).pow(usdtDecimals))), "USDTPriceFeed.setLatestAnswer")
	console.log(`set usdt price: ${USDTPrice} USD`)
}

async function pricefeed() {
	const _vaultPriceFeed = utils.getYamlValue("VaultPriceFeed")
	const BTCarge = [BTCaddress, true, false, false]
	const USDTarge = [USDTaddress, true, false, false]
	const vaultPriceFeed = await contractAt("VaultPriceFeed",_vaultPriceFeed)
	// 访问价格时将精度转化为1e30
	const BTCprice = await callWithRetries(vaultPriceFeed.getPrice.bind(vaultPriceFeed), BTCarge)
	console.log(`get BTC price ${ethers.utils.formatUnits(BTCprice, 30)}`)
	const USDTprice = await callWithRetries(vaultPriceFeed.getPrice.bind(vaultPriceFeed), USDTarge)
	console.log(`get USDT price ${USDTprice.div(ethers.utils.parseUnits("1",30))}`)
}

async function addLiquidity(num) {	
	const GlpManager = await contractAt("GlpManager", _GLPManager)
	const Glp = await contractAt("GLP", GLPaddress)
	await callWithRetries(GlpManager.setHandler.bind(GlpManager), [GOVAddress, true])
	for (const [tokenName, tokenAddress] of Object.entries({
		"BTC": BTCaddress,
		"USDT": USDTaddress
	  })) {
		const token = await contractAt(tokenName, tokenAddress)
		const Decimals = await token.decimals()
		const TokenNum = ethers.utils.parseUnits(num, Decimals)
		const minTokenNum = TokenNum.mul(95).div(100)
		const minGLPNum = minTokenNum
		const args = [GOVAddress, GOVAddress, tokenAddress, TokenNum, minTokenNum, minGLPNum]
		const tx = await callWithRetries(GlpManager.addLiquidityForAccount.bind(GlpManager), args)
		const receipt = await tx.wait()
		const event = receipt.events.find(e => e.event === 'AddLiquidity')
		if (event) {
			const glpAmount = event.args.mintAmount // 具体参数名需要查看合约事件定义
			console.log(`Got ${ethers.utils.formatUnits(glpAmount, USDGDecimals)} GLP`)
		}
	}	
}

async function increasePosition(minTokenNum, time) {
	// 输入USDT，跟随BTC，做多，市价单
	const signer = new ethers.Wallet(user1key).connect(provider)   
	const router = await contractAt("Router", Routeraddress, signer)	
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()	
	const minTokenNum_ = ethers.utils.parseUnits(minTokenNum.toString(), usdtDecimals);
	const _minOut = minTokenNum_.div(BTCPrice).mul(95).div(100); // 假设扣除fee还剩95%
	const _sizeDelta = _minOut.mul(time)
	const args = [[USDTaddress, BTCaddress], BTCaddress, minTokenNum_, _minOut, _sizeDelta, true, BTCPrice-1]
	await callWithRetries(router.increasePosition.bind(router), args)
	console.log("execute increase position")

}

async function main() {
    await mint("100000"); // 给anvil的第二个第三个用户铸造btc、usdt
    await approve("10000");// 将钱授权给gmx的vault
    await updatePrice();// 将btc和USDT的价格传入
    await pricefeed();// 测试传入的价格
	await addLiquidity("1000", "1"); // gov添加流动性，gov获取GLP
	await increasePosition("1000", 10) // 买入1000U的BTC，开10倍多仓
	process.exit(0);
}

main()