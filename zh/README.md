# GMX 核心功能开发

## 核心功能
- **流动性质押与借贷**：优先实现核心功能。
- **附加功能**（若时间和资源允许）：
  - GMX 代币质押
  - 手续费奖励
  - Swap 模块

## GMX 工作流程
### 合约级别执行：
- 铸造并上架交易对。

### 合约触发事件（由后端监听）：
- 开仓/平仓事件
- 流动性质押/赎回
- 市价单状态变更

### 后端触发的合约操作：
- 价格更新
- 限价单触发
- 强制平仓执行

## 主要模块
### 合约：
- **Vault.sol**: 资金管理与交易执行。
- **Router.sol**: 用户交互入口。
- **PositionManager.sol**: 仓位管理。
- **OrderBook.sol**: 订单簿管理。
- **FastPriceFeed.sol**: 快速价格更新。
- **PriceFeed.sol**: 聚合多个价格来源。
- **ChainlinkPriceFeed.sol**: Chainlink 接口。
- **GLP.sol**: LP 代币合约。
- **GlpManager.sol**: LP 管理逻辑。
- **RewardTracker.sol**: 奖励分发。

### 后端：
- **价格监控器**: 监控多个交易所价格（通过配置文件简单实现）。
- **价格更新器**: 调用 `FastPriceFeed` 更新价格。
- **价格验证器**: 验证价格有效性。
- **限价单监控**: 检查触发条件。
- **执行器**: 触发限价单执行。
- **订单状态更新**
- **仓位监控**: 计算维持保证金率。
- **清算执行器**: 触发强制平仓。
- **链上事件监听**
- **数据统计分析**: 用于统计分析（未开源）。
- **API 服务**: 用于集成（未开源）。

## 项目部署
### 前置条件：
- 安装 `npm`。

### 步骤：
1. 进入 `contract` 目录：
   ```bash
   cd contract
   ```

2. **安装依赖：**
   ```bash
   npm install -g npx
   npm i
   ```

3. **编译合约：**
   ```bash
   npx hardhat compile
   ```

4. **合约脚本：**
### 部署主要合约，地址自动记录在contract/CA.yaml，默认gov为anvil的第一个地址
   ```bash
   npx hardhat run scripts/core/deployVault.js --network anvil
   npx hardhat run scripts/core/deployOrderBook.js --network anvil
   npx hardhat run scripts/core/deployPositionManager.js --network anvil
   npx hardhat run scripts/tokens/deployTokens.js --network anvil   
   npx hardhat run scripts/access/deployTokenManager.js --network anvil
   
   ```
6. **功能性测试脚本：**
### 注意部署到anvil的合约在anvil重启后会自动清理需要重新部署
   ```bash
   # 铸造BTC和USDT给anvil的第二、第三个地址
   npx hardhat run scripts/test/mint.js --network anvil
   # 添加白名单
   npx hardhat run scripts/core/whitelistTokens.js --network anvil
   # 质押流动性
   npx hardhat run scripts/test/pledge.js --network anvil
      ```
**已完成的工作：**
   部署gmx核心合约，部署USDT和BTC并铸造给测试账户，授权给vault

**接下来的工作：**
   创建流动性池：    GlpManager：addLiquidity
   设定币价：     Vault：setPriceFeed
   配置交易参数：    Vault：setFees、setTokenConfig
   创建开仓请求：      PositionRouter：createIncreasePosition
   创建平仓请求：      PositionRouter：createDecreasePosition
   确认交易：     PositionRouter：executeIncreasePosition、executeDecreasePosition