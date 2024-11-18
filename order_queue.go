package main

import (
	"github.com/shopspring/decimal"
	"container/list"
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
