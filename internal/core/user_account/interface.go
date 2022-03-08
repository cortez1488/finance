package user_account

import (
	"myFinanceTask/internal/core/deal"
	"myFinanceTask/internal/handler/rest"
)

type UserAccountStorage interface {
	CreatePortfolio(userId int, dto rest.PortfolioDTO) (int, error)
	GetPortfolio(userId, id int) (Portfolio, error)
	GetPortfolioList(userId int) ([]Portfolio, error)

	History(userId int) ([]deal.Deal, error)
	GetSymbolAbbr(symbolID int) (string, error)
}
