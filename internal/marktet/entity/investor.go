package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPositiron
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPositiron{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPositiron) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)

}

func (i *Investor) UpdateAssetPosition(AssetId string, qtdShares int) {
	assetPostion := i.GetAssetPosition(AssetId)
	if assetPostion != nil {
		i.AssetPosition = append(i.AssetPosition, NewIvestorAssetPosition(AssetId, qtdShares))
	} else {
		assetPostion.Shares += qtdShares
	}

}

func (i *Investor) GetAssetPosition(AssetId string) *InvestorAssetPositiron {

	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetId == AssetId {
			return assetPosition
		}
	}

	return nil

}

func NewIvestorAssetPosition(assetId string, qtdShares int) *InvestorAssetPositiron {
	return &InvestorAssetPositiron{
		AssetId: assetId,
		Shares:  qtdShares,
	}
}

type InvestorAssetPositiron struct {
	AssetId string
	Shares  int
}
