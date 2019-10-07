package one

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	rs := make([][]int, 0)
	bc(nums, 0, nil, &rs)
	rs = append(rs, []int{}) // 空集合
	return rs
}

func bc(nums []int, start int, cns []int, rs *[][]int) {
	for index := start; index < len(nums); index++ {
		// 不是开始数字并且和上一次相等就跳过
		if index == start || nums[index] != nums[index-1] {
			cns = append(cns, nums[index])
			*rs = append(*rs, append([]int{}, cns...))
			bc(nums, index+1, cns, rs)
			cns = cns[:len(cns)-1]
		}
	}
}
