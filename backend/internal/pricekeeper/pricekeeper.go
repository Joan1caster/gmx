package pricekeeper

import (
	"gmxBackedn/contracts/oracle/btcpricefeed"
	"gmxBackend/config"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func UpdatePrice(BTCPrice float64) {
	btcPriceFeddAddress := common.HexToAddress(config.AppConfig.Contracts.BTCPriceFeed)

	client, err := ethclient.Dial(config.AppConfig.Account.NodeAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	btcPriceFed := btcpricefeed.NewBTCPriceFeed(btcPriceFeddAddress, client)

	btcdecimals := btcPriceFed.Decimals()

	btcPriceFed.SetLatestAnswer(big.NewInt(int64(BTCPrice) * btcdecimals))
}
