// Package rest acts as an interface to the Echo microframework, mocking a HTTP/2 REST API .
package rest

import (
	"github.com/bareish/captnHook/pkg/config"
	"github.com/labstack/echo/v4"
)
// Server represents a HTTP/2 REST API using the Echo microframework
type Server struct {
	Echo *echo.Echo
	ConfigService *config.Service
}

// New creates a new HTTP server with injected dependencies
func New() *Server {

	return &Server{}
}

// Start initializes the server
func (s *Server) Start() {

}

// Close gracefully shuts down server and closes all connections
func (s *Server) Close() (err error){
	return s.Echo.Close()
}