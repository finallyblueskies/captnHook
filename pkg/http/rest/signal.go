package rest

import (
	"github.com/labstack/echo/v4"
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
	w := WebHookRequest{}
	// bind body to watchlist struct
	if err = c.Bind(&w); err != nil {
		return echo.NewHTTPError(BadRequest, BodyBindingErr)
	}

	return c.JSON(http.StatusOK, w)
}
