package alpaca

import (
	"errors"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/bareish/captnHook/pkg/services"
	"github.com/shopspring/decimal"
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

// Buy buys an asset and will move funds into that ticker
// todo we can allow buy a certain amount but right now we going all in
func (b *BrokerService) Buy(ticker string, quantity float64) (string, error) {
	// check buying power
	buyingPower, err := b.GetBuyingPower()
	if err != nil {
		return "", err
	}
	if buyingPower == 0.0 {
		return "", errors.New(NoFundsErr)
	}
	order := alpaca.PlaceOrderRequest{
		AccountID:   common.EnvApiKeyID,
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(buyingPower), // we going all in bois
		Side:        "buy",
		Type:        "market",
		TimeInForce: "day",
	}
	// make the order from the client
	placedOrder, err := b.Client.PlaceOrder(order)
	if err != nil {
		return "", err
	}

	return placedOrder.ClientOrderID, nil
}

// Sell sells an asset and will move funds into that ticker
func (b *BrokerService) Sell(ticker string, quantity float64) (orderID string, err error) {
	// lets get the position
	position, err := b.Client.GetPosition(ticker)
	if err != nil {
		return "", errors.New(PositionErr)
	}
	order := alpaca.PlaceOrderRequest{
		AccountID:   common.EnvApiKeyID,
		AssetKey:    &ticker,
		Qty:         position.Qty, // sell everything
		Side:        "sell",
		Type:        "market",
		TimeInForce: "day",
	}
	// make the order from the client
	placedOrder, err := b.Client.PlaceOrder(order)
	if err != nil {
		return "", err
	}

	return placedOrder.ClientOrderID, nil
}

// GetBuyingPower returns the user buying power
func (b *BrokerService) GetBuyingPower() (float64, error) {
	// get account information
	account, err := b.Client.GetAccount()
	if err != nil {
		return 0.0, err
	}
	power, _ := account.BuyingPower.Float64()
	return power, nil
}
