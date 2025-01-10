# OrderBookKeeper

1.监听CreateIncreaseOrder事件，信息入库；
2.校验价格是否满足条件、计算手续费；
3.执行订单；
4.反馈执行结果（便于调试和后续扩展）