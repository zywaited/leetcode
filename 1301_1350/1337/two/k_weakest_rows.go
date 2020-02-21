package one

import "sort"

// 如果数据量大时考虑性能+结构巩固, 这里采用二分法+最大堆
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
	// 初始化数量为K的最大堆
	// 调整堆
	var heapify func(int)
	heapify = func(i int) {
		l := (i << 1) + 1
		r := (i << 1) + 2
		// 3者找到最大的
		max := i
		if l < k && ((nodes[l].num > nodes[max].num) || (nodes[l].num == nodes[max].num && nodes[l].index > nodes[max].index)) {
			max = l
		}
		if r < k && ((nodes[r].num > nodes[max].num) || (nodes[r].num == nodes[max].num && nodes[r].index > nodes[max].index)) {
			max = r
		}
		// 交换并继续往下
		if max != i {
			nodes[i], nodes[max] = nodes[max], nodes[i]
			heapify(max)
		}
	}
	for i := (k >> 1) - 1; i >= 0; i-- {
		heapify(i)
	}
	// 依次遍历与最小堆比较
	for i := k; i < len(nodes); i++ {
		if (nodes[i].num < nodes[0].num) || (nodes[i].num == nodes[0].num && nodes[i].index < nodes[0].index) {
			nodes[0] = nodes[i]
			heapify(0) // 重新调整最大堆
		}
	}
	// 最终结果
	// 这里只对K个数排序
	sort.Slice(nodes[:k], func(i, j int) bool {
		return (nodes[i].num < nodes[j].num) || (nodes[i].num == nodes[j].num && nodes[i].index < nodes[j].index)
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = nodes[i].index
	}
	return ans
}
