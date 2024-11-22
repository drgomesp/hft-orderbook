package main

import (
	"container/list"
	"github.com/shopspring/decimal"
)

type OrderQueue struct {
	volume decimal.Decimal
	price  decimal.Decimal
	orders *list.List
}

func NewOrderQueue(price decimal.Decimal) *OrderQueue {
	return &OrderQueue{
		volume: decimal.Zero,
		price:  price,
		orders: list.New(),
	}
}

func (q *OrderQueue) Price() decimal.Decimal {
	return q.price
}

func (q *OrderQueue) Volume() decimal.Decimal {
	return q.volume
}

func (q *OrderQueue) Append(order *Order) *list.Element {
	q.volume = q.volume.Add(order.Volume)
	return q.orders.PushBack(order)
}

func (q *OrderQueue) First() *list.Element {
	return q.orders.Front()
}

func (q *OrderQueue) Last() *Order {
	return q.orders.Back().Value.(*Order)
}

func (q *OrderQueue) IsEmpty() bool {
	return q.orders.Len() == 0
}

func (q *OrderQueue) Len() int {
	return q.orders.Len()
}

func (q *OrderQueue) Update(e *list.Element, partial *Order) *list.Element {
	q.volume = partial.RemainingVolume
	e.Value = partial

	return e
}

func (q *OrderQueue) Remove(e *list.Element) *Order {
	q.volume = q.volume.Sub(e.Value.(*Order).Volume)
	return q.orders.Remove(e).(*Order)
}
