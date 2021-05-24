package one

func StrangePrinter(s string) int {
	// i-j 所需的最少次数
	dp := make([][]int, len(s))
	m := -1
	for j := range s {
		dp[j] = make([]int, len(s))
		dp[j][j] = 1
		for i := j - 1; i >= 0; i-- {
			m = -1
			for k := j - 1; k >= i; k-- {
				if m == -1 || dp[i][k]+dp[k+1][j] < m {
					m = dp[i][k] + dp[k+1][j]
				}
				if s[k] == s[j] && dp[i][k]+dp[k+1][j-1] < m {
					m = dp[i][k] + dp[k+1][j-1]
				}
			}
			dp[i][j] = m
		}
	}
	return dp[0][len(s)-1]
}
