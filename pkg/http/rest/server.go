// Package rest acts as an interface to the Echo microframework, mocking a HTTP/2 REST API .
package rest

import (
	"github.com/bareish/captnHook/pkg/services"
	"github.com/labstack/echo/v4"
)

// Server represents a HTTP/2 REST API using the Echo microframework
type Server struct {
	Echo          *echo.Echo
	ConfigService services.ConfigService
	BrokerService services.BrokerService
}

// New creates a new HTTP server with injected services
func New(cs services.ConfigService, bs services.BrokerService) *Server {

	return &Server{
		Echo: echo.New(),
		ConfigService: cs,
		BrokerService: bs,
	}
}

// Start initializes the HTTP/2 server
func (s *Server) Start() {
	// base configuration variables
	cfg := s.ConfigService.Get().General
	port := cfg.Port
	cert := cfg.CertPath
	key := cfg.KeyPath
	echoHandler := s.Echo
	// register routes
	s.Routes()
	// start
	echoHandler.Logger.Fatal(echoHandler.StartTLS(":"+port, cert, key))
}

// Close gracefully shuts down server and closes all connections
func (s *Server) Close() (err error){
	return s.Echo.Close()
}