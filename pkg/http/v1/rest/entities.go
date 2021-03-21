package rest


// WebHookRequest ...
type WebHookRequest struct {
	Ticker string `json:"ticker"`
	Price  string `json:"price"`
	Action string `json:"action"`
}
