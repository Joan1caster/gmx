package main

import (
	"fmt"
	config "gmxBackend/config"
	"log"
)

func main() {
	err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.AppConfig.Account)
}
