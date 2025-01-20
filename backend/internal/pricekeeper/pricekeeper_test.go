package pricekeeper

import (
	"gmxBackend/config"
	"testing"
)

func TestMain(t *testing.T) {
	config.LoadConfig("../../config/config.yaml")
	UpdatePrice(96002)
}
