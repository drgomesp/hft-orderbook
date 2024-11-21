package main

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

type OrderExecution int

const (
	Market = OrderExecution(iota)
	Limit
)

type Order struct {
	Execution       OrderExecution  `json:"execution"`
	Kind            OrderKind       `json:"kind"`
	Type            OrderType       `json:"type"`
	Id              string          `json:"id"`
	Time            time.Time       `json:"time"`
	Price           decimal.Decimal `json:"price"`
	Volume          decimal.Decimal `json:"volume"`
	RemainingVolume decimal.Decimal `json:"remaining_volume"`
}

func NewMarketOrder(side OrderKind, id string, volume decimal.Decimal) *Order {
	return &Order{
		Execution:       Market,
		Kind:            side,
		Id:              id,
		Time:            time.Now(),
		Volume:          volume,
		RemainingVolume: volume,
		Price:           decimal.Zero,
	}
}

func NewLimitOrder(side OrderKind, id string, volume, price decimal.Decimal) *Order {
	return &Order{
		Execution:       Limit,
		Kind:            side,
		Id:              id,
		Time:            time.Now(),
		Volume:          volume,
		RemainingVolume: volume,
		Price:           price,
	}
}

func (o *Order) Fill(volume decimal.Decimal) error {
	if volume.GreaterThan(o.RemainingVolume) {
		return errors.New("volume is greater than remaining volume")
	}

	o.RemainingVolume = o.RemainingVolume.Sub(volume)

	return nil
}

func (o *Order) IsFilled() bool {
	return o.RemainingVolume.IsZero()
}
