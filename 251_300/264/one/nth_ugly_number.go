package one

func NthUglyNumber(n int) int {
	// 动态规划(2,3,5相乘的数一定是丑数)
	dp := make([]int, n+1)
	dp[1] = 1
	// 2,3,5当前索引
	i2 := 1
	i3 := 1
	i5 := 1
	for num := 2; num <= n; num++ {
		// 那个小就在前面
		dp[num] = min(dp[i2]*2, min(dp[i3]*3, dp[i5]*5))
		// 处理索引
		// 这里判断逻辑是并列是因为存在刚好都相等的情况
		if dp[num] == dp[i2]*2 {
			i2++
		}
		if dp[num] == dp[i3]*3 {
			i3++
		}
		if dp[num] == dp[i5]*5 {
			i5++
		}
	}
	return dp[n]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
