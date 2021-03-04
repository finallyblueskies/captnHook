package alpaca

// GetCurrentPrice will always have an open connection to
func (b *BrokerService) GetCurrentPrice(ticker string, market string) (float32, error) {
	price, err := b.MarketDataService.CurrentPrice(ticker)
	if err != nil {
		return 0.0, err
	}
	return price, nil
}
