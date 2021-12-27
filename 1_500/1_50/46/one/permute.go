package one

func Permute(nums []int) [][]int {
	var rs [][]int
	bc(0, nums, nil, &rs)
	return rs
}

func bc(i int, nums []int, cn []int, rs *[][]int) {
	// 结尾
	if i == len(nums) {
		*rs = append(*rs, append([]int{}, cn...))
		return
	}
	// 依次取出数字，并且递归交换位置继续进行组合
	for j := i; j < len(nums); j++ {
		swap(nums, i, j)
		bc(i+1, nums, append(cn, nums[i]), rs)
		swap(nums, j, i)
	}
}

func swap(nums []int, i, j int) {
	if i == j {
		return
	}
	nums[j], nums[i] = nums[i], nums[j]
}
