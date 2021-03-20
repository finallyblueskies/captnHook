package captnhook

import (
	"github.com/bareish/captnHook/pkg/broker/v2/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/v1/rest"
)

// Execute loads configuration variables starts our services and initializes our server
func Execute() {
	// create config service
	configService := &config.Service{}
	configService.Load()

	// todo: create broker management service

	// create the new broker service
	brokerService := &alpaca.BrokerService{
		ConfigService: configService,
	}
	brokerService.Setup()

	// HTTP/2 REST server
	server := rest.NewRESTServer(configService, brokerService)
	server.Start()

}
