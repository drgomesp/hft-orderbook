package main

import (
	"time"
	"github.com/shopspring/decimal"
)

type Order struct {
	Side   Kind            `json:"side"`
	Id     string          `json:"id"`
	Time   time.Time       `json:"time"`
	Volume decimal.Decimal `json:"volume"`
	Price  decimal.Decimal `json:"price"`
}

func NewOrder(id string, side Kind, volume, price decimal.Decimal) interface{} {
	return &Order{
		Side:   side,
		Id:     id,
		Time:   time.Now(),
		Volume: volume,
		Price:  price,
	}
}
