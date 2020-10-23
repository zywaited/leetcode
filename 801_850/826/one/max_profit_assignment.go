package one

import "sort"

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type node struct {
		profit     int
		difficulty int
	}
	nodes := make([]*node, 0, len(difficulty))
	for index := range difficulty {
		nodes = append(nodes, &node{
			profit:     profit[index],
			difficulty: difficulty[index],
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].difficulty < nodes[j].difficulty {
			return true
		}
		if nodes[i].difficulty == nodes[j].difficulty {
			return nodes[i].profit < nodes[j].profit
		}
		return false
	})
	sort.Ints(worker)

	profits := 0
	fetchIndex := 0
	maxProfit := nodes[fetchIndex].profit
	for _, w := range worker {
		if nodes[fetchIndex].difficulty > w {
			continue
		}
		for ; fetchIndex+1 < len(nodes) && nodes[fetchIndex+1].difficulty <= w; fetchIndex++ {
			if nodes[fetchIndex+1].profit > maxProfit {
				maxProfit = nodes[fetchIndex+1].profit
			}
		}
		profits += maxProfit
	}
	return profits
}
