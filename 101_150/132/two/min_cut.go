package two

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
	// 第一种解法是计算了所有区间的最小值
	// 这次压缩区间和空间，就只算0-j区间即可
	// 这里能压缩区间是因为当[i][j]是回文串的是个整体，不用管dp[i][j]的变化
	dp := make([]int, len(s))
	for j := 1; j < len(s); j++ {
		// dp[i]的最大值就是i
		dp[j] = j
		if pp[0][j] {
			dp[j] = 0
			continue
		}
		for i := j; i >= 0; i-- {
			if !pp[i][j] {
				continue
			}
			dp[j] = min(dp[j], dp[i-1]+1)
		}
	}
	return dp[len(s)-1]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
