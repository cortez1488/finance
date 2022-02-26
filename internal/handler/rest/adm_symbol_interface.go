package rest

type AdmSymbolService interface {
	CreateSymbol(symbol AdmSymbolDTO) (int, error)
	SetPrice(symbol AdmPriceDTO) error
	DeleteSymbol(id int) error
}
