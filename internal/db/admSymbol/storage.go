package psqlAdmSymbol

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/handler/rest"
)

type admSymbolStorage struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewAdmSymbolStorage(db *sqlx.DB, rdb *redis.Client) *admSymbolStorage {
	return &admSymbolStorage{db: db, rdb: rdb}
}

func (r *admSymbolStorage) CreateSymbol(symbol rest.AdmSymbolDTO) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (abbr, full_name) VALUES ($1, $2) RETURNING id", "symbol")
	res := r.db.QueryRow(query, symbol.Abbr, symbol.FullName)
	var id int
	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *admSymbolStorage) SetPrice(symbol rest.AdmPriceDTO) error {
	_, err := r.rdb.Set(context.Background(), symbol.Abbr, symbol.Price, 0).Result()
	return err
}

func (r *admSymbolStorage) DeleteSymbol(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "symbol")
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
