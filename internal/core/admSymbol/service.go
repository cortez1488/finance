package admSymbol

import "myFinanceTask/internal/handler/rest"

type AdmSymbolService struct {
	repo AdmSymbol
}

func NewAdmSymbolService(repo AdmSymbol) *AdmSymbolService {
	return &AdmSymbolService{repo: repo}
}

func (s *AdmSymbolService) CreateSymbol(symbol rest.AdmSymbolDTO) (int, error) {
	//IS ADMIN:
	return s.repo.CreateSymbol(symbol)
}

func (s *AdmSymbolService) SetPrice(price float64) error {
	//IS ADMIN:
	return s.repo.SetPrice(price)
}

func (s *AdmSymbolService) DeleteSymbol(id int) error {
	//IS ADMIN:
	return s.repo.DeleteSymbol(id)
}
