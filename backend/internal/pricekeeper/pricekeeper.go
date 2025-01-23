package pricekeeper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/oracle/btcpricefeed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PriceKeeper struct {
	client          *ethclient.Client
	auth            *bind.TransactOpts
	btcPriceFed     *btcpricefeed.BTCPriceFeed
	btcPriceSession *btcpricefeed.BTCPriceFeedSession
}

// Connect 建立连接并初始化所需组件
func BTCPriceFeedConnect() (*PriceKeeper, error) {
	keeper := &PriceKeeper{}

	// 连接以太坊客户端
	client, err := ethclient.Dial(config.GetString("NodeAddress"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	keeper.client = client

	// 设置私钥和认证
	privateKey, err := crypto.HexToECDSA(config.GetString("GOVKey"))
	if err != nil {
		return nil, fmt.Errorf("get gov key err: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	keeper.auth = auth

	// 初始化合约
	btcPriceFeddAddress := common.HexToAddress(config.GetString("BTCPriceFeed"))
	btcPriceFed, err := btcpricefeed.NewBTCPriceFeed(btcPriceFeddAddress, client)
	if err != nil {
		return nil, err
	}
	keeper.btcPriceFed = btcPriceFed

	// 设置session
	keeper.btcPriceSession = &btcpricefeed.BTCPriceFeedSession{
		Contract:     btcPriceFed,
		CallOpts:     bind.CallOpts{},
		TransactOpts: *auth,
	}

	return keeper, nil
}

// UpdatePrice 更新价格
func (keeper *PriceKeeper) UpdateBTCPrice(BTCPrice *big.Int) error {
	return nil
	_, err := keeper.btcPriceSession.SetLatestAnswer(BTCPrice)
	return err
}
