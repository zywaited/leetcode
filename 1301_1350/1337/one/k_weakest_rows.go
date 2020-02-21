package one

import "sort"

// 这里采用二分法
func KWeakestRows(mat [][]int, k int) []int {
	type node struct {
		num   int
		index int
	}
	nodes := make([]node, len(mat))
	for i := range mat {
		// 二分获取1的个数
		s, e := 0, len(mat[i])-1
		for s <= e {
			m := s + (e-s)>>1
			if mat[i][m] == 0 {
				e = m - 1
				continue
			}
			s = m + 1
		}
		nodes[i].index = i
		nodes[i].num = s
	}
	sort.Slice(nodes, func(i, j int) bool {
		return (nodes[i].num < nodes[j].num) || (nodes[i].num == nodes[j].num && nodes[i].index < nodes[j].index)
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = nodes[i].index
	}
	return ans
}
