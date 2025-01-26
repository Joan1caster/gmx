package main

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/internal/database"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/repository"
	"gmxBackend/service"
	"log"
)

func main() {
	fmt.Println("service running...")

	config.Init() // 加载配置文件
	config.LoadConfig("config/config.yaml")

	rabbitmq.InitMQ()

	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB()

	database.InitTable()

	OrderRepo := repository.NewOrderRepository(database.GetDB())
	PositionRepo := repository.NewPositionReposttory(database.GetDB())

	orderService := service.NewOrderService(OrderRepo, PositionRepo)
	positionService := service.NewPositionService(OrderRepo, PositionRepo)

	go service.GetPrice("btcusdt")           // 订阅币安的BTC/USDT价格
	go orderService.HandlerOrderInfo()       // 订阅挂单请求、删除已执行的订单
	go positionService.HandlerPositionInfo() // 订阅头寸创建，检查平仓条件
	go orderService.HandlerPriceInfo()       // 检查挂单执行条件

	defer rabbitmq.Shutdown()
	select {}
}
