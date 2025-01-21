package service

import (
	"encoding/json"
	"fmt"
	rabbitmq "gmxBackend/middleware/mq"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func GetPrice(Tokens string) {
	producer := rabbitmq.Producers["PriceUpdater"]
	dialer := websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		// HandshakeTimeout: 5 * time.Second,
	}

	fmt.Println("Attempting to connect...")

	url := fmt.Sprintf("wss://testnet.binance.vision/ws/%s@trade", Tokens)

	c, resp, err := dialer.Dial(url, nil)
	if err != nil {
		log.Printf("Dial error: %v\n", err)
		if resp != nil {
			log.Printf("HTTP Response: %+v\n", resp)
		}
		log.Fatal("Cannot connect to WebSocket")
	}
	defer c.Close()

	fmt.Println("Connected to Binance WebSocket")

	latestPrice := make(map[string]string)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v\n", err)
			return
		}
		var messageJson map[string]interface{}
		err = json.Unmarshal(message, &messageJson)
		if err != nil {
			fmt.Println("decode price message error")
		}

		messageStr := messageJson["p"].(string)
		if latestPrice[Tokens] == messageStr {
			continue
		}
		latestPrice[Tokens] = messageStr
		// mq.Publish([]byte(messageStr), "First", "direct", "PriceUpdate")

		err = producer.Publish([]byte(messageStr), "default", "direct", "PriceUpdate")
		if err != nil {
			log.Printf("Failed to publish to consumer1: %v", err)
		}
	}
}

// func UpdatePriceToChain(priceChain chan string, mq *rabbitmq.RabbitMQ) {
// 	BTCPriceConn, err := pricekeeper.BTCPriceFeedConnect()
// 	priceConsumer := mq.GetConsumer("")
// 	if err != nil {
// 		fmt.Println("connect rpc error:", err)
// 	}
// 	for {
// 		select {
// 		case price, ok := <-priceChain:
// 			if !ok {
// 				fmt.Println("channel closed")
// 			}
// 			priceUint256 := utils.StringToUint256(price, 18)
// 			BTCPriceConn.UpdateBTCPrice(priceUint256)
// 			fmt.Println("BTC price for now:", price)

// 		}
// 	}
// }
