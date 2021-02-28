package rest

import (
	"github.com/bareish/captnHook/pkg/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	ConfigService *config.Service
}

func (s *Server) Start() {
}