package one

import (
	"fmt"
	"sort"
)

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	coinIndex := 0
	sum := 0
	addNum := 0
	for sum < target {
		if coinIndex < len(coins) {
			if coins[coinIndex] <= sum || sum+1 == coins[coinIndex] {
				sum += coins[coinIndex]
				coinIndex++
				continue
			}
		}
		addNum++
		sum += sum + 1
	}
	return addNum
}
