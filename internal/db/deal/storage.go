package dealStorage

import (
	"context"
	"errors"
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

type actType string

const (
	typeSell actType = "SELL"
	typeBuy  actType = "BUY"
)

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

//--------------------------------------------|
//---------- BUY / SELL BUSINESS -------------|
//--------------------------------------------|

func (r *dealStorage) BuyShares(shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date string, dType actType) error {

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	dealID, err := createDeal(tx, dType, shareID, portfolioID, userID, quantity, symbolPrice, amount, date)
	if err != nil {
		tx.Rollback()
		return err
	}

	intDealID := int(dealID)
	err = createActiveShare(tx, intDealID, portfolioID, shareID, quantity, symbolPrice)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = changePortfolioAccount(tx, dType, portfolioID, amount)
	return nil
}

func createDeal(tx *sqlx.Tx, dType actType, shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date string) (int64, error) {
	dealCreateQuery := fmt.Sprintf("INSERT INTO %s (type, symbol_id, symbol_price, number, amount, date,"+
		" portfolio_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "deal")

	res, err := tx.Exec(dealCreateQuery, dType, shareID, portfolioID, userID, quantity, symbolPrice, amount, date)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func createActiveShare(tx *sqlx.Tx, dealID, portfolioID, shareID, quantity int, symbolPrice float64) error {
	activeShareQuery := fmt.Sprintf("INSERT INTO %s (price, number, portfolio_id, symbol_id, deal_id) "+
		"VALUES ($1, $2, $3, $4, $5)", "active_share")

	_, err := tx.Exec(activeShareQuery, symbolPrice, quantity, portfolioID, shareID, dealID)
	if err != nil {
		return err
	}

	return nil
}

func changePortfolioAccount(tx *sqlx.Tx, dType actType, portfolioID int, amount float64) error {
	var actChar string
	if dType == typeSell {
		actChar = "+"
	} else if dType == typeBuy {
		actChar = "-"
	} else {
		return errors.New("incorrect type of action")
	}

	portfolioQuery := fmt.Sprintf("UPDATE %s SET account = account %s $1 WHERE id = $2", "portfolio", actChar)
	_, err := tx.Exec(portfolioQuery, amount, portfolioID)
	if err != nil {
		return err
	}

	return nil
}
