package alpaca

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	captnhook "github.com/bareish/captnHook/pkg"
	"github.com/shopspring/decimal"
	"os"
)

type BrokerService struct {
	ConfigService captnhook.ConfigService
}

// Setup initializes configuration variables
func (b *BrokerService) Setup() {
	c := b.ConfigService.Get()
	_ = os.Setenv(common.EnvApiKeyID, c.Alpaca.ClientID)
	_ = os.Setenv(common.EnvApiSecretKey, c.Alpaca.ClientSecret)
}


// Buy buys an assets and will move funds into that ticker
func (b *BrokerService) Buy(ticker string, quantity float64) (err error) {
	order := alpaca.PlaceOrderRequest{
		AccountID:   os.Getenv(common.EnvApiKeyID),
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(quantity),
		Side:        "buy",
		Type:        "market",
		TimeInForce: "day",
	}
	// make the order
	 _, err = alpaca.PlaceOrder(order)
	 if err != nil {
	 	return err
	 }

	 return nil
}