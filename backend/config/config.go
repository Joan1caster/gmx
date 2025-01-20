package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	V          *viper.Viper
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

func Init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	V = viper.New()
	V.SetConfigName("CA.yaml")
	V.SetConfigType("yaml")
	V.AddConfigPath("../contract")

	err = V.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	V.Set("NodeAddress", "http://127.0.0.1:8554")
}

func GetString(key string) string {
	return V.GetString(key)
}

func GetInt(key string) int {
	return V.GetInt(key)
}
