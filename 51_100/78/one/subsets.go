package one

func Subsets(nums []int) [][]int {
	rs := make([][]int, 0)
	bc(nums, 0, nil, &rs)
	rs = append(rs, []int{}) // 空集合
	return rs
}

func bc(nums []int, start int, cns []int, rs *[][]int) {
	for index := start; index < len(nums); index++ {
		cns = append(cns, nums[index])
		*rs = append(*rs, append([]int{}, cns...))
		bc(nums, index+1, cns, rs)
		cns = cns[:len(cns)-1]
	}
}
