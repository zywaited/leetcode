package one

func StoneGameVII(stones []int) int {
	sums := make([]int, len(stones)+1)
	for i := range stones {
		sums[i+1] = sums[i] + stones[i]
	}
	// dp[i][j] 区域i-j的差值
	dp := make([][]int, len(stones)+1)
	for j := 1; j <= len(stones); j++ {
		dp[j] = make([]int, len(stones)+1)
		for i := j - 1; i > 0; i-- {
			dp[i][j] = max(sums[j]-sums[i]-dp[i+1][j], sums[j-1]-sums[i-1]-dp[i][j-1])
		}
	}
	return dp[1][len(stones)]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
