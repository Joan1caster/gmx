# GMX Contracts

## Install Dependencies
If npx is not installed yet:
`npm install -g npx`

Install packages:
`npm i`

## Compile Contracts
`npx hardhat compile`

## Run Tests
`npx hardhat test`

## deploy
`npx hardhat run scripts/core/deployVault.js --network anvil`

`npx hardhat run scripts/core/deployOrderBook.js --network anvil`

`npx hardhat run scripts/core/deployPositionManager.js --network anvil`

        