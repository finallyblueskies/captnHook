

package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
)

// BuyAll ...
func (b *BrokerService) BuyAll(ticker string, currentPrice float64) error {
	// // build order request
	fmt.Println("Buy all")

	res, err := b.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
			fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	usdt := ""
	for _, v := range res.Balances {
    if v.Asset == "USDT" {
      usdt = v.Free
    }
	}
	
	fmt.Println("Selling", usdt, "usdt")
	order, err := b.Client.NewCreateOrderService().
		Symbol("BTCUSD").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).
		QuoteOrderQty(usdt).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return err
	}
	// b.GetBuyingPower()
	// fmt.Println(buyingPower, err)
	fmt.Printf("%+v", order)
	return nil
}

// SellAll ...
func (b *BrokerService) SellAll(ticker string) error {
	fmt.Println("Sell all")

	res, err := b.Client.NewGetAccountService().Do(context.Background())
	fmt.Printf("%+v", res)
	if err != nil {
			fmt.Println(err)
	}
	btc := ""
	for _, v := range res.Balances {
		if v.Asset == "BTC" {
      btc = v.Free
    }
	}
	fmt.Println("Selling", btc, "btc")
	order, err := b.Client.NewCreateOrderService().
		Symbol("BTCUSD").
		Side(binance.SideTypeSell).
		Type(binance.OrderTypeMarket).
		Quantity(btc).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return err
	}
	// b.GetBuyingPower()
	// fmt.Println(buyingPower, err)
	fmt.Printf("%+v", order)
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
	res, err := b.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
			fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	return 0.0, err
}
