package alpaca

import (
	"errors"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/shopspring/decimal"
	"log"
)

// Buy buys an asset and will move funds into that ticker
// todo we can allow buy a certain amount but right now we going all in
func (b *BrokerService) Buy(ticker string, currentPrice int64) (string, error) {
	// check buying power and convert that into share amount
	buyingPower, err := b.GetBuyingPower()
	log.Println(buyingPower)
	// calculate total amount of shares we can afford
	shares := buyingPower / float64(currentPrice)
	log.Println(shares)
	if err != nil {
		return "", err
	}
	if buyingPower == 0.0 {
		return "", errors.New(NoFundsErr)
	}
	order := alpaca.PlaceOrderRequest{
		AccountID:   common.EnvApiKeyID,
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(shares), // we going all in bois
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
func (b *BrokerService) Sell(ticker string) (orderID string, err error) {
	// lets get total amount of shares
	position, err := b.Client.GetPosition(ticker)
	if position == nil {
		return "", errors.New("We have no positions on: " + ticker)
	}
	log.Println(position)
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
	// place the order form the client API
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
