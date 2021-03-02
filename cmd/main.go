package main

import (
	"github.com/bareish/captnHook/pkg/broker/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"log"
)

func main() {
	// setup config service
	configService := &config.Service{}
	configService.Load()
	// create the new broker service
	client := alpaca.New(configService)
	 b, err := client.GetBuyingPower()
	 if err != nil {
	 	panic(err)
	 }
	 log.Println(b)
}
