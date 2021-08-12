package one

func LongestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for j := 0; j <= len(s); j++ {
		dp[j] = make([]int, len(s)+1)
		if j == 0 {
			continue
		}
		dp[j][j] = 1
		for i := j - 1; i > 0; i-- {
			dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1]))
			if s[i-1] == s[j-1] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			}
		}
	}
	return dp[1][len(s)]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
