// Package rest acts as an interface to the Echo microframework, mocking a HTTP/2 REST API .
package rest

import (
	"github.com/bareish/captnHook/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	// BadRequest is an http 400 error
	BadRequest = http.StatusBadRequest
	// BodyBindingErr is an error message when we fail to bind the request body
	BodyBindingErr = "could not bind body"
)

// Server represents a HTTP/2 REST API using the Echo microframework
type Server struct {
	Echo          *echo.Echo // Todo: this couples the application as well - refactor to use basic http server
	ConfigService services.ConfigService
	BrokerService services.BrokerService
}

// NewRESTServer creates a new HTTP server with injected services
func NewRESTServer(cs services.ConfigService, bs services.BrokerService) *Server {

	return &Server{
		Echo:          echo.New(),
		ConfigService: cs,
		BrokerService: bs,
	}
}

// Start initializes the HTTP/2 server
func (s *Server) Start() {
	// base configuration variables
	cfg := s.ConfigService.Get().General
	port := cfg.Port
	//cert := cfg.CertPath
	//key := cfg.KeyPath
	echoHandler := s.Echo
	// register routes
	s.Routes()
	// start
	echoHandler.Logger.Fatal(echoHandler.Start(":"+port))
}

// Close gracefully shuts down server and closes all connections
func (s *Server) Close() (err error) {
	return s.Echo.Close()
}
