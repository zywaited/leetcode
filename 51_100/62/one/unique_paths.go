package one

// f(m, n) = f(m, n-1) + f(m-1, n)
func UniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	// 边缘路线只有1条
	for i := 0; i < m; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, n)
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 逐一计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
