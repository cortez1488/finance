package price_refresh_storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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
	query := fmt.Sprintf("SELECT abbr FROM %s", viper.GetString("db.postgres.tableNames.symbol"))
	var currentSymbols []string

	err := r.db.Select(&currentSymbols, query)
	if err != nil {
		return nil, err
	}

	return currentSymbols, nil
}

func (r *priceRefreshStorage) RefreshPrices(data []price_refresh.Symbol) (time.Time, error) {
	for _, symbol := range data {
		if symbol.Symbol != "" {
			res := r.rdb.Set(context.Background(), symbol.Symbol, symbol.Price, 0)
			if res.Err() != nil {
				return time.Now(), res.Err()
			}
		}
	}

	return time.Now(), nil
}
