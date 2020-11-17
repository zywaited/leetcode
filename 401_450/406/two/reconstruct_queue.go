package two

import "sort"

func ReconstructQueue(people [][]int) [][]int {
	// 先按照身高排序
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	// k即是对应的索引
	var tmp []int
	for i := range people {
		tmp = people[i]
		// 从大到小排序，移动k-i区间即可
		for j := i; j > tmp[1]; j-- {
			people[j] = people[j-1]
		}
		people[tmp[1]] = tmp
	}
	return people
}
