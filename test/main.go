package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}

	fmt.Println("Attempting to connect...")

	url := "wss://testnet.binance.vision/ws/btcusdt@trade"

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

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v\n", err)
			return
		}
		fmt.Printf("Received: %s\n", message)
	}
}
