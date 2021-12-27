package one

import "sort"

func ReconstructQueue(people [][]int) [][]int {
	// 先按照身高排序
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] < people[j][0]) || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	ans := make([][]int, len(people))
	num := 0
	for _, p := range people {
		num = p[1] + 1
		for i := range ans {
			if len(ans[i]) == 0 {
				num--
				if num == 0 {
					ans[i] = p
					break
				}
			}
		}
	}
	return ans
}
