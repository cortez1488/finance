package user_account

import (
	"log"
	"myFinanceTask/internal/core/deal"
	"myFinanceTask/internal/handler/rest"
)

type userAccountService struct {
	repo UserAccountStorage
}

func NewUserAccountService(repo UserAccountStorage) *userAccountService {
	return &userAccountService{repo: repo}
}

func (s *userAccountService) CreatePortfolio(userId int, dto rest.PortfolioDTO) (int, error) {
	log.Printf("Creating portfolio: userID: %d, %v", userId, dto)
	return s.repo.CreatePortfolio(userId, dto)
}

func (s *userAccountService) GetPortfolio(userId int, id int) (rest.PortfolioDTO, error) {
	log.Printf("Getting portfolio: userID: %d", userId)
	bsns, err := s.repo.GetPortfolio(userId, id)
	if err != nil {
		return rest.PortfolioDTO{}, err
	}
	dto := rest.PortfolioDTO{Name: bsns.Name, Account: bsns.Account}
	return dto, nil
}
func (s *userAccountService) GetPortfolioList(userId int) ([]rest.PortfolioDTO, error) {
	log.Printf("Getting portfolio: userID: %d", userId)
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

func (s *userAccountService) History(userId int, timeAfter, timeBefore string) ([]deal.Deal, error) {
	log.Printf("Getting history: userID: %d", userId)
	return s.repo.History(userId, timeAfter, timeBefore)
}
