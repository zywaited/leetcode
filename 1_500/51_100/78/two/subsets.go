package two

// 该方法需要保证数量不超过int的长度
func Subsets(nums []int) [][]int {
	// 子集一共2^N个
	n := 1 << uint(len(nums))
	rs := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		sub := make([]int, 0, len(nums))
		for j := 0; j < len(nums); j++ {
			// 判断当前数字是否应该在当前位中
			if i&(1<<uint(j)) != 0 {
				sub = append(sub, nums[j])
			}
		}
		rs = append(rs, sub)
	}
	return rs
}
