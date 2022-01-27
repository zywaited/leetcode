package one

func MinimumTime(n int, relations [][]int, time []int) int {
	costs := make([]int, n)
	prevs := make([][]int, n)
	for _, relation := range relations {
		prevs[relation[1]-1] = append(prevs[relation[1]-1], relation[0]-1)
	}
	var dfs func(course int) int
	dfs = func(course int) int {
		if costs[course] == 0 {
			for _, nextCourse := range prevs[course] {
				costs[course] = max(costs[course], dfs(nextCourse))
			}
			costs[course] += time[course]
		}
		return costs[course]
	}
	ans := 0
	for course := 0; course < n; course++ {
		if costs[course] == 0 {
			ans = max(ans, dfs(course))
		}
	}
	return ans
}

func max(f, s int) int {
	if s <= f {
		return f
	}
	return s
}
