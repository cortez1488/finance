package deal

import (
	"myFinanceTask/internal/handler/rest"
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
