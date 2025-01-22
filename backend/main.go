package main

import (
	"fmt"
	"gmxBackend/config"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/service"
)

func main() {
	fmt.Println("service running...")

	config.Init() // 加载配置文件

	rabbitmq.InitMQ()

	go service.GetPrice("btcusdt") // 订阅币安的BTC/USDT价格
	go service.HandlerOrderInfo()  // 订阅订单信息
	go service.HandlerPriceInfo()

	// go service.UpdatePriceToChain(priceChain) // 将价格更新到链上

	defer rabbitmq.Shutdown()
	select {}
}
