package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// WebHookRequest ...
type WebHookRequest struct {
	Ticker string `json:"ticker"`
	Price  string `json:"price"`
	Action string `json:"action"`
}

// Signal accepts a JSON request from TradingView and does trades based on the response data
func (s *Server) Signal(c echo.Context) (err error) {
	var request  WebHookRequest
	var price int64
	w := request
	// bind body to watchlist struct
	if err = c.Bind(&w); err != nil {
		return echo.NewHTTPError(BadRequest, BodyBindingErr)
	}
	// what tf are they sending me
	log.Println(w)

	price, err = strconv.ParseInt(w.Price, 10, 32)
	id, err = s.BrokerService.Buy(w.Ticker, price)
	return c.JSON(http.StatusOK, w)
}
