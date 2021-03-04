package alpaca

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/bareish/captnHook/pkg/services"
	"os"
)

const (
	// NoFundsErr is an error message if there is no money in the account
	NoFundsErr = "no funds available in account"
	// PositionErr is an error message if we cannot get the user's current position of a ticker
	PositionErr = "could not get position"
)

// BrokerService represents the Alpaca broker functionality
type BrokerService struct {
	Client        *alpaca.Client
	ConfigService services.ConfigService
	MarketDataService services.MarketDataService
}

// Setup will initialize the Alpaca client
func (b *BrokerService) Setup() {
	// config service
	cs := b.ConfigService
	// set client id
	_ = os.Setenv(common.EnvApiKeyID, cs.Get().Alpaca.ClientID)
	// set client secret
	_ = os.Setenv(common.EnvApiSecretKey, cs.Get().Alpaca.ClientSecret)
	// set Alpaca base URL
	if cs.Get().Alpaca.AccountType == "paper" {
		alpaca.SetBaseUrl(cs.Get().Alpaca.BaseURL)
	}
	// create new client
	b.Client = alpaca.NewClient(common.Credentials())
}

