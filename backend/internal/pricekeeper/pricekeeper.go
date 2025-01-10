package pricekeeper

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func getPositionQueueLengths() {
	return
}

func updatePriceBitsAndOptionallyExecute(client *rpc.Client, configService ConfigService, accountAddress common.Address, gasLimit uint64) {
	priceFeedAddress := common.HexToAddress(configService.Get("contracts.fastPriceFeed"))
	priceFeedContract := NewPriceFeed(priceFeedAddress, client)

	positionQueue, err := getPositionQueueLengths()
	if err != nil {
		log.Fatalf("failed to get position queue lengths: %v", err)
	}
	timestamp := uint64(time.Now().Unix())

	if (positionQueue.IncreaseKeysLength-positionQueue.IncreaseKeyStart > 0) ||
		(positionQueue.DecreaseKeysLength-positionQueue.DecreaseKeyStart > 0) {

		endIndexForIncreasePositions := positionQueue.IncreaseKeysLength
		endIndexForDecreasePositions := positionQueue.DecreaseKeysLength

		log.Println("setPricesWithBitsAndExecute")

		tokens := getTokens()
		prices, priceBits, err := fetchPriceBits(tokens)
		if err != nil {
			log.Fatalf("failed to fetch price bits: %v", err)
		}

		tx, err := priceFeedContract.SetPricesWithBitsAndExecute(
			common.HexToAddress(configService.Get("contracts.positionRouter")),
			priceBits,
			timestamp,
			endIndexForIncreasePositions,
			endIndexForDecreasePositions,
			maxIncreasePositions,
			maxDecreasePositions,
		)
		if err != nil {
			log.Fatalf("failed to set prices with bits and execute: %v", err)
		}

		log.Println("setPricesWithBitsAndExecute", tx.Hash().Hex())

	} else {
		log.Println("setPricesWithBits")

		if time.Since(lastPriceUpdateTimestamp) > 2*time.Minute {
			prices, priceBits, err := fetchPriceBits(tokens)
			if err != nil {
				log.Fatalf("failed to fetch price bits: %v", err)
			}

			tx, err := priceFeedContract.SetPricesWithBits(priceBits, timestamp)
			if err != nil {
				log.Fatalf("failed to set prices with bits: %v", err)
			}

			log.Println("setPricesWithBits", tx.Hash().Hex())
			lastPriceUpdateTimestamp = time.Now()
		} else {
			log.Println("skipUpdatePrice")
		}
	}
}

// Mocked function definitions and structs for completeness
type ConfigService struct{}

func (c ConfigService) Get(key string) string { return "" }

type PositionQueue struct {
	IncreaseKeysLength int
	IncreaseKeyStart   int
	DecreaseKeysLength int
	DecreaseKeyStart   int
}

func getPositionQueueLengths() (PositionQueue, error) {
	return PositionQueue{}, nil
}

type Token struct{}

func getTokens() []Token {
	return []Token{}
}

func fetchPriceBits(tokens []Token) ([]int, []int, error) {
	return []int{}, []int{}, nil
}

type PriceFeed struct{}

func NewPriceFeed(address common.Address, client *rpc.Client) *PriceFeed {
	return &PriceFeed{}
}

func (p *PriceFeed) SetPricesWithBitsAndExecute(address common.Address, priceBits []int, timestamp uint64, endIndexForIncrease, endIndexForDecrease, maxIncrease, maxDecrease int) (*rpc.Transaction, error) {
	return nil, nil
}

func (p *PriceFeed) SetPricesWithBits(priceBits []int, timestamp uint64) (*rpc.Transaction, error) {
	return nil, nil
}

var lastPriceUpdateTimestamp time.Time
var maxIncreasePositions int
var maxDecreasePositions int

func main() {
	return
}
