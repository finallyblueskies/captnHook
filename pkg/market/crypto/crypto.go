package crypto

import "log"

// Todo: this is where we would create a websocket connection to the Polygon real-time data API

// DataService struct provides data from the crypto market
type DataService struct  {

}

// Setup initialized the stream to the crypto market
func (d *DataService) Setup() {
	log.Println("Creating crypto data service")
}