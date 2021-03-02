package rest

import "github.com/labstack/echo/v4/middleware"

// Routes will setup all the routes
func (s *Server) Routes() {
	// logging middleware
	// todo make logger middleware a service
	s.Echo.Use(middleware.Logger())
	// hello
	s.Echo.GET("/", s.Hello)
	// signal
	s.Echo.POST("/signal", s.Signal)
}