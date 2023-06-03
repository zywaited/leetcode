package one

import "sort"

func MaximumTastiness(prices []int, k int) int {
	sort.Ints(prices)
	left := 0
	right := (prices[len(prices)-1] - prices[0]) / (k - 1)
	for left < right {
		mid := left + (right-left+1)/2
		tk := 1
		price := prices[0]
		for index := 1; index < len(prices) && tk < k; index++ {
			if mid <= prices[index]-price {
				tk++
				price = prices[index]
			}
		}
		if tk < k {
			right = mid - 1
			continue
		}
		left = mid
	}
	return left
}
