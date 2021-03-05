package forex

import "log"

// Todo: this is where we would create a websocket connection to the Polygon real-time data API

// DataService struct provides data from the forex market
type DataService struct  {

}

// Setup initializes stream to forex market
func (d *DataService) Setup() {
	log.Println("forex data service")
}

