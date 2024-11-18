package main

type OrderBook struct {
	asks *OrderSide
	bids *OrderSide
}

func NewOrderBook(id string) *OrderBook {
	return &OrderBook{
		asks: NewOrderSide(Sell),
		bids: NewOrderSide(Buy),
	}
}

func (b *OrderBook) AddOrder(order *Order) {

}
