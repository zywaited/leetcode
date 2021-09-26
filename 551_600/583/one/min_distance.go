package one

func MinDistance(word1 string, word2 string) int {
	num := 0
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= len(word2); j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
			num = max(num, dp[i][j])
		}
	}
	return len(word1) + len(word2) - num*2
}

func max(f, s int) int {
	if f <= s {
		return s
	}
	return f
}
