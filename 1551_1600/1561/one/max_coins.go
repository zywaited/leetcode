package one

import "sort"

func MaxCoins(piles []int) int {
	sort.Ints(piles)
	ans := 0
	num := len(piles) / 3
	index := num
	for ; num > 0; num-- {
		ans += piles[index]
		index += 2
	}
	return ans
}
