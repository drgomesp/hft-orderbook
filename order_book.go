package main

import (
	"errors"
	"github.com/shopspring/decimal"
)

var (
	ErrOrderExists   = errors.New("order already exists in order book")
	ErrInvalidVolume = errors.New("invalid volume")
	ErrInvalidPrice  = errors.New("invalid price")
)

type OrderBook struct {
	orders map[string]*Order // orders map where the key is the order id

	asks *OrderSide
	bids *OrderSide
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		orders: make(map[string]*Order),
		asks:   NewOrderSide(Sell),
		bids:   NewOrderSide(Buy),
	}
}

func (b *OrderBook) GetBestBid() (decimal.Decimal, *OrderQueue) {
	if b.bids.MaxPriceQueue().Len() > 0 {
		return b.bids.MaxPriceQueue().First().Value.(*Order).Price, b.bids.MaxPriceQueue()
	}

	return decimal.Zero, nil
}

func (b *OrderBook) GetBestAsk() (decimal.Decimal, *OrderQueue) {
	if b.asks.MinPriceQueue().Len() > 0 {
		return b.asks.MinPriceQueue().First().Value.(*Order).Price, b.asks.MinPriceQueue()
	}

	return decimal.Zero, nil
}

func (b *OrderBook) GetWorstBid() *Order {
	if b.bids.MaxPriceQueue().Len() > 0 {
		return b.bids.MinPriceQueue().Last()
	}

	return nil
}

func (b *OrderBook) GetWorstAsk() *Order {
	if b.asks.MinPriceQueue().Len() > 0 {
		return b.asks.MaxPriceQueue().Last()
	}

	return nil
}

func (b *OrderBook) AddOrder(order *Order) (trades []*Trade, err error) {
	if _, ok := b.orders[order.Id]; ok {
		return nil, ErrOrderExists
	}

	if order.Volume.IsZero() {
		return nil, ErrInvalidVolume
	}

	if order.Execution == Market {
		if order.Kind == Buy && !b.asks.IsEmpty() {
			worstAsk := b.GetWorstAsk()
			order.Price = worstAsk.Price
			order.Type = GoodTilCanceled
		} else if order.Kind == Sell && !b.bids.IsEmpty() {
			worstBid := b.GetWorstBid()
			order.Price = worstBid.Price
			order.Type = GoodTilCanceled
		} else {
			return nil, nil
		}
	}

	if order.Price.IsZero() {
		return nil, ErrInvalidPrice
	}

	if order.Kind == Buy {
		b.bids.Append(order)
	} else {
		b.asks.Append(order)
	}

	b.orders[order.Id] = order

	return b.MatchOrders()
}

func (b *OrderBook) MatchOrders() (trades []*Trade, err error) {
	for !b.bids.IsEmpty() && !b.asks.IsEmpty() {
		bestBid, bids := b.GetBestBid()
		bestAsk, asks := b.GetBestAsk()

		for !bids.IsEmpty() && !asks.IsEmpty() {
			if bestBid.GreaterThanOrEqual(bestAsk) {
				bidEl := bids.First()
				bid := bidEl.Value.(*Order)

				askEl := asks.First()
				ask := askEl.Value.(*Order)

				volume := decimal.Min(bid.RemainingVolume, ask.RemainingVolume)

				if err = bid.Fill(volume); err != nil {
					return nil, err
				}

				if err = ask.Fill(volume); err != nil {
					return nil, err
				}

				if bid.IsFilled() {
					bids.Remove(bidEl)
					b.bids.Remove(bidEl)
					delete(b.orders, bid.Id)
				} else {
					partial := NewLimitOrder(bid.Kind, bid.Id, bid.RemainingVolume, bid.Price)
					bids.Update(bidEl, partial)
				}

				if ask.IsFilled() {
					asks.Remove(askEl)
					b.asks.Remove(askEl)
					delete(b.orders, ask.Id)
				} else {
					partial := NewLimitOrder(ask.Kind, ask.Id, ask.RemainingVolume, ask.Price)
					asks.Update(askEl, partial)
				}

				trades = append(trades, NewTrade(bid, ask, volume, ask.Price))
			}
		}
	}

	return trades, err
}

func (b *OrderBook) String() string {
	return "\r\n------------- asks --------------" +
		b.asks.String() +
		"\r\n------------- bids --------------" +
		b.bids.String()
}
