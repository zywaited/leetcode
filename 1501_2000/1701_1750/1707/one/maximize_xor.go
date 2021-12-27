package one

import "sort"

type node [2]*node

func (n *node) insert(num int) {
	cn := n
	j := 0
	for i := 31; i >= 0; i-- {
		j = num >> uint(i) & 1
		if cn[j] == nil {
			cn[j] = &node{}
		}
		cn = cn[j]
	}
}

func (n *node) find(num int) int {
	cn := n
	j := 0
	m := 0
	for i := 31; i >= 0; i-- {
		j = num >> uint(i) & 1
		if cn[j^1] != nil {
			cn = cn[j^1]
			m |= 1 << uint(i)
			continue
		}
		cn = cn[j]
	}
	return m
}

type queryNode struct {
	index int
	query []int
}

func MaximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	queryNodes := []queryNode{}
	for index := range queries {
		queryNodes = append(queryNodes, queryNode{index: index, query: queries[index]})
	}
	sort.Slice(queryNodes, func(i, j int) bool {
		return queryNodes[i].query[1] < queryNodes[j].query[1]
	})
	ans := make([]int, len(queries))
	nodes := node{}
	i := 0
	for _, qn := range queryNodes {
		for ; i < len(nums) && qn.query[1] >= nums[i]; i++ {
			(&nodes).insert(nums[i])
		}
		if i == 0 {
			ans[qn.index] = -1
			continue
		}
		ans[qn.index] = (&nodes).find(qn.query[0])
	}
	return ans
}
