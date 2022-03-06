package rest

type PricesRefreshService interface {
	RefreshPrices(*[]RefreshDBDTO) error
}
