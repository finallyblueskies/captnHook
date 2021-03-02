package captnhook

// Config is a service that is designed to provide various configuration to the rest of the application
type Config struct {
	General GeneralConfig
	Alpaca AlpacaConfig
	Binance BinanceConfig
}

// ConfigService is an interface that defines the functions needed to implement an interface
type ConfigService interface {
	// Load will do any config setup (load env vars)
	Load()
	// Get returns the config vars
	Get() Config
}

// GeneralConfig controls the general setup of the server
type GeneralConfig struct {
	BaseURL string
	AppEnv string
}

// AlpacaConfig controls Alpaca keys platform keys and base urls
type AlpacaConfig struct {
	ClientID string
	ClientSecret string
	AccountType string
	BaseURL string
}

type BinanceConfig struct {

}
