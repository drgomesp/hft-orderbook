package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestOrderBook_PlaceLimitOrder(t *testing.T) {
	ob := NewOrderBook()

	for i := 50; i < 100; i += 10 {
		volume := decimal.NewFromInt(int64(rand.Intn(10)))

		if !volume.IsZero() {
			price := decimal.NewFromInt(int64(i))
			done, err := ob.PlaceLimitOrder(Buy, fmt.Sprintf("buy-%d", i), volume, price)

			assert.Nil(t, err)
			assert.Empty(t, done)
		}
	}
}
