package deal

import (
	"errors"
	"myFinanceTask/internal/handler/rest"
	"time"
)

type dealService struct {
	repo DealStorage
}

func NewDealService(repo DealStorage) *dealService {
	return &dealService{repo: repo}
}

func (s *dealService) GetShareInfo(id int) (rest.ShareDTO, error) {
	var output rest.ShareDTO
	result, err := s.repo.GetShareInfo(id)
	if err != nil {
		return output, err
	}
	output.Abbr = result.Abbr
	output.FullName = result.FullName
	output.Price = result.Price
	return output, nil
}

func (s *dealService) GetShareListInfo() ([]rest.ShareDTO, error) {
	results, err := s.repo.GetShareListInfo()
	output := make([]rest.ShareDTO, len(results))
	if err != nil {
		return output, err
	}

	for index, res := range results {
		output[index].Abbr = res.Abbr
		output[index].FullName = res.FullName
		output[index].Price = res.Price
	}
	return output, nil
}

//--------- DEAL LOGIC -------------

func (s *dealService) BuyShares(shareID, portfolioID, userID, quantity int) (float64, error) {
	owner, err := s.repo.IsPortfoliosOwner(userID, portfolioID)
	if err != nil {
		return 0, err
	}
	if !owner {
		return 0, errors.New("you're not owner of portfolio")
	}

	share, err := s.GetShareInfo(shareID)
	if err != nil {
		return 0, err
	}

	price := share.Price
	amount := price * float64(quantity)

	err = s.repo.BuyShares(shareID, portfolioID, userID, quantity, price, amount, time.Now().String(), TypeBuy)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (s *dealService) SellShares(activeShareID, portfolioID, userID, quantity int) (float64, error) {
	owner, err := s.repo.IsPortfoliosOwner(userID, portfolioID)
	if err != nil {
		return 0, err
	}
	if !owner {
		return 0, errors.New("you're not owner of portfolio")
	}

	share, err := s.repo.GetShareInfoOfActiveShareID(activeShareID)
	if err != nil {
		return 0, err
	}

	price := share.Price
	amount := price * float64(quantity)

	err = s.repo.SellShares(activeShareID, share.ID, portfolioID, userID, quantity, price, amount, time.Now().String(), TypeSell)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
