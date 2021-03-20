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

// Setup initializes client and configuration variable
func (b *BrokerService) Setup() {
	// binance configuration vars
	binanceConfig := b.ConfigService.Get().Binance
	// client secret
	clientSecret := binanceConfig.ClientSecret
	// client id
	clientID := binanceConfig.ClientId
	// set client
	b.Client = binance.NewClient(clientID, clientSecret)
}