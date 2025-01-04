package one

import "sort"

func minCost(nums []int, costs []int) int64 {
	type node struct {
		num  int64
		cost int64
	}
	nodes := make([]node, len(nums))
	for index := range nums {
		nodes[index] = node{num: int64(nums[index]), cost: int64(costs[index])}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].num < nodes[j].num
	})
	sumCost := int64(0)
	sumDiff := int64(0)
	minCost := int64(0)
	for index := range nodes {
		diff := nodes[index].num - nodes[0].num
		sumDiff += diff
		sumCost += nodes[index].cost
		minCost += nodes[index].cost * diff
	}
	prevSumCost := nodes[0].cost
	currCost := minCost
	for index := 1; index < len(nodes); index++ {
		diff := nodes[index].num - nodes[index-1].num
		currCost = currCost - (sumCost-prevSumCost)*diff + prevSumCost*diff
		prevSumCost += nodes[index].cost
		if currCost < minCost {
			minCost = currCost
		}
	}
	return minCost
}
