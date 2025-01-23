package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
)

func GetClient() (*ethclient.Client, error) {
	var err error
	client, err = ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	return client, nil
}
