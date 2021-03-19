package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
)

// BuyAll ...
func (b *BrokerService) BuyAll(ticker string) {
	// build order request
	order, err := b.Client.NewCreateOrderService().Symbol(ticker).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity("5").
		Price("0.0030000").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

// SellAll ...
func (b *BrokerService) SellAll(ticker string) {

}

// Sell ..
func (b *BrokerService) Sell(ticker string, shares float64) {

}
// Buy ...
func (b *BrokerService) Buy(ticker string, shares float64) {

}

// GetBuyingPower ...
func (b *BrokerService) GetBuyingPower() {

}
