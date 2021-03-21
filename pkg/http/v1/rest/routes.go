package rest

import "github.com/labstack/echo/v4/middleware"

// Routes will setup all the routes
func (s *Server) Routes() {
	// logging middleware
	s.Echo.Use(middleware.Logger())
	// hello
	s.Echo.GET("/", s.Hello)
	// group API version 1 routes
	v1 := s.Echo.Group("/v1")
	// stock market route
	v1.POST("/stock", s.Stock)
	// forex market route
	// v1.POST("/stock", s.Signal)
	// crypto market route
	v1.POST("/crypto", s.Crypto)
}
