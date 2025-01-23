package rabbitmq

import (
	"fmt"
	"log"
)

func InitMQ() {
	err := InitConnection("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// 创建生产者
	CreateProducer("PriceUpdater")
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	// defer PriceUpdater.Close()

	// 创建消费者
	Price_Order, err := CreateConsumer("OrderUpdater")
	if err != nil {
		log.Fatalf("Failed to create consumer1: %v", err)
	}

	Price_User, err := CreateConsumer("UserAssets")
	if err != nil {
		log.Fatalf("Failed to create consumer2: %v", err)
	}

	Price_Platform, err := CreateConsumer("UserAssets")
	if err != nil {
		log.Fatalf("Failed to create consumer2: %v", err)
	}

	// 绑定队列到交换机
	exchangeName := "default"

	err = Price_Order.BindQueue(exchangeName, "PriceUpdate", "PriceOrder")
	if err != nil {
		log.Fatalf("Failed to bind queue1: %v", err)
	}

	err = Price_User.BindQueue(exchangeName, "PriceUpdate", "PriceUser")
	if err != nil {
		log.Fatalf("Failed to bind queue2: %v", err)
	}

	err = Price_Platform.BindQueue(exchangeName, "PriceUpdate", "PricePlatform")
	if err != nil {
		log.Fatalf("Failed to bind queue2: %v", err)
	}

	// 设置消费者处理函数
	err = Price_Order.Consume(func(msg []byte) error {
		fmt.Printf("Price_Order received: %s\n", string(msg))
		return nil
	}, "PriceOrder")
	if err != nil {
		log.Fatalf("Failed to start consuming on consumer1: %v", err)
	}

	err = Price_User.Consume(func(msg []byte) error {
		fmt.Printf("Price_User received: %s\n", string(msg))
		return nil
	}, "PriceUser")
	if err != nil {
		log.Fatalf("Failed to start consuming on consumer2: %v", err)
	}

	err = Price_Platform.Consume(func(msg []byte) error {
		fmt.Printf("Price_Platform received: %s\n", string(msg))
		return nil
	}, "PricePlatform")
	if err != nil {
		log.Fatalf("Failed to start consuming on consumer2: %v", err)
	}
}
