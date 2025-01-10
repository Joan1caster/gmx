# GMX Contracts

## Install Dependencies
If npx is not installed yet:
`npm install -g npx`

Install packages:
`npm i`
## network
    need install foundry
    ```bash
    anvil --code-size-limit 100000000
    ```

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
