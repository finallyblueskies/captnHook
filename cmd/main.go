package main

import (
	"github.com/bareish/captnHook/pkg/broker/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/rest"
)

func main() {
	// setup config service
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
