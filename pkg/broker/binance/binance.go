package binance

import (
	"github.com/adshao/go-binance/v2"
	"github.com/bareish/captnHook/pkg/services"
)

// BrokerService represent the interface to the Binance API
type BrokerService struct {
	Client        *binance.Client
	ConfigService services.ConfigService
}

func (b *BrokerService) Setup() {
	// binanceConfig := b.ConfigService.Get().Binance
}