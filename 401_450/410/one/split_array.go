package one

func SplitArray(nums []int, m int) int {
	// 前缀和
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	// dp[i][j] 区间1-i的分割为j连续数组的最大值
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = prefixSum[len(nums)]
		}
	}
	dp[0][0] = 0

	for i := 1; i <= len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			for k := 1; k <= m; k++ {
				if i < k || j < k-1 {
					break
				}
				dp[i][k] = min(dp[i][k], max(dp[j][k-1], prefixSum[i]-prefixSum[j]))
			}
		}
	}

	return dp[len(nums)][m]
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
