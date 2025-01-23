package rabbitmq

import (
	"fmt"
	"gmxBackend/utils"
	"testing"
)

func TestInitConnection(t *testing.T) {
	InitMQ()
	order, err := GetConsumer("OrderUpdater")
	if err != nil {
		t.Errorf("Failed to get consumer: %v", err)
	}
	for {
		order.Consume(func(msg []byte) error {
			fmt.Printf("OrderUpdater received: %s\n", utils.StringToUint256(string(msg), 18))
			return nil
		}, "PriceOrder")
	}
}
