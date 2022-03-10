package one

import "sort"

func SortJumbled(mapping []int, nums []int) []int {
	type node struct {
		index   int
		mapping int
	}
	nodes := make([]node, len(nums))
	for index, num := range nums {
		nodes[index].index = index
		if num == 0 {
			nodes[index].mapping = mapping[num]
			continue
		}
		multi := 1
		for ; num > 0; num /= 10 {
			nodes[index].mapping += mapping[num%10] * multi
			multi *= 10
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].mapping == nodes[j].mapping {
			return nodes[i].index < nodes[j].index
		}
		return nodes[i].mapping < nodes[j].mapping
	})
	ans := make([]int, len(nums))
	for index, cn := range nodes {
		ans[index] = nums[cn.index]
	}
	return ans
}
