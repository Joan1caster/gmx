package pricekeeper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/oracle/btcpricefeed"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func UpdatePrice(BTCPrice float64) {
	btcPriceFeddAddress := common.HexToAddress(config.AppConfig.Contract.BtcPriceFeddAddress)

	client, err := ethclient.Dial(config.AppConfig.Account.NodeAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(config.AppConfig.Account.PrivateKey)
	fmt.Println(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	btcPriceFed, _ := btcpricefeed.NewBTCPriceFeed(btcPriceFeddAddress, client)

	btcPriceSession := btcpricefeed.BTCPriceFeedSession{
		Contract:     btcPriceFed,
		CallOpts:     bind.CallOpts{},
		TransactOpts: *auth,
	}
	btcdecimals, _ := btcPriceFed.Decimals(&bind.CallOpts{})

	BTCPriceBig := big.NewFloat(BTCPrice)

	BTCPriceBigInt, _ := BTCPriceBig.SetMantExp(BTCPriceBig, BTCPriceBig.MantExp(nil)+int(btcdecimals.Int64())).Int64()

	btcPriceSession.SetLatestAnswer(big.NewInt(BTCPriceBigInt))
}
