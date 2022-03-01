package rest

type DealService interface {
	GetShareInfo(id int) (ShareDTO, error)
	GetShareListInfo() ([]ShareDTO, error)

	BuyShares(shareID, portfolioID, userID, quantity int) (float64, error)
	SellShares(activeShareID, portfolioID, userID, quantity int) (float64, error)
}
