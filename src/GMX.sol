// MyVaultDeployment.sol
pragma solidity ^0.8.0;

import "lib/gmx-contracts/contracts/core/Vault.sol";
import "lib/gmx-contracts/contracts/core/PositionManager.sol";

contract MyVault {
    // 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
    Vault public vault;

    constructor() {
        vault = new Vault();
    }
}

// contract MyPositionManager {
//     PositionManager public positionManager;

//     constructor() {
//         positionManager = new PositionManager();
//     }
// }