package one

import (
	"math"
	"sort"
)

// 1: d1 < d2 那么选1这门课是最优
// 2: d1 == d2 && t1 <= t2 那么选1这门课是最优
func ScheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})
	num := 0
	// 实际消耗最优时间
	ts := make([]int, len(courses)+1)
	for i, td := range courses {
		ts[i+1] = math.MaxInt32
		for j := i; j >= 0; j-- {
			t := td[0] + ts[j]
			if t <= td[1] && t < ts[j+1] {
				ts[j+1] = t
				if j+1 > num {
					num = j + 1
				}
			}
		}
	}
	return num
}
