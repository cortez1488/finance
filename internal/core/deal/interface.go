package deal

type DealStorage interface {
	GetShareInfo(id int) (Symbol, error)
	GetShareListInfo() ([]Symbol, error)

	BuyShares(shareID, portfolioID, userID, quantity int, amount float64) float64
	SellShares(shareID, portfolioID, userID, quantity int, amount float64) float64
}
