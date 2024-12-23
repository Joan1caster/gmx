// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {MyVault} from "../src/GMX.sol";
// 合约依赖
//      PositionManager
//          Router.sol
//          vault.sol
contract DeployScript is Script {
    MyVault public myvault;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        myvault = new MyVault();

        vm.stopBroadcast();
    }
}
