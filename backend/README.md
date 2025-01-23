# OrderBookKeeper

1.监听CreateIncreaseOrder事件，信息入库；
2.校验价格是否满足条件、计算手续费；
3.执行订单；
4.反馈执行结果（便于调试和后续扩展）

# Service
## 订单服务 (OrderService)
处理限价/市价开仓请求
处理平仓请求
维护订单状态

## 仓位服务 (PositionService)
管理用户持仓信息
计算仓位PnL
处理爆仓逻辑

## 价格服务 (PriceService)
接收和处理价格更新
维护最新价格
触发限价订单

## 资产服务 (AssetService)
管理用户保证金
处理保证金冻结/解冻
手续费计算和扣除

## 风控服务 (RiskService)
爆仓监控
仓位风险度计算
交易限额控制

# 交互流程
## 限价开仓
用户-> createIncreaseOrder（OrderBook）
后端-> 监听事件
    -> 订单服务更新
    -> 价格服务监听
    -> 订单、仓位、资产服务更新

## 市价开仓(不做)
用户-> createIncreaseOrder（Router）
后端-> 监听事件
    -> 订单、仓位、资产服务更新

## 限价平仓(可选)
用户-> createDecreaseOrder（OrderBook）
后端-> 监听事件
    -> 订单服务更新
    -> 价格服务监听
    -> 订单、仓位、资产服务更新

## 市价平仓（不做）
用户-> createDecreaseOrder（Router）
后端-> 监听事件
    -> 订单、仓位、资产服务更新

## 强制平仓
后端-> 价格监听
    -> CreateDecrease(Router)
    -> 订单、仓位、资产服务更新
(盈利-平仓手续费-仓位费-提现金额)

优化点：
事件监听断点续传
订单二级缓存，独立缓存价格附近的订单
按交易量优先执行订单


未来优化提高：
微服务、分布式事务处理数据一致性、价格更新幅度告警、限流保护

性能指标描述：
系统支持的TPS
价格更新延迟
订单处理延迟