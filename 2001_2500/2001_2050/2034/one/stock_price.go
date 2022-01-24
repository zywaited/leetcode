package one

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type StockPrice struct {
	prices          *redblacktree.Tree
	timestamp2Price map[int]int
	maxTimestamp    int
}

func Constructor() StockPrice {
	return StockPrice{
		prices:          redblacktree.NewWithIntComparator(),
		timestamp2Price: make(map[int]int),
	}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	if sp.maxTimestamp < timestamp {
		sp.maxTimestamp = timestamp
	}
	if sp.timestamp2Price[timestamp] > 0 {
		times, _ := sp.prices.Get(sp.timestamp2Price[timestamp])
		if times.(int) <= 1 {
			sp.prices.Remove(sp.timestamp2Price[timestamp])
		} else {
			sp.prices.Put(sp.timestamp2Price[timestamp], times.(int)-1)
		}
	}
	times := 0
	node, found := sp.prices.Get(price)
	if found {
		times = node.(int)
	}
	sp.prices.Put(price, times+1)
	sp.timestamp2Price[timestamp] = price
}

func (sp *StockPrice) Current() int {
	return sp.timestamp2Price[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	return sp.prices.Right().Key.(int)
}

func (sp *StockPrice) Minimum() int {
	return sp.prices.Left().Key.(int)
}
