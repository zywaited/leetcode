package two

const mod = 1e9 + 7

func CountVowelPermutation(n int) int {
	// 初始化数据(DP代表从前往后的方向)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 5)
		if i > 0 {
			continue
		}
		// 第一位都是1
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	for i := 1; i < n; i++ {
		// a前面是e, i, u
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % mod
		// e前面是a, i
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		// i前面是e, o
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % mod
		// o前面是i
		dp[i][3] = dp[i-1][2] % mod
		// u前面是i, o
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % mod
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum = (sum + dp[n-1][i]) % mod
	}
	return sum
}
