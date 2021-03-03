package binance

import (
	"github.com/adshao/go-binance/v2"
	"github.com/bareish/captnHook/pkg/services"
)

// BrokerService
type BrokerService struct {
	Client *binance.Client
	ConfigService services.ConfigService
}

// Buy
func (b *BrokerService) Buy(ticker string) {

}

// Sell
func (b *BrokerService) Sell(ticker string) {

}

// GetBuyingPower
func (b *BrokerService) GetBuyingPower() {

}