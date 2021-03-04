/*
Package market:

This is where we would create a websocket connection to the Alpaca real-time data API

 */
package stock

import (
	"github.com/bareish/captnHook/pkg/services"
)



// MarketDataService ...
type MarketDataService struct  {
	ConfigService services.ConfigService
	Stream *Stream
	PriceChan chan StreamQuote
}

// Setup ...
func (m *MarketDataService) Setup() {
	// config service
	cs := m.ConfigService
	// create stream client
	m.Stream = NewStreamClient(cs)
	// start websocket connection and stream data
	go m.Stream.Start()
}

// CurrentPrice ...
func (m *MarketDataService) CurrentPrice(ticker string) (float32, error) {
	// subscribe to ticker
	// m.Stream.StreamQuote(ticker,)
	return 0.0, nil
}




