package one

import "math"

func MaxSubArray(nums []int) int {
	// dp[n] = max(nums[n], dp[n+1]+nums[n])
	dp := make([]int, len(nums)+1)
	m := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(nums[i], dp[i+1]+nums[i])
		if dp[i] > m {
			m = dp[i]
		}
	}
	return m
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
