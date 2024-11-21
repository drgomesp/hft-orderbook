package main

//
//import (
//	"container/list"
//	"errors"
//	"github.com/shopspring/decimal"
//	"log"
//)
//
//var (
//	ErrOrderExists   = errors.New("order already exists in order book")
//	ErrInvalidVolume = errors.New("invalid volume")
//	ErrInvalidPrice  = errors.New("invalid price")
//)
//
//type OrderBook struct {
//	orders map[string]*list.Element // orders map where the key is the order id
//
//	asks *OrderSide
//	bids *OrderSide
//}
//
//func NewOrderBook() *OrderBook {
//	return &OrderBook{
//		orders: make(map[string]*list.Element),
//
//		asks: NewOrderSide(Sell),
//		bids: NewOrderSide(Buy),
//	}
//}
//
//func (b *OrderBook) PlaceMarketOrder(kind OrderKind, id string, volume decimal.Decimal) {
//	panic("TODO")
//}
//
//// PlaceLimitOrder places a new limit order to the order book.
//func (b *OrderBook) PlaceLimitOrder(orderKind OrderKind, id string, volume, price decimal.Decimal) ([]*Trade, error) {
//	if _, ok := b.orders[id]; ok {
//		return nil, ErrOrderExists
//	}
//
//	if volume.IsZero() || volume.IsNegative() {
//		return nil, ErrInvalidVolume
//	}
//
//	if price.IsZero() || price.IsNegative() {
//		return nil, ErrInvalidPrice
//	}
//
//	var (
//		sideToAdd *OrderSide
//		compare   func(decimal.Decimal) bool
//		iterator  func() *OrderQueue
//	)
//
//	if orderKind == Buy {
//		sideToAdd = b.bids
//		//sideToProcess = b.asks
//		compare = price.GreaterThanOrEqual
//		iterator = b.asks.MinPriceQueue
//	} else {
//		sideToAdd = b.asks
//		//sideToProcess = b.bids
//		compare = price.LessThanOrEqual
//		iterator = b.bids.MaxPriceQueue
//	}
//
//	trades := make([]*Trade, 0)
//	volumeLeft := volume
//	bestPrice := iterator()
//
//	// while there is a best price and the opposite side is not
//	// empty, and as long as the remaining order volume is not zero
//	for bestPrice != nil && volumeLeft.GreaterThan(decimal.Zero) && compare(bestPrice.Price()) {
//		trades, volumeLeft = b.executeOrders(id, bestPrice, volumeLeft)
//		bestPrice = iterator()
//	}
//
//	if volumeLeft.GreaterThan(decimal.Zero) {
//		order := NewOrder(id, orderKind, volume, price)
//		b.orders[order.Id] = sideToAdd.Append(order)
//	}
//
//	return trades, nil
//}
//
//func (b *OrderBook) executeOrders(id string, orderQueue *OrderQueue, remainingVolume decimal.Decimal) (
//	trades []*Trade,
//	volumeLeft decimal.Decimal,
//) {
//	volumeLeft = remainingVolume
//
//	for orderQueue.Len() > 0 && volumeLeft.GreaterThan(decimal.Zero) {
//		order := orderQueue.First()
//
//		// if the order can fill the entire volume left
//		if volumeLeft.LessThan(order.Volume) {
//			partial := NewOrder(
//				order.Id,
//				order.Kind,
//				order.Volume.Sub(volumeLeft),
//				orderQueue.Price(),
//			)
//
//			orderQueue.Update(order, partial)
//			trades = append(trades, NewTrade(id, partial.Id, volumeLeft, partial.Price))
//			volumeLeft = decimal.Zero
//		} else {
//			volumeLeft = volumeLeft.Sub(order.Volume)
//			canceled := b.CancelOrder(order.Id)
//			trades = append(trades, NewTrade(id, canceled.Id, canceled.Volume, canceled.Price))
//
//			log.Printf("canceled order (%v)", canceled)
//		}
//	}
//
//	return trades, volumeLeft
//}
//
//func (b *OrderBook) CancelOrder(id string) *Order {
//	e, ok := b.orders[id]
//	if !ok {
//		return nil
//	}
//
//	delete(b.orders, id)
//	order := e.Value.(*Order)
//
//	if order.Kind == Buy {
//		return b.bids.Remove(e)
//	}
//
//	return b.asks.Remove(e)
//}
//
//func (b *OrderBook) String() string {
//	return "\r\n------------- asks --------------" +
//		b.asks.String() + "\r\n------------- bids --------------" + b.bids.String()
//}
