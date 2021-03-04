package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// WebHookRequest ...
type WebHookRequest struct {
	Ticker string `json:"ticker"`
	Price  string `json:"price"`
	Action string `json:"action"`
}

// Signal accepts a JSON request from TradingView and does trades based on the response data
func (s *Server) Signal(c echo.Context) (err error) {
	var empty interface{}
	w := empty
	// bind body to watchlist struct
	if err = c.Bind(&w); err != nil {
		return echo.NewHTTPError(BadRequest, BodyBindingErr)
	}
	// what tf are they sending me
	log.Println(w)
	return c.JSON(http.StatusOK, w)
}
