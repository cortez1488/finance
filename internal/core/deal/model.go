package deal

type Symbol struct {
	ID       int    `db:"id"`
	Abbr     string `db:"abbr"`
	FullName string `db:"full_name"`
	Price    float64
}

type Deal struct {
	ID          int     `db:"id"`
	Type        string  `db:"type"`
	SymbolID    int     `db:"symbol_id"`
	SymbolPrice float64 `db:"symbol_price"`
	Number      int     `db:"number"`
	Amount      float64 `db:"amount"`
	Date        string  `db:"date"`
	PortfolioID int     `db:"portfolio_id"`
	UserID      int     `db:"user_id"`
}

type ActiveShare struct {
	ID          int
	Price       float64
	Number      int
	PortfolioID int
	SymbolID    int
	DealID      int
}

type ActType string
