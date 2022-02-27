package deal

type Symbol struct {
	ID       int    `db:"id"`
	Abbr     string `db:"abbr"`
	FullName string `db:"full_name"`
	Price    float64
}

type Deal struct {
	ID          int
	Type        string
	SymbolID    int
	SymbolPrice float64
	Number      int
	Amount      float64
	Date        string
	PortfolioID int
	UserID      int
}

type ActiveShare struct {
	ID          int
	Price       float64
	Number      int
	PortfolioID int
	SymbolID    int
	DealID      int
}
