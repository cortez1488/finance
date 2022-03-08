package dealStorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"myFinanceTask/internal/core/deal"
	"strconv"
	"time"
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
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", viper.GetString("db.postgres.tableNames.symbol"))
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
	query := fmt.Sprintf("SELECT * FROM %s", viper.GetString("db.postgres.tableNames.symbol"))
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

func (r *dealStorage) SellShares(activeShareID, shareID, portfolioID, userID, quantity int, symbolPrice, amount float64,
	date time.Time, dType deal.ActType) error {

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = createDeal(tx, dType, shareID, portfolioID, userID, quantity, symbolPrice, amount, date)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = downActiveShare(tx, activeShareID, quantity)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = changePortfolioAccount(tx, dType, portfolioID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *dealStorage) BuyShares(shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date time.Time,
	dType deal.ActType) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	dealID, err := createDeal(tx, dType, shareID, portfolioID, userID, quantity, symbolPrice, amount, date)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createActiveShare(tx, dealID, portfolioID, shareID, quantity, symbolPrice)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = changePortfolioAccount(tx, dType, portfolioID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func createDeal(tx *sqlx.Tx, dType deal.ActType, shareID, portfolioID, userID, quantity int, symbolPrice, amount float64,
	date time.Time) (int, error) {

	dealCreateQuery := fmt.Sprintf("INSERT INTO %s (type, symbol_id, symbol_price, number, amount, date,"+
		" portfolio_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		viper.GetString("db.postgres.tableNames.deal"))

	row := tx.QueryRow(dealCreateQuery, dType, shareID, symbolPrice, quantity, amount, date, portfolioID, userID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}
	return id, nil
}

func createActiveShare(tx *sqlx.Tx, dealID, portfolioID, shareID, quantity int, symbolPrice float64) error {
	activeShareQuery := fmt.Sprintf("INSERT INTO %s (price, number, portfolio_id, symbol_id, deal_id) "+
		"VALUES ($1, $2, $3, $4, $5)", viper.GetString("db.postgres.tableNames.activeShare"))

	_, err := tx.Exec(activeShareQuery, symbolPrice, quantity, portfolioID, shareID, dealID)
	if err != nil {
		return err
	}

	return nil
}

func downActiveShare(tx *sqlx.Tx, activeShareID, quantity int) error {
	downActiveShareQuery := fmt.Sprintf("UPDATE %s SET number = number - $1 WHERE id = $2",
		viper.GetString("db.postgres.tableNames.activeShare"))
	_, err := tx.Exec(downActiveShareQuery, quantity, activeShareID)
	if err != nil {
		return err
	}
	return nil
}

func changePortfolioAccount(tx *sqlx.Tx, dType deal.ActType, portfolioID int, amount float64) error {
	var actChar string
	if dType == deal.TypeSell {
		actChar = "+"
	} else if dType == deal.TypeBuy {
		actChar = "-"
	} else {
		return errors.New("incorrect type of action")
	}

	portfolioQuery := fmt.Sprintf("UPDATE %s SET account = account %s $1 WHERE id = $2",
		viper.GetString("db.postgres.tableNames.portfolio"), actChar)
	_, err := tx.Exec(portfolioQuery, amount, portfolioID)
	if err != nil {
		return err
	}

	return nil
}

func (r *dealStorage) GetShareInfoOfActiveShareID(activeShareID int) (deal.Symbol, error) {
	queryFromActiveShare := fmt.Sprintf("SELECT symbol_id FROM %s WHERE id = $1",
		viper.GetString("db.postgres.tableNames.activeShare"))
	querySymbol := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", viper.GetString("db.postgres.tableNames.symbol"))

	var symbolID int
	row := r.db.QueryRow(queryFromActiveShare, activeShareID)
	err := row.Scan(&symbolID)
	if err != nil {
		return deal.Symbol{}, err
	}

	var symbol deal.Symbol
	row = r.db.QueryRow(querySymbol, symbolID)
	err = row.Scan(&symbol.ID, &symbol.Abbr, &symbol.FullName)
	if err != nil {
		return deal.Symbol{}, err
	}

	price, err := getSymbolPrice(r.rdb, symbol)
	if err != nil {
		return deal.Symbol{}, err
	}
	symbol.Price = price

	return symbol, nil
}

func (r *dealStorage) IsPortfoliosOwner(userID, portfolioID int) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1 AND id = $2",
		viper.GetString("db.postgres.tableNames.portfolio"))
	count := make([]int, 1)
	err := r.db.Select(&count, query, userID, portfolioID)
	if err != nil {
		return false, err
	}

	//log.Println("IsPortfoliosOwner(): ", "userID =", userID, " portfolioID =", portfolioID, count)
	if count[1] == 1 {
		return true, nil
	} else if count[1] == 0 {
		return false, nil
	} else {
		return false, errors.New("unknown error on portfolio's owner")
	}
}
