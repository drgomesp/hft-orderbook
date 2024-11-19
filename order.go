package main

import (
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	Side   OrderKind       `json:"side"`
	Id     string          `json:"id"`
	Time   time.Time       `json:"time"`
	Price  decimal.Decimal `json:"price"`
	Volume decimal.Decimal `json:"volume"`
}

func NewOrder(id string, side OrderKind, volume, price decimal.Decimal) interface{} {
	return &Order{
		Side:   side,
		Id:     id,
		Time:   time.Now(),
		Volume: volume,
		Price:  price,
	}
}
