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
	var price int64
	w := request
	// bind body to watchlist struct
	if err = c.Bind(&w); err != nil {
		return echo.NewHTTPError(BadRequest, BodyBindingErr)
	}
	var id string
	// determine action
	if w.Action == "Buy" {
		price, err = strconv.ParseInt(w.Price, 10, 32)
		id, err = s.BrokerService.Buy(w.Ticker, price)
		if err != nil {
			return echo.NewHTTPError(InternalServerErr, "Could not buy "+ w.Ticker, err)
		}
	} else if w.Action == "Sell" {
		id, err = s.BrokerService.Sell(w.Ticker)
		if err != nil{
			return echo.NewHTTPError(InternalServerErr, "No positions on ticker")
		} else {
			return echo.NewHTTPError(InternalServerErr, "Action not recognized")
		}
	}
	return c.JSON(http.StatusOK, id)
}

// determineAction will look at the incoming request and determine where to send the data
func determineAction(request WebHookRequest) {

}