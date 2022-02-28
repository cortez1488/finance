package deal

type DealStorage interface {
	GetShareInfo(id int) (Symbol, error)
	GetShareListInfo() ([]Symbol, error)

	BuyShares(shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date string, dType ActType) error
	SellShares(activeShareID, shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date string, dType ActType) error
}
