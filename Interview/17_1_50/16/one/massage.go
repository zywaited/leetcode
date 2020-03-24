package one

func Massage(nums []int) int {
	last := 0 // 上一次的最优值
	ans := 0  // 本次的最优值
	for _, num := range nums {
		last, ans = ans, max(ans, num+last)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
