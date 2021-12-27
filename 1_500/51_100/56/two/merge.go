package two

import (
	"sort"
)

// 设置排序规则
type slice [][]int

func (s slice) Len() int {
	return len(s)
}

func (s slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s slice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func Merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size < 2 {
		return intervals
	}

	// 起点从小到大
	sort.Sort(slice(intervals))
	var r [][]int
	for i := 1; i < size; i++ {
		if intervals[i-1][1] < intervals[i][0] {
			r = append(r, intervals[i-1])
			continue
		}

		intervals[i][0] = intervals[i-1][0]
		if intervals[i-1][1] > intervals[i][1] {
			intervals[i][1] = intervals[i-1][1]
		}
	}

	// 添加最后一个
	r = append(r, intervals[size-1])
	return r
}
