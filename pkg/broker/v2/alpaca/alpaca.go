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
	ClientID string
	ClientSecret string
}

// Setup will initialize the Alpaca client
func (b *BrokerService) Setup() {
	// alpaca config variables
	alpacaConfig := b.ConfigService.Get().Alpaca
	clientID := alpacaConfig.ClientID
	secretID := alpacaConfig.ClientSecret
	baseURL := alpacaConfig.BaseURL
	accountType := alpacaConfig.AccountType
	// set client id
	_ = os.Setenv(common.EnvApiKeyID, clientID)
	// set client secret
	_ = os.Setenv(common.EnvApiSecretKey, secretID)
	// set Alpaca base URL
	if accountType == "paper" {
		alpaca.SetBaseUrl(baseURL)
	}
	// create new client
	b.Client = alpaca.NewClient(common.Credentials())
}

