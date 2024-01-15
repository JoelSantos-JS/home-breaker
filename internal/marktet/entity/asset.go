package entity

type Asset struct {
	Id          string
	Name        string
	MarketVolum int
}

func newAsset(id string, name string, marketVolum int) *Asset {
	return &Asset{
		Id:          id,
		Name:        name,
		MarketVolum: marketVolum,
	}

}
