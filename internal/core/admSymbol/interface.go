package admSymbol

import "myFinanceTask/internal/handler/rest"

type AdmSymbol interface {
	CreateSymbol(symbol rest.AdmSymbolDTO) (int, error)
	SetPrice(price float64) error
	DeleteSymbol(id int) error
}
