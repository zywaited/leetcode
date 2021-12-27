package one

func CountSubstrings(s string) int {
	ans := 0
	// 区间i-j是否是回文串
	dp := make([][]bool, len(s))
	for j := 0; j < len(s); j++ {
		dp[j] = make([]bool, len(s))
		for i := j; i >= 0; i-- {
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}
