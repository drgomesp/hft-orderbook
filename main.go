package main

import (
	"log"
	"github.com/shopspring/decimal"
)

func main() {
	o1 := NewOrder("sell-1", Sell, decimal.NewFromInt(10), decimal.NewFromFloat(99.56))
	o2 := NewOrder("sell-2", Sell, decimal.NewFromInt(10), decimal.NewFromFloat(99.58))
	o3 := NewOrder("sell-3", Sell, decimal.NewFromInt(10), decimal.NewFromFloat(99.60))
	o4 := NewOrder("sell-4", Sell, decimal.NewFromInt(10), decimal.NewFromFloat(99.62))
	o5 := NewOrder("sell-5", Sell, decimal.NewFromInt(10), decimal.NewFromFloat(99.64))

	log.Println("hello world")
}
