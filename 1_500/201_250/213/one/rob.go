package one

func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 就是需要计算两次
	// 计算nums[0]和不计算nums[0]
	f, s := dp(nums, 0, len(nums)-1), dp(nums, 1, len(nums))
	if f < s {
		return s
	}
	return f
}

func dp(nums []int, l, r int) int {
	f := 0
	s := 0
	t := 0
	for i := l; i < r; i++ {
		t = f
		if s+nums[i] > f {
			f = s + nums[i]
		}
		s = t
	}
	return f
}
