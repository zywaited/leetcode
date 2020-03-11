package one

func CuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := i >> 1; j > 0; j-- {
			dp[i] = max(dp[i], j*(i-j))
			dp[i] = max(dp[i], j*dp[i-j])
			dp[i] = max(dp[i], (i-j)*dp[j])
		}
	}
	return dp[n]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
