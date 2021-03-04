package services

// MarketData is a service that allows any broker to subscribe to realtime market data (stock, crypto, forex)
type MarketDataService interface {
	// Setup will initialize any live market data service.
	Setup()
	// CurrentPrice will return the live current price of
	CurrentPrice(ticker string)
}
