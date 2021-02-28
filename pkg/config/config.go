package config

import (
	"log"

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

