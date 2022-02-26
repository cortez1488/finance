package admSymbol

import "myFinanceTask/internal/handler/rest"

type admSymbolService struct {
	repo AdmSymbol
}

func NewAdmSymbolService(repo AdmSymbol) *admSymbolService {
	return &admSymbolService{repo: repo}
}

func (s *admSymbolService) CreateSymbol(symbol rest.AdmSymbolDTO) (int, error) {
	//IS ADMIN:
	return s.repo.CreateSymbol(symbol)
}

func (s *admSymbolService) SetPrice(price float64) error {
	//IS ADMIN:
	return s.repo.SetPrice(price)
}

func (s *admSymbolService) DeleteSymbol(id int) error {
	//IS ADMIN:
	return s.repo.DeleteSymbol(id)
}
