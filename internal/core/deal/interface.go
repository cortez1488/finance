package deal

import "time"

type DealStorage interface {
	GetShareInfo(id int) (Symbol, error)
	GetShareListInfo() ([]Symbol, error)

	BuyShares(shareID, portfolioID, userID, quantity int, symbolPrice, amount float64, date time.Time, dType ActType) error
	SellShares(activeShareID, portfolioID, shareID, userID, quantity int, symbolPrice, amount float64, date time.Time, dType ActType) error
	GetShareInfoOfActiveShareID(activeShareID int) (Symbol, error)
	IsPortfoliosOwner(userID, portfolioID int) (bool, error)
}
