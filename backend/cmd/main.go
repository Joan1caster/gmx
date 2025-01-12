package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
)

func main() {
	apiKey := "B9VBgVdsMjfRbAaXWaYe0BnwT1tEcBmgV0NRSTrsJadA8oPeXTVXpuVkpDWOLvKO"
	secretKey := "hZzOyAaiLCILYTCdCkA0L9c9KyxzvvIcs3lTsiybGlqwX28JR5hoCqQxrjmC0t0k"

	client := binance.NewClient(apiKey, secretKey)

	priceChan := make(chan float64)

	go fetchPrice(client, priceChan)

	handler(priceChan)
}

func fetchPrice(client *binance.Client, priceChan chan<- float64) {
	for {
		// Create a context
		ctx := context.Background()

		// Fetch the latest Bitcoin price
		prices, err := client.NewListPricesService().Symbol("BTCUSDT").Do(ctx)
		if err != nil {
			log.Printf("Error fetching price: %v", err)
			continue
		}

		if len(prices) > 0 {
			price, err := strconv.ParseFloat(prices[0].Price, 64)
			if err != nil {
				log.Printf("Error converting price to float64: %v", err)
				continue
			}
			priceChan <- price
		}

		// Wait for a minute before fetching the price again
		time.Sleep(1 * time.Second)
	}
}

func handler(priceChan <-chan float64) {
	for {
		select {
		case price := <-priceChan:
			fmt.Printf("Latest BTC Price: %f\n", price)
		}
	}
}
