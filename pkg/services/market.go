package services


// todo: refactor so that MarketDataService has many markets to choose from eg (stock, crytpo, forex)
// MarketDataService is a service that allows any broker to subscribe to realtime market data (stock, crypto, forex)
type DataManagerService interface {
	// Setup
	Setup()
	// GetCurrentPrice
	GetCurrentPrice(ticker string, market string) (float32,error)
}


// DataService is a service that connects to the MarketDataService for switching between types of markets
type DataService interface {
	// Setup will initialize a live market data service
	Setup()
	// CurrentPrice will return the live current price of
	CurrentPrice(ticker string) (float32, error)
}