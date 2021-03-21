package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
)

// BuyAll ...
func (b *BrokerService) BuyAll(ticker string, currentPrice float64) error {
	// build order request
	order, err := b.Client.NewCreateOrderService().Symbol(ticker).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity("5").
		Price("0.0030000").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(order)
	return nil
}

// SellAll ...
func (b *BrokerService) SellAll(ticker string) error {
	return nil
}

// Sell ..
func (b *BrokerService) Sell(ticker string, shares float64) error {
	return nil
}
// Buy ...
func (b *BrokerService) Buy(ticker string, shares float64) error {
	return nil
}

// GetBuyingPower ...
func (b *BrokerService) GetBuyingPower() (float64, error){
	return 0.0, nil
}
