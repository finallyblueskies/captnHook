package captnhook

import (
	"github.com/bareish/captnHook/pkg/broker/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/rest"
)

// Run loads configuration variables starts our services and initialize our server
func Run() {
	// create config service
	configService := &config.Service{}
	configService.Load()

	// create the new broker service
	brokerService := &alpaca.BrokerService{
		ConfigService: configService,
	}
	brokerService.Setup()

	// HTTP/2 REST server
	server := rest.New(configService, brokerService)
	server.Start()
}
