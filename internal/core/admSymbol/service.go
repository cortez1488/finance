package admSymbol

import (
	"log"
	"myFinanceTask/internal/handler/rest"
)

type admSymbolService struct {
	repo AdmSymbol
}

func NewAdmSymbolService(repo AdmSymbol) *admSymbolService {
	return &admSymbolService{repo: repo}
}

func (s *admSymbolService) CreateSymbol(symbol rest.AdmSymbolDTO) (int, error) {
	log.Println("Creating symbol", symbol)
	return s.repo.CreateSymbol(symbol)
}

func (s *admSymbolService) SetPrice(symbol rest.AdmPriceDTO) error {
	log.Println("Set symbol price", symbol)
	return s.repo.SetPrice(symbol)
}

func (s *admSymbolService) DeleteSymbol(id int) error {
	log.Println("Delete symbol", id)
	return s.repo.DeleteSymbol(id)
}
