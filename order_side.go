package main

import (
	"container/list"
	"fmt"
	"github.com/emirpasic/gods/examples/redblacktreeextended"
	"github.com/emirpasic/gods/trees/redblacktree"
	_ "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/shopspring/decimal"
	"strings"
)

type OrderSide struct {
	kind        OrderKind
	priceTree   *redblacktreeextended.RedBlackTreeExtended
	prices      map[string]*OrderQueue
	volume      decimal.Decimal
	size, depth int
}

func NewOrderSide(kind OrderKind) *OrderSide {
	return &OrderSide{
		kind: kind,
		priceTree: &redblacktreeextended.RedBlackTreeExtended{
			Tree: redblacktree.NewWith(func(a, b interface{}) int {
				return a.(decimal.Decimal).Cmp(b.(decimal.Decimal))
			}),
		},
		prices: make(map[string]*OrderQueue),
		volume: decimal.Zero,
		size:   0,
		depth:  0,
	}
}

// Append appends an order.
func (s *OrderSide) Append(order *Order) *list.Element {
	price := order.Price

	priceQueue, ok := s.prices[price.String()]
	if !ok {
		priceQueue = NewOrderQueue(price)
		s.prices[price.String()] = priceQueue
		s.priceTree.Put(price, priceQueue)
		s.depth++
	}

	s.size++
	s.volume = s.volume.Add(order.Volume)

	return priceQueue.Append(order)
}

func (s *OrderSide) MinPriceQueue() *OrderQueue {
	if s.depth > 0 {
		if value, found := s.priceTree.GetMin(); found {
			return value.(*OrderQueue)
		}
	}

	return nil
}

func (s *OrderSide) MaxPriceQueue() *OrderQueue {
	if s.depth > 0 {
		if value, found := s.priceTree.GetMax(); found {
			return value.(*OrderQueue)
		}
	}

	return nil
}

func (s *OrderSide) GreaterThan(price decimal.Decimal) *OrderQueue {
	panic("TODO")
}

// LessThan returns nearest OrderQueue with price less than given
func (s *OrderSide) LessThan(price decimal.Decimal) *OrderQueue {
	tree := s.priceTree.Tree
	node := tree.Root

	var floor *redblacktree.Node
	for node != nil {
		if tree.Comparator(price, node.Key) > 0 {
			floor = node
			node = node.Right
		} else {
			node = node.Left
		}
	}

	if floor != nil {
		return floor.Value.(*OrderQueue)
	}

	return nil
}

func (s *OrderSide) IsEmpty() bool {
	return s.size == 0
}

func (s *OrderSide) String() string {
	sb := strings.Builder{}

	level := s.MaxPriceQueue()
	for level != nil {
		sb.WriteString(fmt.Sprintf("\n  price: %s  |  volume: %s", level.Price(), level.Volume()))
		level = s.LessThan(level.Price())
	}

	return sb.String()
}

func (s *OrderSide) Remove(e *list.Element) *Order {
	priceQueue := s.prices[e.Value.(*Order).Price.String()]
	o := priceQueue.Remove(e)

	if priceQueue.Len() == 0 {
		delete(s.prices, o.Price.String())
		s.priceTree.Remove(o.Price)
		s.depth--
	}

	s.size--
	s.volume = s.volume.Sub(o.Volume)
	return o
}
