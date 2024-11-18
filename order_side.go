package main

import (
	rbtx "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/shopspring/decimal"
)

type OrderSide struct {
	kind        Kind
	priceTree   *rbtx.Tree
	prices      map[decimal.Decimal]*OrderQueue
	volume      decimal.Decimal
	size, depth int
}

func NewOrderSide(kind Kind) *OrderSide {
	return &OrderSide{
		kind: kind,
		priceTree: rbtx.NewWith(func(a, b interface{}) int {
			return a.(decimal.Decimal).Cmp(b.(decimal.Decimal))
		}),
		prices: make(map[decimal.Decimal]*OrderQueue),
		volume: decimal.Zero,
		size:   0,
		depth:  0,
	}
}
