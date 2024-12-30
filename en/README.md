# GMX Core Functionality Development

## Core Features
- **Liquid Staking and Lending**: Implement the core functionalities first.
- **Additional Features** (time and resources permitting):
  - GMX token staking
  - Fee rewards
  - Swap module

## GMX Workflow
### Contract-Level Execution:
- Mint and list trading pairs.

### Events Triggered by Contracts (monitored by backend):
- Opening/closing positions
- Staking/redeeming liquidity
- Market order status updates

### Backend-Initiated Contract Actions:
- Price updates
- Limit order triggers
- Forced liquidation execution

## Primary Modules
### Contracts:
- **Vault.sol**: Manages funds and executes trades.
- **Router.sol**: User interaction entry point.
- **PositionManager.sol**: Position management.
- **OrderBook.sol**: Order book management.
- **FastPriceFeed.sol**: Fast price updates.
- **PriceFeed.sol**: Aggregates multiple price sources.
- **ChainlinkPriceFeed.sol**: Interface for Chainlink.
- **GLP.sol**: LP token contract.
- **GlpManager.sol**: LP management logic.
- **RewardTracker.sol**: Reward distribution.

### Backend:
- **Price Monitor**: Monitors prices across exchanges (implemented with a simple configuration file).
- **Price Updater**: Calls `FastPriceFeed` to update prices.
- **Price Validator**: Validates price accuracy.
- **Limit Order Monitor**: Checks trigger conditions.
- **Executor**: Executes triggered limit orders.
- **Order Status Updater**
- **Position Monitor**: Calculates margin maintenance rates.
- **Liquidation Executor**: Triggers forced liquidations.
- **On-Chain Event Listener**
- **Data Analytics**: For statistics and analysis (not open-sourced).
- **API Services**: For integration (not open-sourced).

## Project Deployment
### Prerequisites:
- Install `npm`.

### Steps:
1. Navigate to the `contract` directory:
   ```bash
   cd contract
   ```

2. **Install Dependencies:**
   ```bash
   npm install -g npx
   npm i
   ```

3. **Compile Contracts:**
   ```bash
   npx hardhat compile
   ```

4. **Run Tests -- Cost Time :**
   ```bash
   npx hardhat test
   ```

5. **Deploy Contracts:**
   ```bash
   npx hardhat run scripts/core/deployVault.js --network anvil
   npx hardhat run scripts/core/deployOrderBook.js --network anvil
   npx hardhat run scripts/core/deployPositionManager.js --network anvil
   ```

