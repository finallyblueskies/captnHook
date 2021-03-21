package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Crypto accepts a JSON request from TradingView and does trades in the crypto market based on the response data
func (s *Server) Crypto(c echo.Context) (err error) {
	var request WebHookRequest
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
		err = s.ManageBrokerService.GetCryptoService().BuyAll(w.Ticker, price)
		if err != nil {
			return echo.NewHTTPError(ErrInternalServer, err)
		}
	}
	if w.Action == "Sell" {
		err = s.ManageBrokerService.GetCryptoService().SellAll(w.Ticker)
		if err != nil{
			return echo.NewHTTPError(ErrInternalServer, "No positions on ticker")
		}

	}
	return c.JSON(http.StatusOK, id)
}

