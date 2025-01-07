# GMX Contracts
Contracts for GMX.

Docs at https://gmxio.gitbook.io/gmx/contracts.

## Install Dependencies
If npx is not installed yet:
`npm install -g npx`

Install packages:
`npm i`

## deploy 
    ```bash
    npx hardhat run scripts/deploy/deployVault.js --network anvil
    npx hardhat run scripts/deploy/deployOrderbook.js --network anvil
    npx hardhat run scripts/deploy/deployPositionManager.js --network anvil

    ```

## interact
    ```bash
    npx hardhat run scripts/test/mint.js --network anvil
    ```

## recompile
    删除 .\cache 

    删除 .\artifacts 

    重新编译 npx hardhat compile
