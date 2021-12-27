package one

func CoinChange(coins []int, amount int) int {
	// 0不参与计算
	if amount == 0 {
		return 0
	}
	// 背包问题
	dp := make([]int, amount+1)
	for _, coin := range coins {
		for k := coin; k <= amount; k++ {
			if k == coin || (dp[k-coin] > 0 && (dp[k] == 0 || dp[k-coin]+1 < dp[k])) {
				dp[k] = dp[k-coin] + 1
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}
