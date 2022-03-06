package price_refresh

type PricesRefreshStorage interface {
	GetCurrentSymbols() ([]string, error)
	RefreshPrices(*[]Symbol) error
}
