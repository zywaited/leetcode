package one

import "sort"

func FindRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	hi := 0
	cl := 0
	for _, house := range houses {
		hi = sort.SearchInts(heaters, house)
		if hi == len(heaters) {
			ans = max(ans, house-heaters[hi-1])
			continue
		}
		cl = heaters[hi] - house
		if 1 <= hi {
			cl = min(cl, house-heaters[hi-1])
		}
		ans = max(ans, cl)
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
