package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello is a simple 200 status check that can be used to see if the server is running.
func (s *Server) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "hello there!"})
}
