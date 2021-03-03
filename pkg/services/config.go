package services

// Config is a service that is designed to provide various configuration to the rest of the application
type Config struct {
	General GeneralConfig
	Alpaca  AlpacaConfig
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
	Port     string
	BaseURL  string
	AppEnv   string
	CertPath string
	KeyPath  string
}

// AlpacaConfig controls Alpaca configuration variables
type AlpacaConfig struct {
	ClientID     string
	ClientSecret string
	AccountType  string
	BaseURL      string
}

// BinanceConfig controls Binance configuration variables
type BinanceConfig struct {
	// TODO: implement
}

// CoinbaseConfig controls Coinbase configuration variables
type CoinbaseConfig struct {
	// TODO: implement
}
