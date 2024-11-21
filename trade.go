package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Trade struct {
	bid    *Order
	ask    *Order
	volume decimal.Decimal
	price  decimal.Decimal
}

func NewTrade(bid, ask *Order, volume, price decimal.Decimal) *Trade {
	return &Trade{
		bid:    bid,
		ask:    ask,
		volume: volume,
		price:  price,
	}
}

func (t *Trade) String() string {
	return fmt.Sprintf("bid: %s | ask: %s | volume: %s | price: %s", t.bid.Price, t.ask.Price, t.volume, t.price)
}
