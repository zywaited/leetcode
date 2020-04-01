package three

func Partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
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
	// 压缩第二种解法的空间和区间(0-j)
	// 这里能压缩区间是因为当[i][j]是回文串的是个整体，不用管dp[i][j]的变化
	dp := make([][][]string, len(s))
	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			if !pp[i][j] {
				continue
			}
			if i == 0 {
				dp[j] = append(dp[j], []string{s[:j+1]})
				continue
			}
			for _, d := range dp[i-1] {
				dp[j] = append(dp[j], append(append([]string{}, d...), s[i:j+1]))
			}
		}
	}
	return dp[len(s)-1]
}
