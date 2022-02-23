package rest

type AdmSymbolService interface {
	CreateSymbol(symbol AdmSymbolDTO) (int, error)
	SetPrice(price float64) error
	DeleteSymbol(id int) error
}
