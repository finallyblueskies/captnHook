/*
Package market:

This is where we would create a websocket connection to the Alpaca real-time data API

 */
package stock

import (
	"github.com/bareish/captnHook/pkg/services"
	cmap "github.com/orcaman/concurrent-map"
	"log"
)

// DataService struct provides data from the stock market
type DataService struct {

}
// MarketDataService ...
type MarketDataService struct  {
	ConfigService services.ConfigService
	Stream *Stream
	PriceChan chan StreamQuote
	PriceMap cmap.ConcurrentMap
	SeenTickers map[string]bool
	PriceVal float32
}

// Setup ...
func (m *MarketDataService) Setup() {
	m.PriceChan = make(chan StreamQuote, 128)
	// config service
	cs := m.ConfigService
	// create stream client
	m.Stream = NewStreamClient(cs)
	// start websocket connection and stream data
	go m.Stream.Start()
}

// CurrentPrice ...
func (m *MarketDataService) CurrentPrice(ticker string) (float32, error) {


	/* concurrent map implementation
	if ticker not in map start streaming data to map on its own go routine
	price, ok := m.PriceMap.Get(ticker)
	if !ok {
		// start stream
		go m.currentPriceStream(ticker)
		return m.PriceMap.Get(ticker)
	}
	// deconstruct interface
	currentPrice := price.(atomic.Value)
	*/

	//if _, ok := m.SeenTickers[ticker]; ok {
	//	return m.PriceVal, nil
	//}
	// start stream
	go m.currentPriceStream(ticker)
	// add ticker to seen
	//m.SeenTickers[ticker] = true

	return m.PriceVal, nil

}

func (m *MarketDataService) currentPriceStream(ticker string) error {
	// subscribe to ticker
	err := m.Stream.StreamQuote(ticker, m.PriceChan)
	if err != nil {
		return err
	}

	for stock := range m.PriceChan {
		log.Println(stock.AskPrice)
	}
	return nil
}




