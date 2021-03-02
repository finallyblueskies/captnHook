package alpaca

import (
	"errors"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/bareish/captnHook/pkg/services"
	"github.com/shopspring/decimal"
	"os"
)

// BrokerService
type BrokerService struct {
	Client        *alpaca.Client
	ConfigService services.ConfigService
}

// Get will initialize the alpaca client
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


// Buy buys an asset and will move funds into that ticker
func (b *BrokerService)Buy(ticker string, quantity float64) (err error) {
	// check buying power
	buyingPower, err := b.GetBuyingPower()
	if err != nil {
		return err
	}
	if buyingPower == 0.0 {
		return errors.New("no funds in account")
	}
	order := alpaca.PlaceOrderRequest{
		AccountID:   common.EnvApiKeyID,
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(quantity),
		Side:        "buy",
		Type:        "market",
		TimeInForce: "day",
	}
	// make the order from the client
	 _, err = b.Client.PlaceOrder(order)
	 if err != nil {
	 	return err
	 }

	 return nil
}

// Sell sells an asset and will move funds into that ticker
func (b *BrokerService)Sell(ticker string, quantity float64) (err error) {
	return nil
}

// GetBuyingPower returns the user buying power
func (b *BrokerService)GetBuyingPower() (float64, error) {
	// get account information
	account, err := b.Client.GetAccount()
	if err != nil {
		return 0.0, err
	}
	power, _ := account.BuyingPower.Float64()
	return power, nil
}

