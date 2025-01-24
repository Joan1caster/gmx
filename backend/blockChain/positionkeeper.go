package blockChain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/vault"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PositionKeeper struct {
	client       *ethclient.Client
	auth         *bind.TransactOpts
	vault        *vault.Vault
	vaultSession *vault.VaultSession
}

// Connect 建立连接并初始化所需组件
func VaultConnect() (*PositionKeeper, error) {
	keeper := &PositionKeeper{}

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
	vaultAddress := common.HexToAddress(config.GetString("Vault"))
	vault_, err := vault.NewVault(vaultAddress, client)
	if err != nil {
		return nil, err
	}
	keeper.vault = vault_

	// 设置session
	keeper.vaultSession = &vault.VaultSession{
		Contract:     vault_,
		CallOpts:     bind.CallOpts{},
		TransactOpts: *auth,
	}

	return keeper, nil
}

// UpdatePrice 更新价格
func (keeper *PositionKeeper) LiquidatePosition(_account, _collateralToken, _indexToken common.Address, _isLone bool, _feeReceiver common.Address) error {
	_, err := keeper.vaultSession.LiquidatePosition(_account, _collateralToken, _indexToken, _isLone, _feeReceiver)
	return err
}
