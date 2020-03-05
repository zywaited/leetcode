package one

var scores = [4]int{1, 5, 10, 25}

const mod = 1e9 + 7

func WaysToChange(n int) int {
	// 记录每个值对应的方案数量
	dp := make([]int, n+1)
	dp[0] = 1
	for _, score := range scores {
		for i := score; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-score]) % mod
		}
	}
	return dp[n]
}
