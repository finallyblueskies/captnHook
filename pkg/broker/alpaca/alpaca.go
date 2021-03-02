package alpaca

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	captnhook "github.com/bareish/captnHook/pkg"
	"github.com/shopspring/decimal"
	"os"
)

type Alpaca struct {
	Client *alpaca.Client
	ConfigService captnhook.ConfigService
}

// New will create a new alpaca client instance
func New(cs captnhook.ConfigService) *Alpaca {
	// set client id
	_ = os.Setenv(common.EnvApiKeyID, cs.Get().Alpaca.ClientID)
	// set client secret
	_ = os.Setenv(common.EnvApiSecretKey, cs.Get().Alpaca.ClientSecret)
	// set Alpaca base URL
	if cs.Get().Alpaca.AccountType == "paper" {
		alpaca.SetBaseUrl(cs.Get().Alpaca.BaseURL)
	}
	// create new client
	client := alpaca.NewClient(common.Credentials())

	return &Alpaca{
		Client: client,
		ConfigService: cs,
	}
}


// Buy buys an asset and will move funds into that ticker
func (a *Alpaca) Buy(ticker string, quantity float64) (err error) {
	order := alpaca.PlaceOrderRequest{
		AccountID:   os.Getenv(common.EnvApiKeyID),
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(quantity),
		Side:        "buy",
		Type:        "market",
		TimeInForce: "day",
	}

	// make the order from the client
	 _, err = a.Client.PlaceOrder(order)
	 if err != nil {
	 	return err
	 }

	 return nil
}

// Sell sells an asset and will move funds into that ticker
func (a *Alpaca) Sell(ticker string, quantity float64) (err error) {
	return nil
}

// GetBuyingPower returns the user buying power
func (a *Alpaca) GetBuyingPower() (float64, error) {
	// get account information
	account, err := a.Client.GetAccount()
	if err != nil {
		return 0.0, err
	}
	power, _ := account.BuyingPower.Float64()
	return power, nil
}

