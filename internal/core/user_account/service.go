package user_account

import (
	"myFinanceTask/internal/core/deal"
	"myFinanceTask/internal/handler/rest"
)

type UserAccountService struct {
	repo UserAccountStorage
}

func NewUserAccountService(repo UserAccountStorage) *UserAccountService {
	return &UserAccountService{repo: repo}
}

func (s *UserAccountService) CreatePortfolio(userId int, dto rest.PortfolioDTO) (int, error) {
	return s.repo.CreatePortfolio(userId, dto)
}

func (s *UserAccountService) GetPortfolio(userId int, id int) (rest.PortfolioDTO, error) {
	bsns, err := s.repo.GetPortfolio(userId, id)
	if err != nil {
		return rest.PortfolioDTO{}, err
	}
	dto := rest.PortfolioDTO{Name: bsns.Name, Account: bsns.Account}
	return dto, nil
}
func (s *UserAccountService) GetPortfolioList(userId int) ([]rest.PortfolioDTO, error) {
	bsns, err := s.repo.GetPortfolioList(userId)
	dto := make([]rest.PortfolioDTO, len(bsns))
	if err != nil {
		return dto, err
	}

	for index, obj := range bsns {
		dto[index].Name = obj.Name
		dto[index].Account = obj.Account
	}
	return dto, nil
}

func (s *UserAccountService) History(userId int, timeAfter, timeBefore string) ([]deal.Deal, error) {
	return s.repo.History(userId, timeAfter, timeBefore)
}
