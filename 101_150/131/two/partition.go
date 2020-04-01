package two

func Partition(s string) [][]string {
	ans := make([][]string, 0)
	if len(s) == 0 {
		ans = append(ans, []string{})
		return ans
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
	// 动态规划i-j的可能性, 但也可能会导致保存数据过大(分治)
	// 测试能通用, 内存占用了120M
	// dp[i][j] = dp[k+1][j] * s[i:k+1]  (如果是回文串) [i <= k <= j]
	dp := make([][][][]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([][][]string, len(s))
		for j := i; j < len(s); j++ {
			for k := i; k <= j; k++ {
				if !pp[i][k] {
					continue
				}
				if k+1 > j { // i-j整体是回文串
					dp[i][j] = append(dp[i][j], []string{s[i : k+1]})
					continue
				}
				// 性能相比回溯慢就是这里频繁的进行大量的拷贝
				for _, d := range dp[k+1][j] {
					dp[i][j] = append(dp[i][j], append([]string{s[i : k+1]}, d...))
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
