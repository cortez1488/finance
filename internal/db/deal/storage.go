package dealStorage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/core/deal"
	"strconv"
)

type dealStorage struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewDealStorage(db *sqlx.DB, rdb *redis.Client) *dealStorage {
	return &dealStorage{db: db, rdb: rdb}
}

func (r *dealStorage) GetShareInfo(id int) (deal.Symbol, error) {
	var output deal.Symbol
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", "symbol")
	row := r.db.QueryRow(query, id)
	err := row.Scan(&output.ID, &output.Abbr, &output.FullName)
	if err != nil {
		return deal.Symbol{}, err
	}

	err = setPrice(r.rdb, &output)
	if err != nil {
		return deal.Symbol{}, err
	}

	return output, nil
}

func (r *dealStorage) GetShareListInfo() ([]deal.Symbol, error) {
	var output []deal.Symbol
	query := fmt.Sprintf("SELECT * FROM %s", "symbol")
	err := r.db.Select(&output, query)
	if err != nil {
		return []deal.Symbol{}, err
	}

	for index, _ := range output {
		err := setPrice(r.rdb, &output[index])
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func setPrice(rdb *redis.Client, symb *deal.Symbol) error {
	price, err := getSymbolPrice(rdb, *symb)
	if err != nil {
		return err
	}

	symb.Price = price
	return nil
}

func getSymbolPrice(rdb *redis.Client, symb deal.Symbol) (float64, error) {
	priceStr, err := rdb.Get(context.Background(), symb.Abbr).Result()
	if err != nil {
		return 0, err
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}
