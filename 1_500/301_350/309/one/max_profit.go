package one

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	p := -prices[0] // dp[i-1][0]
	pp := 0         // dp[i-2][1]
	for _, price := range prices {
		tp := ans
		p = max(p, pp-price)
		ans = max(ans, p+price)
		pp = tp // 保留上次的数据
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
