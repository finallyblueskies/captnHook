package crypto

import (
	"github.com/bareish/captnHook/pkg/services"
	"log"
)

// Todo: this is where we would create a websocket connection to the Polygon real-time data API

// DataService struct provides data from the crypto market
type DataService struct  {
	ConfigService services.ConfigService
}

// Setup initialized the stream to the crypto market
func (d *DataService) Setup() {
	log.Println("Creating crypto data service")
}

// CurrentPrice will get the current price of a ticker in the crypto market
func (d *DataService) CurrentPrice(ticker string) (float32, error) {
	return 0.0, nil
}