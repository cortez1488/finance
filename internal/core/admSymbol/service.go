package admSymbol

import "myFinanceTask/internal/handler/rest"

type admSymbolService struct {
	repo AdmSymbol
}

func NewAdmSymbolService(repo AdmSymbol) *admSymbolService {
	return &admSymbolService{repo: repo}
}

func (s *admSymbolService) CreateSymbol(symbol rest.AdmSymbolDTO) (int, error) {
	return s.repo.CreateSymbol(symbol)
}

func (s *admSymbolService) SetPrice(symbol rest.AdmPriceDTO) error {
	return s.repo.SetPrice(symbol)
}

func (s *admSymbolService) DeleteSymbol(id int) error {
	return s.repo.DeleteSymbol(id)
}
