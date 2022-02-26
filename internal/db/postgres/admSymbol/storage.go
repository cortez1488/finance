package psqlAdmSymbol

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/handler/rest"
)

type admSymbolStorage struct {
	db *sqlx.DB
}

func NewAdmSymbolStorage(db *sqlx.DB) *admSymbolStorage {
	return &admSymbolStorage{db: db}
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

func (r *admSymbolStorage) SetPrice(price float64) error {
	fmt.Println("set price redis logic")
	return nil
}

func (r *admSymbolStorage) DeleteSymbol(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "symbol")
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
