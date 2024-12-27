const { readTmpAddresses, contractAt, callWithRetries } = require("../shared/helpers.js")
const { expandDecimals } = require("../../test/shared/utilities.js")
const hre = require("hardhat");
const utils = require("../../utils.js");
const { ethers } = require("hardhat");


async function main() {
    const network = hre.network;
    const provider = new ethers.providers.JsonRpcProvider(network.config.url)
    
    const BTCaddress = utils.getYamlValue("BTC")
    const USDTaddress = utils.getYamlValue("USDT")
    const user1key = utils.getYamlValue("TestUser1Key")
    const user2key = utils.getYamlValue("TestUser2Key")
    const Vaultaddress = utils.getYamlValue("Vault")
    const allowanceNum = "1000" // *10^18
    for (const userkey of [user1key, user2key]) {
        for (const [name, token] of Object.entries({
            "BTC": BTCaddress,
            "USDT": USDTaddress
          })) {
            const signer = new ethers.Wallet(userkey).connect(provider)            
            const Contract = await contractAt(name, token, signer)
            const amount = ethers.utils.parseUnits(allowanceNum, 18); // 假设代币有 18 位小数
            await callWithRetries(Contract.approve.bind(Contract), [Vaultaddress, amount])
            const AllowanceNum = await Contract.allowance(signer.address, Vaultaddress);
            console.log(`${signer.address} allowance ${name} :`, ethers.utils.formatUnits(AllowanceNum, 18));
        }
    }
}

main()

/*
    const Vault = await contractAt("Vault", Vaultaddress)
    const TestUser1Address = utils.getYamlValue("TestUser1Address")
    const TestUser2Address = utils.getYamlValue("TestUser2Address")

*/