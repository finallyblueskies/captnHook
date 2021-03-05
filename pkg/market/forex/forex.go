package forex

import (
	"github.com/bareish/captnHook/pkg/services"
	"log"
)

// Todo: this is where we would create a websocket connection to the Polygon real-time data API

// DataService struct provides data from the forex market
type DataService struct  {
	ConfigService services.ConfigService
}

// Setup initializes stream to forex market
func (d *DataService) Setup() {
	log.Println("forex data service")
}

// CurrentPrice will give the price of a ticker from the forex market
func (d *DataService) CurrentPrice(ticker string) (float32, error) {
	return 0.0, nil
}