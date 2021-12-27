package one

const mod = 1e9 + 7

func NumOfArrays(n int, m int, k int) int {
	// dp[n][m][k] 代表长度为n，当前最大值为m，search_cost为k的数量
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for c := 1; c <= k; c++ {
				if i == 1 && c == 1 {
					dp[i][j][c] = 1
					continue
				}
				// 如果i-1长度已经到c了，那么最后一位只要不比j大就行(1-j)
				dp[i][j][c] = (dp[i-1][j][c] * j) % mod
				// 如果i-1长度到c-1, 那么遍历所有比j小的，最后加上j就行了
				for pj := 1; pj < j; pj++ {
					dp[i][j][c] = (dp[i][j][c] + dp[i-1][pj][c-1]) % mod
				}
			}
		}
	}
	ans := 0
	for j := 1; j <= m; j++ {
		ans = (ans + dp[n][j][k]) % mod
	}
	return ans
}
