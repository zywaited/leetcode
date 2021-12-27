package one

import "math"

func MaxProfit(prices []int) int {
	// 某一索引的做大值就是左边交易最大值+右边交易最大值
	lm := make([]int, len(prices)+1)
	rm := make([]int, len(prices)+1)
	index := 0
	// 计算左边所有的最大值
	pn := math.MaxInt32
	for index = 1; index <= len(prices); index++ {
		lm[index] = max(lm[index-1], prices[index-1]-pn)
		pn = min(pn, prices[index-1])
	}
	// 计算右边的最大值
	pn = 0
	for index = len(prices) - 1; index >= 0; index-- {
		rm[index] = max(rm[index+1], pn-prices[index])
		pn = max(pn, prices[index])
	}
	ans := 0
	for index = range prices {
		ans = max(ans, lm[index]+rm[index])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
