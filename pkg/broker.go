package captnhook


// BrokerService is a service that allows us to buy and sell market assets from a broker
type BrokerService interface {
	// Setup will initialize any broker service.
	Setup()
	// Buy buys an asset and will move funds into that ticker
	Buy(ticker string, amount float64) error
	// Sell sells an asset and will remove funds from that ticker
	Sell(ticker string, amount float64) error
	// GetAvailableFunds will get the amount of money from the account
	GetAvailableFunds() float64
}