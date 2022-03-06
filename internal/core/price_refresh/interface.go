package price_refresh

import "time"

type PricesRefreshStorage interface {
	GetCurrentSymbols() ([]string, error)
	RefreshPrices([]Symbol) (time.Time, error)
}
