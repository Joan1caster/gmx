const { readTmpAddresses, contractAt, callWithRetries, sendTxn} = require("../shared/helpers.js")
const { expandDecimals } = require("../../test/shared/utilities.js")
const utils = require("../../utils.js");
const BigNumber = require('bignumber.js');
const { log } = require("easy-table");

const BTCaddress = utils.getYamlValue("BTC")
const USDTaddress = utils.getYamlValue("USDT")
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

const BTCPriceInt = new BigNumber('96290e18');
const BTCPrice = ethers.BigNumber.from(BTCPriceInt.toFixed());

async function approve(allowanceNum) {
	const mintAmount = new BigNumber(allowanceNum);
	const amount = ethers.BigNumber.from(mintAmount.toFixed());
    for (const userkey of [user1key, user2key, GOVKey]) {
        for (const [tokenName, tokenAddress] of Object.entries({
            "BTC": BTCaddress,
            "USDT": USDTaddress
          })) {
            const signer = new ethers.Wallet(userkey).connect(provider)            
            const Contract = await contractAt(tokenName, tokenAddress, signer)
			const AllowanceNum = await Contract.allowance(signer.address, _GLPManager);
            console.log(`succeed ${signer.address} allowance ${tokenName} :`, ethers.utils.formatUnits(AllowanceNum, 18));

            await callWithRetries(Contract.approve.bind(Contract), [_GLPManager, amount])
			}
    }
}

async function mint(_amount) {
	const mintAmount = new BigNumber(_amount);
	const amount = ethers.BigNumber.from(mintAmount.toFixed());
	for (const user of [TestUser1Address, TestUser2Address, GOVAddress]) {
		const usdt = await contractAt("USDT", USDTaddress)
		await callWithRetries(usdt.mint.bind(usdt), [user, amount])
		const btc = await contractAt("BTC", BTCaddress)
		await callWithRetries(btc.mint.bind(btc), [user, amount])

		const balanceUSDT = await usdt.balanceOf(user);
    	console.log("Balance of USDT ", user, ":", ethers.utils.formatUnits(balanceUSDT, 18));
		const balanceBTC = await usdt.balanceOf(user);
    	console.log("Balance of BTC ", user, ":", ethers.utils.formatUnits(balanceUSDT, 18));
	}
}


async function updatePrice() {
	const _BTCPriceFeed = utils.getYamlValue("BTCPriceFeed")
	const _USDTPriceFeed = utils.getYamlValue("USDTPriceFeed")
	const BTCPriceFeed = await contractAt("BTCPriceFeed",_BTCPriceFeed)
	const USDTPriceFeed = await contractAt("USDTPriceFeed",_USDTPriceFeed)
	// 将数字原封不动存进去
	await sendTxn(BTCPriceFeed.setLatestAnswer(BTCPrice), "BTCPriceFeed.setLatestAnswer")
	console.log(`set btc price: ${BTCPrice} USD`)

	const USDTPriceInt = new BigNumber('1e18');
	const USDTPrice = ethers.BigNumber.from(USDTPriceInt.toFixed());
	await sendTxn(USDTPriceFeed.setLatestAnswer(USDTPrice), "USDTPriceFeed.setLatestAnswer")
	console.log(`set usdt price: ${USDTPrice} USD`)
}

async function pricefeed() {
	const _vaultPriceFeed = utils.getYamlValue("VaultPriceFeed")
	const BTCarge = [BTCaddress, true, false, false]
	const USDTarge = [USDTaddress, true, false, false]
	const vaultPriceFeed = await contractAt("VaultPriceFeed",_vaultPriceFeed)
	const vault = await contractAt("Vault", Vaultaddress)
	const BTCprice = await callWithRetries(vault.getMinPrice.bind(vault), [BTCaddress])

	// 访问价格时将精度转化为1e30
	// const BTCprice = await callWithRetries(vaultPriceFeed.getPrice.bind(vaultPriceFeed), BTCarge)
	console.log(`get BTC price ${BTCprice/1e30}`)
	const USDTprice = await callWithRetries(vaultPriceFeed.getPrice.bind(vaultPriceFeed), USDTarge)
	console.log(`get USDT price ${USDTprice/1e30}`)
}

async function addLiquidity(num, min_num) {	
	const GlpManager = await contractAt("GlpManager", _GLPManager)
	await callWithRetries(GlpManager.setHandler.bind(GlpManager), [GOVAddress, true])
	const _TokenNum = new BigNumber(num);
	const TokenNum = ethers.BigNumber.from(_TokenNum.toFixed());
	const _minTokenNum = new BigNumber(min_num);
	const minTokenNum = ethers.BigNumber.from(_minTokenNum.toFixed());
	for (const token of [USDTaddress, BTCaddress]) {
		const args = [GOVAddress, GOVAddress, token, TokenNum, minTokenNum, minTokenNum]
		const tx = await callWithRetries(GlpManager.addLiquidityForAccount.bind(GlpManager), args)
		const receipt = await tx.wait()
		const event = receipt.events.find(e => e.event === 'AddLiquidity')
		if (event) {
			const glpAmount = event.args.mintAmount // 具体参数名需要查看合约事件定义
			console.log(`Got ${ethers.utils.formatUnits(glpAmount, 18)} GLP`)
		} 
	}	
}

async function increasePosition(minTokenNum, time) {
	// 输入USDT，跟随BTC，做多，市价单
	const signer = new ethers.Wallet(user1key).connect(provider)   
	const router = await contractAt("Router", Routeraddress, signer)		
	const minTokenNum_ = ethers.BigNumber.from(new BigNumber(minTokenNum).toFixed());
	const _minOut = Math.floor(minTokenNum_ / BTCPrice);
	const _sizeDelta = _minOut * time
	const args = [[USDTaddress, BTCaddress], BTCaddress, minTokenNum_, _minOut, _sizeDelta, true, BTCPrice-1]
	await callWithRetries(router.increasePosition.bind(router), args)

//Router.sol:     function increasePosition(address[] memory _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price)
}

async function main() {
    // await mint("10000e18"); // 给anvil的第二个第三个用户铸造btc、usdt
    // await approve("2000e18");// 将钱授权给gmx的vault
    // await updatePrice();// 将btc和USDT的价格传入
    // await pricefeed();// 测试传入的价格
	// await addLiquidity("1000e18", "1"); // gov添加流动性，gov获取GLP
	await increasePosition("1000e16", 10) 
	process.exit(0);
}

main()