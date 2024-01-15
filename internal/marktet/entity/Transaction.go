package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func newTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
	}

}

func (t *Transaction) calcTotal(shares int, price float64) {
	t.Total = float64(shares) * price

}

func (t *Transaction) closedBuyTransaction() {

	if t.BuyingOrder.PendingShates == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) closedSellTransaction() {

	if t.SellingOrder.PendingShates == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}
