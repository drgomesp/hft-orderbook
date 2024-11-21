package main

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestOrderBook_PlaceLimitOrder(t *testing.T) {
	ob := NewOrderBook()

	//ob.PlaceLimitOrder(Buy, "buy-1", decimal.NewFromInt(5), decimal.NewFromFloat(96.98))

	//ob.PlaceLimitOrder(Sell, "sell-1", decimal.NewFromInt(10), decimal.NewFromFloat(96.98))
	//ob.PlaceLimitOrder(Sell, "sell-2", decimal.NewFromInt(20), decimal.NewFromFloat(96.96))
	//ob.PlaceLimitOrder(Sell, "sell-3", decimal.NewFromInt(15), decimal.NewFromFloat(96.94))
	//ob.PlaceLimitOrder(Sell, "sell-4", decimal.NewFromInt(25), decimal.NewFromFloat(96.94))
	//
	//ob.PlaceLimitOrder(Buy, "buy-2", decimal.NewFromInt(35), decimal.NewFromFloat(96.95))

	_, _ = ob.AddOrder(NewLimitOrder(Sell, "sell-1", decimal.NewFromInt(10), decimal.NewFromFloat(96.98)))
	_, _ = ob.AddOrder(NewLimitOrder(Sell, "sell-2", decimal.NewFromInt(20), decimal.NewFromFloat(96.96)))
	_, _ = ob.AddOrder(NewLimitOrder(Sell, "sell-3", decimal.NewFromInt(15), decimal.NewFromFloat(96.94)))
	_, _ = ob.AddOrder(NewLimitOrder(Sell, "sell-4", decimal.NewFromInt(25), decimal.NewFromFloat(96.94)))

	trades, err := ob.AddOrder(NewLimitOrder(Buy, "buy-1", decimal.NewFromInt(35), decimal.NewFromFloat(96.95)))
	assert.NoError(t, err)

	log.Print(ob.String())
	for _, trade := range trades {
		log.Println(trade)
	}

	//trades, err = ob.AddOrder(NewLimitOrder(Buy, "buy-1", decimal.NewFromInt(10), decimal.NewFromFloat(96.99)))
	trades, err = ob.AddOrder(NewMarketOrder(Buy, "buy-2", decimal.NewFromInt(10)))
	assert.NoError(t, err)

	log.Print(ob.String())
	for _, trade := range trades {
		log.Println(trade)
	}

	//assert.NoError(t, err)
	//spew.Dump(trades)

	//_, _ = ob.AddOrder(NewMarketOrder(Buy, "buy-2", decimal.NewFromInt(35)))

}
