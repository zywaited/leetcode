package one

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	type node struct {
		num   int
		index int
	}
	ns := make([]node, 0, len(nums))
	for index, num := range nums {
		ns = append(ns, node{num: num, index: index})
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].num < ns[j].num
	})
	ans := make([]int, len(nums))
	nn := make([]int, 101)
	for index, n := range ns {
		ans[n.index] = index - nn[n.num]
		nn[n.num]++
	}
	return ans
}
