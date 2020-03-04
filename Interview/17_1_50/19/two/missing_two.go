package two

// 转化为只出现一次的两个数字
func MissingTwo(nums []int) []int {
	num := 0
	for i := range nums {
		num ^= nums[i]
	}
	for i := len(nums) + 2; i > 0; i-- {
		num ^= i
	}
	// 取出最后不同一位
	num &= -num
	ans := make([]int, 2)
	for i := range nums {
		if nums[i]&num != 0 {
			ans[0] ^= nums[i]
		} else {
			ans[1] ^= nums[i]
		}
	}
	for i := len(nums) + 2; i > 0; i-- {
		if i&num != 0 {
			ans[0] ^= i
		} else {
			ans[1] ^= i
		}
	}
	return ans
}
