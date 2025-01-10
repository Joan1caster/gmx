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
const OrderBookaddress = utils.getYamlValue("OrderBook")
const TestUser1Address = utils.getYamlValue("Usr1Address")
const TestUser2Address = utils.getYamlValue("Usr2Address")
const GOVAddress = utils.getYamlValue("GOVAddress")
const _GLPManager = utils.getYamlValue("GlpManager")

const BTCPrice = ethers.BigNumber.from(96000);
const USDTPrice = ethers.BigNumber.from(1);
const ExecutionFee = ethers.BigNumber.from("10000000000000000")

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
            await callWithRetries(Contract.approve.bind(Contract), [Routeraddress, amount.mul(ethers.BigNumber.from(10).pow(TokenDecimal))])
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


async function updatePrice(_BTCPrice, _USDTPrice) {
	const _BTCPriceFeed = utils.getYamlValue("BTCPriceFeed")
	const _USDTPriceFeed = utils.getYamlValue("USDTPriceFeed")
	const BTCPriceFeed = await contractAt("BTCPriceFeed",_BTCPriceFeed)
	const USDTPriceFeed = await contractAt("USDTPriceFeed",_USDTPriceFeed)
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()
	const btc = await contractAt("BTC", BTCaddress)
	const btcDecimals = await btc.decimals()
	// 将数字乘以原本的decimal
	await sendTxn(BTCPriceFeed.setLatestAnswer(_BTCPrice.mul(ethers.BigNumber.from(10).pow(btcDecimals))), "BTCPriceFeed.setLatestAnswer")
	console.log(`set btc price: ${BTCPrice} USD`)

	await sendTxn(USDTPriceFeed.setLatestAnswer(_USDTPrice.mul(ethers.BigNumber.from(10).pow(usdtDecimals))), "USDTPriceFeed.setLatestAnswer")
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
			console.log(`Got ${ethers.utils.formatUnits(glpAmount, Decimals)} GLP`)
		}
	}	
}

async function approvePlugin() {
	const routerGOV = await contractAt("Router", Routeraddress)
	await callWithRetries(routerGOV.addPlugin.bind(routerGOV), [OrderBookaddress])
	const signer = new ethers.Wallet(user1key).connect(provider)  
	const router = await contractAt("Router", Routeraddress, signer)	
	await callWithRetries(router.approvePlugin.bind(router), [OrderBookaddress])
	const isApproved = await router.approvedPlugins(TestUser1Address, OrderBookaddress);
	console.log("Is approved:", isApproved);  // 返回 true/false
}

async function increasePosition(amountIn, time) {
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()	
	const btc = await contractAt("BTC", BTCaddress)
	const btcDecimals = await btc.decimals()	
	const amount = ethers.BigNumber.from(amountIn)
	const _amountIn = ethers.utils.parseUnits(amountIn, usdtDecimals)
	const _minOut = amount.div(BTCPrice).mul(10).pow(btcDecimals).mul(95).div(100);
	const _sizeDelta = amount.mul(time).mul(ethers.BigNumber.from(10).pow(30))	
	const signer = new ethers.Wallet(user1key).connect(provider) 
	const orderbook = await contractAt("OrderBook", OrderBookaddress, signer)
	const _BTCPrice = BTCPrice.mul(ethers.BigNumber.from(10).pow(30))
	const arg = [[USDTaddress, BTCaddress], _amountIn, BTCaddress, _minOut, _sizeDelta, BTCaddress, true, _BTCPrice, true, ExecutionFee, false, {value:ExecutionFee}]
	const tx = await callWithRetries(orderbook.createIncreaseOrder.bind(orderbook), arg)
	console.log(`create increase order succeed.`)
}

async function ExecuteIncreaseOrder() {
	const orderbook = await contractAt("OrderBook", OrderBookaddress)
	let orderIndex = await orderbook.increaseOrdersIndex(TestUser1Address)
	orderIndex = orderIndex.add(-1)
	await callWithRetries(orderbook.executeIncreaseOrder.bind(orderbook), [TestUser1Address, orderIndex, GOVAddress])
	console.log(`Execute Increase Order succeed.`)
}

async function DecreasePosition(sizeDelta, _collateralDelta) {
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()	
	const signer = new ethers.Wallet(user1key).connect(provider) 
	const orderbook = await contractAt("OrderBook", OrderBookaddress, signer)
	const _sizeDelta = ethers.utils.parseUnits(sizeDelta, 30)
	const collateralDelta = ethers.utils.parseUnits(_collateralDelta, usdtDecimals)
	const _BTCPrice = BTCPrice.mul(ethers.BigNumber.from(10).pow(30))
	// 低于目标价格触发
	arg = [BTCaddress, _sizeDelta, BTCaddress, collateralDelta, true, _BTCPrice, false, {value:ExecutionFee+1}]
	await callWithRetries(orderbook.createDecreaseOrder.bind(orderbook), arg)
	console.log("create decrease order succeed")
	
}

