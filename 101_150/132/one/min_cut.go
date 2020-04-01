package one

import "math"

func MinCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	// 动态规划i-j是否是回文串
	pp := make([][]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		pp[i] = make([]bool, len(s))
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || pp[i+1][j-1]) {
				pp[i][j] = true
			}
		}
	}
	// 动态规划次数
	// 计算i-j的次数
	dp := make([][]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(s))
		for j := i; j < len(s); j++ {
			if pp[i][j] {
				// 不用切割
				continue
			}
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				if !pp[i][k] {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[k+1][j]+1)
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
