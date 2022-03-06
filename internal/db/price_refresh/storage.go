package price_refresh_storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/core/price_refresh"
)

type priceRefreshStorage struct {
	db *sqlx.DB
}

func NewPriceRefreshStorage(db *sqlx.DB) *priceRefreshStorage {
	return &priceRefreshStorage{db: db}
}

func (r *priceRefreshStorage) GetCurrentSymbols() ([]string, error) {
	query := fmt.Sprintf("SELECT abbr FROM symbol")
	var currentSymbols []string

	err := r.db.Select(&currentSymbols, query)
	if err != nil {
		return nil, err
	}

	return currentSymbols, nil
}

func (r *priceRefreshStorage) RefreshPrices(data *[]price_refresh.Symbol) error {
	fmt.Println(*data)
	return nil
}
