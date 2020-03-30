package one

func MinDistance(word1 string, word2 string) int {
	// 全是增或者减
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([][]int, len(word1)+1)
	dp[0] = make([]int, len(word2)+1)
	for i := 1; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i // 只能删除
		for j := 1; j <= len(word2); j++ {
			dp[0][j] = j // 只能插入(这里会重复赋值, 不过不用多循环一次)
			// 3中操作选择最小的
			// 插入
			dp[i][j] = dp[i][j-1] + 1
			// 删除
			if dp[i-1][j]+1 < dp[i][j] {
				dp[i][j] = dp[i-1][j] + 1
			}
			// 替换(如果相等就不用替换)
			if word1[i-1] == word2[j-1] {
				if dp[i-1][j-1] < dp[i][j] {
					dp[i][j] = dp[i-1][j-1]
				}
				continue
			}
			if dp[i-1][j-1]+1 < dp[i][j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
