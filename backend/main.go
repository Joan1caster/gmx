package main

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/internal/database"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/service"
	"log"
)

func main() {
	fmt.Println("service running...")

	config.Init() // 加载配置文件
	config.LoadConfig("config/config.yaml")

	rabbitmq.InitMQ()
	// Connect to the database
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB()

	database.InitTable()

	go service.GetPrice("btcusdt")                // 订阅币安的BTC/USDT价格
	go service.HandlerOrderInfo(database.GetDB()) // 订阅订单信息
	// go service.HandlerPriceInfo(database.GetDB())

	// go service.UpdatePriceToChain(priceChain) // 将价格更新到链上

	defer rabbitmq.Shutdown()
	select {}
}
