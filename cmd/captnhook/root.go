package captnhook

import (
	"github.com/bareish/captnHook/pkg/broker/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/rest"
	"github.com/bareish/captnHook/pkg/market"
	"github.com/bareish/captnHook/pkg/market/crypto"
	"github.com/bareish/captnHook/pkg/market/forex"
	"github.com/bareish/captnHook/pkg/market/stock"
)

// Run loads configuration variables starts our services and initializes our server
func Run() {
	// create config service
	configService := &config.Service{}
	configService.Load()

	// create stock market data service
	stockDataService := &stock.DataService{ConfigService: configService}

	// create crypto market data service
	cryptoDataService := &crypto.DataService{ConfigService: configService}

	// create forex market data service
	forexDataService := &forex.DataService{ConfigService: configService}

	// create market manager
	dataMangerService := &market.DataManagerService{
		ConfigService: configService,
		StockMarket: stockDataService,
		ForexMarket: forexDataService,
		CryptoMarket: cryptoDataService,
	}
	dataMangerService.Setup()

	// create the new broker service
	brokerService := &alpaca.BrokerService{
		ConfigService: configService,
		DataMangerService: dataMangerService,
	}
	brokerService.Setup()

	// HTTP/2 REST server
	server := rest.NewRESTServer(configService, brokerService)
	server.Start()

}
