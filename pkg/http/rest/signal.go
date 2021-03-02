package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// WebHookRequest ...
type WebHookRequest struct {

}


// Signal accepts a JSON request from TradingView and does trades based on the response data
func (s *Server) Signal(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "hello there!"})
}
