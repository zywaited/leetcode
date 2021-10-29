package two

func KInversePairs(n int, k int) int {
	mod := 1000000007
	dp := make([]int, k+1)
	dp[0] = 1
	sum := make([]int, k+1)
	for i := 1; i <= n; i++ {
		sum[0] = dp[0]
		for j := 1; j <= k; j++ {
			sum[j] = (sum[j-1] + dp[j]) % mod
		}

		for j := min(i*(i-1)/2, k); j >= 0; j-- {
			if j <= i-1 {
				dp[j] = sum[j]
				continue
			}
			dp[j] = (sum[j] - sum[j-(i-1)-1] + mod) % mod

		}
	}
	return dp[k]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
