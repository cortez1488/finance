package rest

import (
	"myFinanceTask/internal/core/deal"
)

type UserAccountService interface {
	CreatePortfolio(userId int, dto PortfolioDTO) (int, error)
	GetPortfolio(userId, id int) (PortfolioDTO, error) //Переделать на DTO
	GetPortfolioList(userId int) ([]PortfolioDTO, error)

	History(userId int, timeAfter, timeBefore string) ([]deal.Deal, error) //Переделать на DTO
}
