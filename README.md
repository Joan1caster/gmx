# GMX 订单引擎后端
基于事件驱动的链上订单执行系统，实现订单管理、仓位监控、资产风控等核心交易功能。
## ✨ 核心功能
- **事件监听**：实时监听链上 CreateIncreaseOrder/CreateDecreaseOrder 事件
- **价格校验**：滑点验证 + 手续费计算（支持动态费率配置）
- **订单执行**：多条件触发式订单处理引擎
- **结果反馈**：执行状态实时追踪（支持Webhook/Event推送）
## 🛠 服务架构
### OrderBookKeeper
1. 监听 CreateIncreaseOrder 事件并持久化存储
2. 价格条件校验与手续费计算
3. 订单执行引擎
4. 交易结果反馈机制
### 微服务模块
| 服务模块         | 核心能力                                                                 |
|------------------|--------------------------------------------------------------------------|
| **订单服务**     | 处理开/平仓请求，维护订单生命周期状态机                                  |
| **仓位服务**     | 实时仓位管理，PnL动态计算，自动爆仓检测                                  |
| **价格服务**     | 价格流处理，限价订单触发，市场数据聚合                                   |
| **风控服务**     | 实时风险度监控，交易限额控制，异常波动熔断机制                            |
## 🔄 交易流程
### 限价开仓流程
graph TD
    A[用户调用 createIncreaseOrder] --> B[OrderBookKeeper 监听事件]
    B --> C[订单服务创建挂单]
    C --> D{价格服务监控市场}
    D -- 达到触发价 --> E[执行引擎处理]
    E --> F[更新仓位/资产状态]
## 强制平仓流程
graph TD
    A[价格更新] --> B[风控服务风险校验]
    B --> C[生成DecreaseOrder]
    C --> D[执行引擎强平]
    D --> E[计算最终损益]
    E --> F[资产清算处理]

## 🚀 快速开始 
### 
### 安装依赖
    npm install -g npx
    need install foundry
### 部署合约
    npx hardhat run scripts/deploy/deployVault.js --network anvil
    npx hardhat run scripts/deploy/deployOrderbook.js --network anvil
    npx hardhat run scripts/deploy/deployPositionManager.js --network anvil

### 测试
    npx hardhat run scripts/test/mint.js --network anvil

### 启动keeper服务
    anvil --code-size-limit 100000000
    cd backend 
    go mod tidy
    go run .
## 🤝 贡献指南
欢迎提交PR并遵循以下规范：

新功能开发需包含单元测试
重大变更需更新接口文档
代码风格遵循Airbnb规范