package two

func MinDistance(word1 string, word2 string) int {
	// 全是增或者减
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([]int, len(word2)+1)
	// 当i==0时，对应的距离
	for j := 1; j <= len(word2); j++ {
		dp[j] = j
	}
	// pi=[i-1][j] pj=[i-1][j-1]
	pi, pj := 0, 0
	for i := 1; i <= len(word1); i++ {
		// [i-1][0]
		pi = dp[0]
		// i--0距离
		dp[0] = i
		for j := 1; j <= len(word2); j++ {
			// [i-1][j-1]
			pj = pi
			// [i-1][j]
			pi = dp[j]
			// 删除
			dp[j] = pi + 1
			// 插入dp[j-1]=[i][j-1]
			if dp[j-1]+1 < dp[j] {
				dp[j] = dp[j-1] + 1
			}
			if word1[i-1] == word2[j-1] {
				if pj < dp[j] {
					dp[j] = pj
				}
				continue
			}
			if pj+1 < dp[j] {
				dp[j] = pj + 1
			}
		}
	}
	return dp[len(word2)]
}
