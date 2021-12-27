package one

func StoneGameV(stoneValue []int) int {
	// 区间对应的分数
	sum := make([]int, len(stoneValue)+1)
	for index := range stoneValue {
		sum[index+1] = stoneValue[index] + sum[index]
	}

	// dp[i][j] 区间i-j的最大分数
	dp := make([][]int, len(stoneValue)+1)
	for j := 1; j <= len(stoneValue); j++ {
		dp[j] = make([]int, len(stoneValue)+1)
		for i := j - 1; i > 0; i-- {
			for k := i; k < j; k++ {
				// 以k为分割线
				f, s := sum[k]-sum[i-1], sum[j]-sum[k]
				if f > s {
					dp[i][j] = max(dp[i][j], s+dp[k+1][j])
					continue
				}
				if f < s {
					dp[i][j] = max(dp[i][j], f+dp[i][k])
					continue
				}
				dp[i][j] = max(dp[i][j], max(f, s)+max(dp[i][k], dp[k+1][j]))
			}
		}
	}

	return dp[1][len(stoneValue)]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
