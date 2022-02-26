package rest

type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdmSymbolDTO struct {
	Abbr     string `json:"abbr"`
	FullName string `json:"fullName"`
}

type AdmPriceDTO struct {
	Abbr  string  `json:"abbr"`
	Price float64 `json:"price"`
}

type PortfolioDTO struct {
	Name    string `json:"name"`
	Account int    `json:"account,omitempty"`
}
