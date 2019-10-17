package one

import "sort"

func NumMovesStones(a int, b int, c int) []int {
	pointers := []int{a, b, c}
	sort.Ints(pointers)
	// 3个点之间移动距离
	ab := pointers[1] - pointers[0] - 1
	bc := pointers[2] - pointers[1] - 1
	rs := make([]int, 2)
	// 最小值就是每个最大一次移动
	if ab > 0 {
		rs[0]++
	}
	if bc > 0 {
		rs[0]++
	}
	if (ab == 1 || bc == 1) && rs[0] > 1 {
		// 只需移动1次
		rs[0] = 1
	}
	// 最大值就是都移动到一边
	rs[1] = ab + bc
	return rs
}
