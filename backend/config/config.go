package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	AppConfig  Config
	configOnce sync.Once
)

type Config struct {
	Contract struct {
		OrderBook                string `yaml:"OrderBook"`
		Vault                    string `yaml:"Vault"`
		PositionManager          string `yaml:"PositionManager"`
		NFTcontractAddress       string `yaml:"NFTcontractAddress"`
		NFTMarketcontractAddress string `yaml:"NFTMarketcontractAddress"`
		ContractApiFile          string `yaml:"contractApiFile"`
		GOVAddress               string `yaml:"GOVAddress"`
	} `yaml:"contract"`
	Account struct {
		PrivateKey     string `yaml:"privatekey"`
		AccountAddress string `yaml:"address"`
		NodeAddress    string `yaml:"eth_rpc_url"`
	} `yaml:"account"`
}

// LoadConfig loads the configuration from the YAML file.
func LoadConfig(configPath string) error {
	var err error
	configOnce.Do(func() {
		var data []byte
		data, err = os.ReadFile(configPath)
		if err != nil {
			log.Printf("Error reading config file: %v", err)
			return
		}
		err = yaml.Unmarshal(data, &AppConfig)
		if err != nil {
			log.Printf("Error unmarshalling config file: %v", err)
			return
		}
	})
	return err
}
