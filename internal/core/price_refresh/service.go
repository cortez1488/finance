package price_refresh

import (
	"fmt"
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
	existingSymbolsStrings, err := s.repo.GetCurrentSymbols()
	if err != nil {
		return err
	}
	sort.Strings(existingSymbolsStrings)
	fmt.Println(existingSymbolsStrings)

	serviceData := make([]Symbol, len(existingSymbolsStrings))

	start := time.Now()
	for index, symbol := range *rawData {

		if binaryStringSearch(symbol.Symbol, existingSymbolsStrings) {
			serviceData[index].Price, err = strconv.ParseFloat(symbol.Price, 64)
			if err != nil {
				return err
			}
			serviceData[index].Symbol = symbol.Symbol
		}
	}

	log.Println("Time fore searching existing symbols:", time.Since(start))
	return s.repo.RefreshPrices(&serviceData)
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
