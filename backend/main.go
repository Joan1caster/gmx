package main

import (
	"fmt"
	"gmxBackend/models"
	"gmxBackend/repository"
	"gmxBackend/service"
)

func main() {
	fmt.Println("service running...")

	priceBuffer := models.NewPriceBuffer()

	priceRepo := repository.NewPriceRepo()

	priceService := service.NewPriceService(priceBuffer, priceRepo)
	newPrice := []models.Price{}
	go priceService.UpdatePrices(newPrice)
}
