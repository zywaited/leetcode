package one

import "sort"

func BestSeqAtIndex(height []int, weight []int) int {
	// 题目没有说人数大于0
	if len(height) <= 1 {
		return len(height)
	}
	people := make([][2]int, len(height))
	for i := range height {
		people[i][0] = height[i]
		people[i][1] = weight[i]
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1] // 防止重复使用相同身高的
	})
	dp := make([]int, 0, len(people))
	dp = append(dp, people[0][1]) // 用weight为依据，height已经排过序
	for i := 1; i < len(people); i++ {
		// 可以满足
		if people[i][1] < dp[len(dp)-1] {
			dp = append(dp, people[i][1])
			continue
		}
		if people[i][1] == dp[len(dp)-1] {
			continue
		}
		// 二分查找
		j := 0
		s := 0
		e := len(dp) - 1
		for s <= e {
			m := s + (e-s)>>1
			// 想等就不需要任何操作了
			if dp[m] == people[i][1] {
				j = m
				break
			}
			// 找到小的了
			if dp[m] < people[i][1] {
				j = m
				e = m - 1
				continue
			}
			s = m + 1
		}
		dp[j] = people[i][1]
	}
	return len(dp)
}
