package rest

//----------- SYMBOL ADMINISTRATING -------------
type AdmSymbolDTO struct {
	Abbr     string `json:"abbr"`
	FullName string `json:"fullName"`
}

type AdmPriceDTO struct {
	Abbr  string  `json:"abbr"`
	Price float64 `json:"price"`
}

//----------- USER ACCOUNT -------------
type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PortfolioDTO struct {
	Name    string  `json:"name"`
	Account float64 `json:"account"` //float
}

type HistoryDTO struct {
	Type        string  `db:"type"`
	SymbolAbbr  string  `db:"symbol_id"`
	SymbolPrice float64 `db:"symbol_price"`
	Number      int     `db:"number"`
	Amount      float64 `db:"amount"`
	Date        string  `db:"date"`
}

//----------- DEAL  -------------
type ShareDTO struct {
	Abbr     string  `json:"abbr"`
	FullName string  `json:"fullName"`
	Price    float64 `json:"price"`
}

type CreateDealDTO struct {
	ShareID     int `json:"shareID"`
	PortfolioID int `json:"portfolioID"`
	Quantity    int `json:"quantity"`
}

//---------- REFRESH -----------

type RefreshDBDTO struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
