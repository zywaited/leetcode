package one

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// dp[i][j] s1中i个字符+s2中j个字符是否可以构建为s3的i+j
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				// 全部由s2构建
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1]
				continue
			}
			if j == 0 {
				// 全部由s1构建
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1]
				continue
			}
			// 这时候有两种情况
			// 最后一个字符由s1构建
			// 最后一个字符由s2构建
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}
