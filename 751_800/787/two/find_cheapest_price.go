package two

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// dp[i][j] 经过i站到达j处的最小花费
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][src] = 0
	ans := -1
	cost := 0
	for i := 0; i <= k; i++ {
		for _, flight := range flights {
			if dp[i][flight[0]] == -1 {
				continue
			}
			cost = dp[i][flight[0]] + flight[2]
			if flight[1] == dst && (ans == -1 || cost < ans) {
				ans = cost
			}
			if i == k {
				continue
			}
			if dp[i+1][flight[1]] == -1 || cost < dp[i+1][flight[1]] {
				dp[i+1][flight[1]] = cost
			}
		}
	}
	return ans
}
