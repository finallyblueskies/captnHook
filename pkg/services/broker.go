package services

// BrokerService is a service that allows us to buy and sell market assets from a broker
type BrokerService interface {
	// Setup will initialize any broker service.
	Setup()
	// BuyAll ...
	BuyAll(ticker string, currentPrice float64) error
	// SellAll sells an asset and will remove funds from that ticker
	SellAll(ticker string) error
	// Buy buys an asset and will move funds into that ticker
	Buy(ticker string, amount float64)  error
	// Sell ..
	Sell(ticker string, shares float64) error
	// GetBuyingPower will get the amount of money from the account
	GetBuyingPower() (float64, error)
}
