package one

import "math"

func StoneGameIII(stoneValue []int) string {
	// dp[i][j] 剩余i堆石子，取走前j堆石子的最大值
	dp := make([][4]int, len(stoneValue)+1)
	// 剩余石子总分
	sum := 0
	for i := 1; i <= len(stoneValue); i++ {
		sum += stoneValue[len(stoneValue)-i]
		// 1-3的最大值
		ms := math.MinInt32
		for j := 1; j <= 3; j++ {
			dp[i][j] = math.MinInt32
			if j >= i {
				// 全取
				dp[i][j] = sum
			} else {
				dp[i][j] = max(dp[i][j], sum-dp[i-j][0])
			}
			ms = max(ms, dp[i][j])
		}
		dp[i][0] = ms
	}
	if dp[len(stoneValue)][0] > sum-dp[len(stoneValue)][0] {
		return "Alice"
	}
	if dp[len(stoneValue)][0] < sum-dp[len(stoneValue)][0] {
		return "Bob"
	}
	return "Tie"
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
