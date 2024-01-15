package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shates        int
	PendingShates int
	Price         float64
	OrderType     string
	Status        string
	Transaction   []*Transaction
}

func newOrder(orderID string, investor *Investor, asset *Asset, shates int, price float64, orderType string, status string) *Order {
	return &Order{
		ID:          orderID,
		Investor:    investor,
		Asset:       asset,
		Shates:      shates,
		Price:       price,
		OrderType:   orderType,
		Status:      "OPEN",
		Transaction: []*Transaction{},
	}
}
