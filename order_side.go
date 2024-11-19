package main

import (
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
	prices      map[decimal.Decimal]*OrderQueue
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
		prices: make(map[decimal.Decimal]*OrderQueue),
		volume: decimal.Zero,
		size:   0,
		depth:  0,
	}
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

func (s *OrderSide) String() string {
	sb := strings.Builder{}
	level := s.MaxPriceQueue()

	for level != nil {
		sb.WriteString(
			fmt.Sprintf(
				"\n%s -> %s",
				level.Price(),
				level.Volume(),
			),
		)

		level = s.LessThan(level.Price())
	}

	return sb.String()
}

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
