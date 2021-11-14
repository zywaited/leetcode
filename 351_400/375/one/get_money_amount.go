package one

func GetMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = make([]int, n+1)
		if j <= 1 {
			continue
		}
		for i := j - 1; i > 0; i-- {
			dp[i][j] = min(dp[i][j-1]+j, dp[i+1][j]+i)
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
			}
		}
	}
	return dp[1][n]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
