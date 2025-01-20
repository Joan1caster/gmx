package main

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/service"
)

func main() {
	fmt.Println("service running...")

	config.Init()
	priceFeed := config.GetString("BTCPriceFeed")
	fmt.Printf("btc pricefeed ca: %s\n", priceFeed)

	// priceBuffer := models.NewPriceBuffer()

	// priceRepo := repository.NewPriceRepo()

	// go priceService.UpdatePrices(newPrice)

	priceChain := make(chan string, 10)

	go service.GetPrice("btcusdt", priceChain)

	for {
		select {
		case price, ok := <-priceChain:
			if !ok {
				fmt.Println("channel closed")
			}
			fmt.Println("price:", price)

		}
	}
}