async function ExecuteDecreaseOrder() {
	const orderbook = await contractAt("OrderBook", OrderBookaddress)
	let orderIndex = await orderbook.decreaseOrdersIndex(TestUser1Address)
	orderIndex = orderIndex.add(-1)
	await callWithRetries(orderbook.executeDecreaseOrder.bind(orderbook), [TestUser1Address, orderIndex, GOVAddress])
	console.log(`Execute Decrease Order succeed.`)
}

async function marketPriceIncrease(amountIn, time) {
	const btc = await contractAt("BTC", BTCaddress)
	const btcDecimals = await btc.decimals()
	const signer = new ethers.Wallet(user1key).connect(provider) 
	const router = await contractAt("Router", Routeraddress, signer)
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()	
	const _amountIn = ethers.utils.parseUnits(amountIn, usdtDecimals)
	const amount = ethers.BigNumber.from(amountIn)
	const _minOut = amount.div(BTCPrice).mul(10).pow(btcDecimals).mul(95).div(100)
	const _sizeDelta = amount.mul(time).mul(ethers.BigNumber.from(10).pow(30))
	const _BTCPrice = BTCPrice.mul(ethers.BigNumber.from(10).pow(30))
	arg = [[USDTaddress, BTCaddress], BTCaddress, _amountIn, _minOut, _sizeDelta, true, _BTCPrice.add(1)]
	await callWithRetries(router.increasePosition.bind(router), arg)
	console.log("execulte market Price Increase succeed.")
}

async function marketPriceDecrease(sizeDelta, _collateralDelta) {
	const usdt = await contractAt("USDT", USDTaddress)
	const usdtDecimals = await usdt.decimals()	
	const signer = new ethers.Wallet(user1key).connect(provider) 
	const router = await contractAt("Router", Routeraddress, signer)
	const _sizeDelta = ethers.utils.parseUnits(sizeDelta, 30)
	const collateralDelta = ethers.utils.parseUnits(_collateralDelta, usdtDecimals)
	const _BTCPrice = BTCPrice.mul(ethers.BigNumber.from(10).pow(30))
	// 低于目标价格触发
	arg = [BTCaddress, BTCaddress, collateralDelta, _sizeDelta, true, TestUser1Address, _BTCPrice.sub(20)]
	await callWithRetries(router.decreasePosition.bind(router), arg)
	console.log("execulte market Price Decrease succeed.")
	
}

async function getLatestEvents(contract) {
	// 用于输出制定合约的事件，调试用
    const latestBlock = await provider.getBlockNumber();
    
    const filter = contract.filters.Test();
    
    const events = await contract.queryFilter(filter, latestBlock, latestBlock);
	console.log(`latestBlock:${latestBlock}, events num: ${events.length}`)
    events.forEach(event => {
		const bignum = new BigNumber(event.args.arg.toString())
        console.log('Event Args:', {
            arg: bignum.toExponential(),
            name: event.args.name
        });
    });
}

async function main() {
	// 提供流动性、余额、授权等，只执行一次即可
    await mint("100000"); // 给anvil的第1、2、3个用户铸造btc、usdt
    await approve("10000");// 将钱授权给gmx的router(用作杠杆交易)/glpmanager(用作质押)
    await updatePrice(BTCPrice, USDTPrice);// 将btc和USDT的价格传入,单位为原token单位，精度为pricefeed精度
	await addLiquidity("1000", "1"); // gov添加流动性，gov获取GLP
	await approvePlugin() // gov授权plugin，用户授权plugin

	// // 限价单：1k USDT开10倍BTC的多仓
	// await increasePosition("1000", 10)// user1买入1000U的BTC，开10倍多仓
	// await updatePrice(BTCPrice.add(1),USDTPrice);// 假设BTC价格上涨了1美元
	// await ExecuteIncreaseOrder()// 执行开辟多仓

	// // 限价单：将上一步开的持仓量减小一半，押金退500USD，借贷规模减小5k USD
	// // 如果要完全平仓，将借贷规模设定成全部，押金写0即可，自动清算
	// await DecreasePosition("5000", "500") // 退出5000借贷规模，以及500押金
	// await updatePrice(BTCPrice.sub(1),USDTPrice);// 假设BTC价格下跌了1美元
	// await ExecuteDecreaseOrder()// 执行减仓

	// 市价单：1k USDT开10倍BTC的多仓
	await marketPriceIncrease("1000", 10)

	// 市价单：平仓一半
	await marketPriceDecrease("5000", "500")
	process.exit(0);
}

main()


