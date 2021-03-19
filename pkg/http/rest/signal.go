package rest

import (
	"github.com/labstack/echo/v4"
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
	var price float64
	w := request
	// bind body to watchlist struct
	if err = c.Bind(&w); err != nil {
		return echo.NewHTTPError(ErrBadRequest, ErrBodyBinding)
	}
	var id string
	// determine action
	if w.Action == "Buy" {
		price, err = strconv.ParseFloat(w.Price, 64)
		if err != nil {
			return echo.NewHTTPError(ErrInternalServer, ErrStringConversion)
		}
		err = s.BrokerService.BuyAll(w.Ticker, price)
		if err != nil {
			return echo.NewHTTPError(ErrInternalServer, "Could not buy "+ w.Ticker, err)
		}
	}
	if w.Action == "Sell" {
		err = s.BrokerService.SellAll(w.Ticker)
		if err != nil{
			return echo.NewHTTPError(ErrInternalServer, "No positions on ticker")
		}

	}
	return c.JSON(http.StatusOK, id)
}

