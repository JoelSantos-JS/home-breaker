package entity

type OrderQueue struct {
	Orders []*Order
}

// Swap implements heap.Interface.
func (*OrderQueue) Swap(i int, j int) {
	panic("unimplemented")
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price

}

func (oq *OrderQueue) swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func newOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
