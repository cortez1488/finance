package admSymbol

import "myFinanceTask/internal/handler/rest"

type AdmSymbol interface {
	CreateSymbol(symbol rest.AdmSymbolDTO) (int, error)
	SetPrice(symbol rest.AdmPriceDTO) error
	DeleteSymbol(id int) error
}
