package price_refresh

import (
	"log"
	"myFinanceTask/internal/handler/rest"
	"sort"
	"strconv"
	"time"
)

type priceRefreshService struct {
	repo PricesRefreshStorage
}

func NewPriceRefreshService(repo PricesRefreshStorage) *priceRefreshService {
	return &priceRefreshService{repo: repo}
}

func (s *priceRefreshService) RefreshPrices(rawData *[]rest.RefreshDBDTO) error {
	log.Println("Started refreshing redis prices.")
	existingSymbolsStrings, err := s.repo.GetCurrentSymbols()
	if err != nil {
		return err
	}
	sort.Strings(existingSymbolsStrings)

	start := time.Now()
	serviceData, err := makeServiceData(rawData, existingSymbolsStrings)
	if err != nil {
		return err
	}

	resTime, err := s.repo.RefreshPrices(serviceData)
	log.Println("Time fore searching existing symbols:", resTime.Sub(start))

	return nil
}

func makeServiceData(rawData *[]rest.RefreshDBDTO, existingSymbolsStrings []string) ([]Symbol, error) {
	var err error

	serviceData := make([]Symbol, len(existingSymbolsStrings))
	for index, symbol := range *rawData {

		if binaryStringSearch(symbol.Symbol, existingSymbolsStrings) {
			serviceData[index].Price, err = strconv.ParseFloat(symbol.Price, 64)
			if err != nil {
				return nil, err
			}
			serviceData[index].Symbol = symbol.Symbol
		}
	}
	return serviceData, nil
}

func binaryStringSearch(needle string, haystack []string) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
