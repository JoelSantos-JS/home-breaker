package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order         []*Order
	Transaction   []*Transaction
	OrdersChan    chan *Order
	OrdersChanOut chan *Order
	Wg            *sync.WaitGroup
}

func NewBook(orderChan chan *Order, orderChanout chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transaction:   []*Transaction{},
		OrdersChan:    orderChan,
		OrdersChanOut: orderChanout,
		Wg:            wg,
	}
}

func (b *Book) trade() {
	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	/*
		buyOrders := newOrderQueue()
		sellOrders := newOrderQueue()
	*/
	//	heap.Init(buyOrders)
	//	heap.Init(sellOrders)

	for order := range b.OrdersChan {
		asset := order.Asset.Id

		if buyOrders[asset] == nil {
			buyOrders[asset] = newOrderQueue()
			heap.Init(buyOrders[asset])
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = newOrderQueue()
			heap.Init(sellOrders[asset])
		}

		if order.OrderType == "BUY" {
			buyOrders[asset].Push(order)
			if sellOrders[asset].Len() > 0 && sellOrders[asset].Orders[0].Price <= order.Price {
				sellOrder := sellOrders[asset].Pop().(*Order)

				if sellOrder.PendingShates > 0 {
					transcation := newTransaction(sellOrder, order, order.Shates, sellOrder.Price)
					b.AddTransaction(transcation, b.Wg)
					sellOrder.Transaction = append(sellOrder.Transaction)
					order.Transaction = append(order.Transaction, transcation)
					b.OrdersChanOut <- sellOrder
					b.OrdersChanOut <- order
					if sellOrder.PendingShates > 0 {
						sellOrders[asset].Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == "SELL" {
			sellOrders[asset].Push(order)

			if buyOrders[asset].Len() > 0 && buyOrders[asset].Orders[0].Price >= order.Price {
				buyOrder := buyOrders[asset].Pop().(*Order)

				if buyOrder.PendingShates > 0 {
					transaction := newTransaction(buyOrder, order, order.Shates, buyOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					buyOrder.Transaction = append(buyOrder.Transaction)
					order.Transaction = append(order.Transaction, transaction)
					b.OrdersChanOut <- buyOrder
					b.OrdersChanOut <- order
					if buyOrder.PendingShates > 0 {
						buyOrders[asset].Push(buyOrder)

					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	sellingShares := transaction.SellingOrder.PendingShates
	buyingShares := transaction.BuyingOrder.PendingShates

	minShares := sellingShares

	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.Id, -minShares)
	transaction.SellingOrder.PendingShates -= minShares
	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.Id, minShares)
	transaction.BuyingOrder.PendingShates -= minShares

	transaction.calcTotal(transaction.Shares, transaction.BuyingOrder.Price)
	transaction.closedBuyTransaction()
	transaction.closedSellTransaction()

	b.Transaction = append(b.Transaction, transaction)
}
