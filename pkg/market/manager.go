package market

import (
	"github.com/bareish/captnHook/pkg/services"
)

// DataManagerService
type DataManagerService struct {
	ConfigService services.ConfigService
	CryptoMarket services.DataService
	StockMarket services.DataService
	ForexMarket services.DataService
}



// Setup will just configure our services
func (d *DataManagerService) Setup() {
	// setup our services
	// todo: add switch
	d.CryptoMarket.Setup()
	d.ForexMarket.Setup()
	d.StockMarket.Setup()
}

// GetCurrentPrice will look the current price of a stock depending on the market
func (d *DataManagerService) GetCurrentPrice(ticker string, market string) (float32,error){
	// todo: do some cool shit in here
	return 0.0, nil
}
