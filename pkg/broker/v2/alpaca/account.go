package alpaca

import (
	"errors"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/shopspring/decimal"
	"math"
)

// BuyAll will move all available funds into a ticker
func (b *BrokerService) BuyAll(ticker string, currentPrice float64)  error {
	var shares float64
	// determine if we can buy fractions of asset
	asset, err := b.Client.GetAsset(ticker)
	if err != nil {
		return err
	}
	// initialize a stop loss that is equivalent to the initial price
	// limitPrice  := decimal.NewFromFloat(currentPrice)
	// account buying power
	buyingPower, err := b.GetBuyingPower()
	if err != nil {
		return err
	}
	// check if we have have any available funds so we can exit early
	if buyingPower == 0.0 {
		return errors.New(NoFundsErr)
	}
	// if we can buy fractions of asset
	if asset.Shortable {
		// calculate total amount of shares we can afford
		shares = buyingPower / currentPrice
	} else {
		// round to the nearest who number
		shares = math.Round(buyingPower / currentPrice) - 1
	}
	// build order request
	order := alpaca.PlaceOrderRequest{
		AccountID:   b.ConfigService.Get().Alpaca.ClientID,
		AssetKey:    &ticker,
		Qty:         decimal.NewFromFloat(shares),
		Side:        "buy",
		Type:        "market",
		TimeInForce: "day",
		//LimitPrice: &limitPrice,
	}
	// place the order from Alpaca client
	_, err = b.Client.PlaceOrder(order)
	if err != nil {
		return err
	}

	return  nil
}

// SellAll sells all available shares of a position
func (b *BrokerService) SellAll(ticker string) error {
	// what position do we have on specified ticker
	position, err := b.Client.GetPosition(ticker)
	// if we have no position on current ticker then we can exit early
	if position == nil {
		return errors.New("We have no positions on: " + ticker)
	}
	// position err
	if err != nil {
		return errors.New(PositionErr)
	}
	// build order request
	order := alpaca.PlaceOrderRequest{
		AccountID:   common.EnvApiKeyID,
		AssetKey:    &ticker,
		Qty:         position.Qty,
		Side:        "sell",
		Type:        "market",
		TimeInForce: "day",
	}
	// place the order form the client API
	_, err = b.Client.PlaceOrder(order)
	if err != nil {
		return err
	}

	return  nil
}

// Sell ...
func (b *BrokerService)Sell(ticker string, amount float64) error {
	return nil
}

// Buy ...
func (b *BrokerService)Buy(ticker string, amount float64) error {
	return nil
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
