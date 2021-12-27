package one

import "sort"

func DeleteAndEarn(nums []int) int {
	type node struct {
		num   int
		score int
	}
	nodes := []node{}
	nm := map[int]int{}
	for _, num := range nums {
		index := nm[num]
		if index == 0 {
			nodes = append(nodes, node{num: num, score: num})
			nm[num] = len(nodes)
			continue
		}
		nodes[index-1].score += num
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].num < nodes[j].num
	})
	ps := 0
	cs := 0
	for i := 0; i < len(nodes); i++ {
		if i == 0 {
			cs = nodes[i].score
			continue
		}
		if nodes[i].num-1 == nodes[i-1].num {
			ps, cs = cs, max(ps+nodes[i].score, cs)
			continue
		}
		ps, cs = cs, max(ps, cs)+nodes[i].score
	}
	return max(ps, cs)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
