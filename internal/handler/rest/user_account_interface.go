package rest

type UserAccountService interface {
	CreatePortfolio(userId int, dto PortfolioDTO) (int, error)
	GetPortfolio(userId, id int) (PortfolioDTO, error) //Переделать на DTO
	GetPortfolioList(userId int) ([]PortfolioDTO, error)

	History(userId int) ([]HistoryDTO, error) //Переделать на DTO
}
