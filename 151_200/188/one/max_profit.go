package one

import "math"

func MaxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	// 如果k > len(prices)>>1, 那么就是无限次交易
	// 基于前面几道题的做法
	if k > len(prices)>>1 || k == 1 {
		ans := 0
		bp := -prices[0] // 买入
		tp := 0
		for _, price := range prices {
			tp = ans
			if k == 1 {
				tp = 0
			}
			bp = max(bp, tp-price)
			ans = max(ans, bp+price)
		}
		return max(ans, 0)
	}
	// dp[k][j] k: 交易次数, j: 买卖
	dp := make([][]int, k+1)
	ck := 0               // 临时存储k的值
	for ; ck <= k; ck++ { // 初始化数据
		dp[ck] = make([]int, 2)
		dp[ck][0] = math.MinInt32
		dp[ck][1] = math.MinInt32
	}
	dp[0][1] = 0 // 0次交易数据都是0
	dp[1][0] = -prices[0]
	for _, price := range prices {
		for ck = 1; ck <= k; ck++ {
			dp[ck][0] = max(dp[ck][0], dp[ck-1][1]-price) // 第ck次买入
			dp[ck][1] = max(dp[ck][1], dp[ck][0]+price)   // 第ck次卖出
		}
	}
	return max(dp[k][1], 0)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
