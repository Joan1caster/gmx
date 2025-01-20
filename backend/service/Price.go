package service

import (
	"encoding/json"
	"fmt"
	"gmxBackend/internal/pricekeeper"
	"gmxBackend/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func GetPrice(Tokens string, priceChan chan string) {
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
		if latestPrice[Tokens] == messageJson["p"].(string) {
			continue
		}
		latestPrice[Tokens] = messageJson["p"].(string)
		priceChan <- messageJson["p"].(string)
	}
}

func UpdatePriceToChain(priceChain chan string) {
	BTCPriceConn, err := pricekeeper.BTCPriceFeedConnect()
	if err != nil {
		fmt.Println("connect rpc error:", err)
	}
	for {
		select {
		case price, ok := <-priceChain:
			if !ok {
				fmt.Println("channel closed")
			}
			priceUint256 := utils.StringToUint256(price, 18)
			BTCPriceConn.UpdateBTCPrice(priceUint256)
			fmt.Println("BTC price for now:", price)

		}
	}
}
