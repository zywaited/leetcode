package one

import "sort"

func DistributeCandies(candies []int) int {
	sort.Ints(candies)
	half := len(candies) >> 1
	ans := 1
	for i := 1; i < len(candies) && ans < half; i++ {
		if candies[i] > candies[i-1] {
			ans++
		}
	}
	return ans
}
