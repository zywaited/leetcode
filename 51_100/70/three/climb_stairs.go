package three

// 循环
// f(n) = f(n-1) + f(n-2)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 数组和map都可以
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for index := 2; index < n; index++ {
		dp[index] = dp[index-1] + dp[index-2]
	}
	return dp[n-1]
}
