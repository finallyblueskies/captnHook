// config package
package config

import (
	captnhook "github.com/bareish/captnHook/pkg"
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
func (c *Service) Get() captnhook.Config {
	return captnhook.Config{
		General: getGeneralConfig(),
		Alpaca: getAlpacaConfig(),
	}
}

// return the server general configuration
func getGeneralConfig() captnhook.GeneralConfig {
	// default to dev settings
	config := captnhook.GeneralConfig{
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
func getAlpacaConfig() captnhook.AlpacaConfig {
	return captnhook.AlpacaConfig{
		// client id
		ClientID: os.Getenv("ALPACA_CLIENT_ID"),
		// client secret
		ClientSecret: os.Getenv("ALPACA_CLIENT_SECRET"),
		// account type
		AccountType: os.Getenv("ALPACA_ACCOUNT_TYPE"),
		// base
		BaseURL: os.Getenv("ALPACA_BASE_URL"),
	}
}




