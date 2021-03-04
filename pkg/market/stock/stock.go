/*
Package market:

This is where we would create a websocket connection to the Alpaca real-time data API

 */
package stock

import (
	"github.com/bareish/captnHook/pkg/services"
)

const (
	MaxConnectionAttempts = 5 			// max retries
	TradeUpdates   = "trade_updates" 	// alpaca trade updates string
	AccountUpdates = "account_updates"  // alpaca account updates string
)

// MarketDataService ...
type MarketDataService struct  {
	ConfigService services.ConfigService
	Stream *Stream

}

// Setup ...
func (m *MarketDataService) Setup() {
	// config service
	cs := m.ConfigService
	m.Stream = NewStreamClient(cs)
}

// CurrentPrice ...
func (m *MarketDataService) CurrentPrice(msg interface{}) float32 {
	quote, _ := msg.(StreamQuote)

	return quote.BidPrice
}





