package dto

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetID       string  `json:"asset_id"`
	CurrentShares int     `json:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderID    string `json:"order_id"`
	InvestorID string `json:"investor_id"`
	AssetID    string `json:"asset_id"`
	OrderType  string `json:"current_shares"`
	Status     string `json:"status"`
	Partial    int    `json:"partial"`
	Shares     int    `json:"shares"`
}

type TranscationOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerId       string  `json:"buyer_id"`
	SellerId      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
}
