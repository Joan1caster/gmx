package main

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/internal/pricekeeper"
	"gmxBackend/service"
	"gmxBackend/utils"
)

func main() {
	fmt.Println("service running...")

	config.Init()
	priceFeed := config.GetString("BTCPriceFeed")
	fmt.Printf("btc pricefeed ca: %s\n", priceFeed)

	// priceBuffer := models.NewPriceBuffer()

	// priceRepo := repository.NewPriceRepo()

	// go priceService.UpdatePrices(newPrice)
	BTCPriceConn, err := pricekeeper.BTCPriceFeedConnect()
	if err != nil {
		fmt.Println("connect rpc error:", err)
	}

	priceChain := make(chan string, 10)

	go service.GetPrice("btcusdt", priceChain)

	for {
		select {
		case price, ok := <-priceChain:
			if !ok {
				fmt.Println("channel closed")
			}
			priceUint256 := utils.StringToUint256(price, 18)
			BTCPriceConn.UpdateBTCPrice(priceUint256)
			fmt.Println("price:", price)

		}
	}
}
