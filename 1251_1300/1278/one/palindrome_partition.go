package one

func PalindromePartition(s string, k int) int {
	if len(s) == k {
		return 0
	}
	// 动态规划i-j变成回文串需要的次数
	pp := make([][]int, len(s)+1)
	for i := len(s); i > 0; i-- {
		pp[i] = make([]int, len(s)+1)
		for j := i; j <= len(s); j++ {
			if j-i <= 2 {
				pp[i][j] = 0
				if s[i-1] != s[j-1] {
					pp[i][j] = 1
				}
				continue
			}
			pp[i][j] = pp[i+1][j-1]
			if s[i-1] != s[j-1] {
				pp[i][j]++
			}
		}
	}
	// 0-i区间切割成j个回文串所需修改的次数
	dp := make([][]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		dp[i] = make([]int, k+1)
		for ck := 0; ck <= k; ck++ {
			dp[i][ck] = len(s) + 1 // 切割次数置为最大
			if ck == 0 || ck > i {
				continue
			}
			if ck == i {
				dp[i][ck] = 0
				continue
			}
			if ck == 1 {
				dp[i][ck] = pp[1][i]
				continue
			}
			for j := i; j > 1; j-- {
				dp[i][ck] = min(dp[i][ck], dp[j-1][ck-1]+pp[j][i])
			}
		}
	}
	return dp[len(s)][k]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
