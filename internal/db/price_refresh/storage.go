package price_refresh_storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/core/price_refresh"
	"time"
)

type priceRefreshStorage struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewPriceRefreshStorage(db *sqlx.DB, rdb *redis.Client) *priceRefreshStorage {
	return &priceRefreshStorage{db: db, rdb: rdb}
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

func (r *priceRefreshStorage) RefreshPrices(data []price_refresh.Symbol) (time.Time, error) {)
	for _, symbol := range data {
		if symbol.Symbol != "" {
			r.rdb.Set(context.Background(), symbol.Symbol, symbol.Price, 0)
		}
	}

	return time.Now(), nil
}
