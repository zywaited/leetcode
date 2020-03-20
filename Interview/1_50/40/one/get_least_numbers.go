package one

import "sort"

func GetLeastNumbers(arr []int, k int) []int {
	// 排序
	sort.Ints(arr)
	return arr[:k]
}
