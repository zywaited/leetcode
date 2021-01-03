package one

import (
	"math"
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := 0
	max := math.MinInt32
	for _, interval := range intervals {
		if interval[0] >= max {
			max = interval[1]
			continue
		}
		ans++
		if interval[1] < max {
			max = interval[1]
		}
	}
	return ans
}
