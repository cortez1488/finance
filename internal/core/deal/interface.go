package deal

type DealStorage interface {
	GetShareInfo(id int) (Symbol, error)
	GetShareListInfo() ([]Symbol, error)
}
