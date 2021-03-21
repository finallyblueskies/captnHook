package coinbase

import (
	"github.com/bareish/captnHook/pkg/services"
	"github.com/fabioberger/coinbase-go"
	"log"
)

// BrokerService represent the interface to the Coinbase API
type BrokerService struct {
	Client coinbase.Client
	ConfigService services.ConfigService
}

// Setup ...
func (b *BrokerService) Setup() {
	coinbaseConfig := b.ConfigService.Get().Coinbase
	// client id
	clientID := coinbaseConfig.ClientID
	// client secret
	clientSecret := coinbaseConfig.ClientSecret
	// set client
	b.Client = coinbase.ApiKeyClient(clientID, clientSecret)
	// get balance
	balance, _ := b.Client.GetBalance()
	log.Println("Balance")
	log.Println(balance)
}