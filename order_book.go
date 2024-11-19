package main

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
)

var (
	ErrOrderExists = errors.New("order already exists in order book")
)

type OrderBook struct {
	asks *OrderSide
	bids *OrderSide

	orders map[string]*Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		asks: NewOrderSide(Sell),
		bids: NewOrderSide(Buy),
	}
}

func (ob *OrderBook) PlaceMarketOrder(kind OrderKind, id string, volume decimal.Decimal) {
	panic("TODO")
}

// PlaceLimitOrder places a new limit order to the order book.
func (ob *OrderBook) PlaceLimitOrder(kind OrderKind, id string, price, volume decimal.Decimal) ([]*Order, error) {
	if _, ok := ob.orders[id]; ok {
		return nil, ErrOrderExists
	}

	if volume.IsZero() || volume.IsNegative() {
		return nil, errors.New("invalid volume")
	}

	if price.IsZero() || price.IsNegative() {
		return nil, errors.New("invalid price")
	}

	var (
		side           *OrderSide
		comparatorFunc func(decimal.Decimal) bool
		iter           func() *OrderQueue
	)

	if kind == Buy {
		side = ob.bids
		comparatorFunc = price.GreaterThanOrEqual
		iter = ob.asks.MinPriceQueue
	} else {
		side = ob.asks
		comparatorFunc = price.LessThanOrEqual
		iter = ob.bids.MaxPriceQueue
	}

	spew.Dump(side)
	_, _ = comparatorFunc, iter

	return nil, nil
}

func (ob *OrderBook) String() string {
	return ob.asks.String() +
		"\r\n-----\r\n" +
		ob.bids.String()
}
