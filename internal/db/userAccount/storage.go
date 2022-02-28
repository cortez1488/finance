package userAccount

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"myFinanceTask/internal/core/deal"
	"myFinanceTask/internal/core/user_account"
	"myFinanceTask/internal/handler/rest"
)

type userAccountStorage struct {
	db *sqlx.DB
}

func NewUserAccountStorage(db *sqlx.DB) *userAccountStorage {
	return &userAccountStorage{db: db}
}

func (r *userAccountStorage) CreatePortfolio(userId int, dto rest.PortfolioDTO) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, user_id) VALUES ($1, $2) RETURNING id", "portfolio")
	var id int
	err := r.db.Get(&id, query, dto.Name, userId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userAccountStorage) GetPortfolio(userId int, id int) (user_account.Portfolio, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 and user_id = $2", "portfolio")
	var output user_account.Portfolio
	err := r.db.Get(&output, query, id, userId)
	if err != nil {
		return user_account.Portfolio{}, errors.New("not owner of portfolio" + err.Error())
	}

	return output, nil
}

func (r *userAccountStorage) GetPortfolioList(userId int) ([]user_account.Portfolio, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", "portfolio")
	var output []user_account.Portfolio
	err := r.db.Select(&output, query, userId)
	if err != nil {
		return output, err
	}
	return output, nil

}

func (r *userAccountStorage) History(userId int, timeAfter, timeBefore string) ([]deal.Deal, error) {
	return []deal.Deal{}, nil
}
