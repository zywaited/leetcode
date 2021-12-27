package one

func MaxProfit(prices []int) int {
	// 多次购买就是121的改编版
	// 只要后一天比前一天大，就可以买卖
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
