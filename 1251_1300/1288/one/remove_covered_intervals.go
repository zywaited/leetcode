package one

import "sort"

func RemoveCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	num := 1
	min := intervals[0][0]
	max := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == min || intervals[i][1] <= max {
			continue
		}
		max = intervals[i][1]
		num++
	}
	return num
}
