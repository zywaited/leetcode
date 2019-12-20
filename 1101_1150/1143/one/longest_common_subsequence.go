package one

// 动态规划
// dp[i][j] 代表 text1前i个字符与text2前j个字符的最大子序列长度(因为能这样计算是因为不需要连续)
func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
