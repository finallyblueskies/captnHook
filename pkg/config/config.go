// config package
package config

import (
	captnhook2 "github.com/bareish/captnHook/pkg/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Service is the environment implementation
type Service struct{}

// Load will load the configuration variables
func (c *Service) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("did not load env vars from .env")
	}
}

// Get will return the default config
func (c *Service) Get() captnhook2.Config {
	return captnhook2.Config{
		General: getGeneralConfig(),
		Alpaca: getAlpacaConfig(),
	}
}

// return the server general configuration
func getGeneralConfig() captnhook2.GeneralConfig {
	// default to dev settings
	config := captnhook2.GeneralConfig{
		Port: os.Getenv("PORT"),
		CertPath: os.Getenv("CERT_PATH"),
		KeyPath: os.Getenv("KEY_PATH"),
		AppEnv: os.Getenv("MODE"),
		BaseURL: os.Getenv("BASE_URL"),
	}
	// change urls if we are in production
	if config.AppEnv == "PROD" {
		config.BaseURL = "https://0.0.0.0"
	}

	return config
}

// return the Alpaca configuration
func getAlpacaConfig() captnhook2.AlpacaConfig {
	// setup config values
	config := captnhook2.AlpacaConfig{
		// client id
		ClientID: os.Getenv("ALPACA_CLIENT_ID"),
		// client secret
		ClientSecret: os.Getenv("ALPACA_CLIENT_SECRET"),
		// account type
		AccountType: os.Getenv("ALPACA_ACCOUNT_TYPE"),
		// base - we default to paper url lmao
		BaseURL: "https://paper-api.alpaca.markets",
	}
	// check if we are using the live url
	if config.AccountType == "live" {
		config.BaseURL = "https://api.alpaca.markets"
	}

	return config
}




