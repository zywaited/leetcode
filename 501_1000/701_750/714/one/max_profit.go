package one

func MaxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	prev := -prices[0] // dp[i-1][0]
	for _, price := range prices {
		prev = max(prev, ans-price)
		ans = max(ans, prev+price-fee)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
