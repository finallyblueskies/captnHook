package captnhook

import (
	"github.com/bareish/captnHook/pkg/broker"
	"github.com/bareish/captnHook/pkg/broker/v2/alpaca"
	"github.com/bareish/captnHook/pkg/broker/v2/coinbase"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/v1/rest"
)

// Execute loads configuration variables starts our services and initializes our server
func Execute() {
	// create config service
	configService := &config.Service{}
	configService.Load()

	// create the new stock broker service
	stockBrokerService := &alpaca.BrokerService{
		ConfigService: configService,
	}

	// create the new crypto broker service
	cryptoBrokerService := &coinbase.BrokerService{
		ConfigService: configService,
	}

	// create the new forex broker service
	// forexBrokerService

	// create broker management service
	brokerManager := &broker.ManageBrokerService{
		StockBroker: stockBrokerService,
		CryptoBroker: cryptoBrokerService,
	}
	brokerManager.Setup()

	// HTTP/2 REST server
	server := rest.NewRESTServer(configService, brokerManager)
	server.Start()

}
