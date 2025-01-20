package main

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/service"
)

func main() {
	fmt.Println("service running...")

	config.Init() // 加载配置文件

	priceChain := make(chan string, 10) //储存BTC价格订阅

	go service.GetPrice("btcusdt", priceChain) // 订阅币安的BTC/USDT价格

	go service.UpdatePriceToChain(priceChain) // 将价格更新到链上

	select {} // 防止main退出
}
