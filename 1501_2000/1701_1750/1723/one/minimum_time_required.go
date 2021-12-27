package one

import (
	"math"
	"math/bits"
)

func MinimumTimeRequired(jobs []int, k int) int {
	n := 1 << uint(len(jobs))
	sum := make([]int, n)
	for i := 1; i < n; i++ {
		// note: 系统有对应的方法
		zn := bits.TrailingZeros(uint(i))
		sum[i] = sum[i^(1<<uint(zn))] + jobs[zn]
	}
	dp := make([][]int, k)
	for tk := 0; tk < k; tk++ {
		dp[tk] = make([]int, n)
		if tk == 0 {
			for i, m := range sum {
				dp[tk][i] = m
			}
			continue
		}
		for i := 1; i < n; i++ {
			m := math.MaxInt32
			// 各种组合逻辑
			for c := i; c > 0; c = (c - 1) & i {
				m = min(m, max(dp[tk-1][i-c], sum[c]))
			}
			dp[tk][i] = m
		}
	}
	return dp[k-1][n-1]
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
