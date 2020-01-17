package one

func FindOrder(numCourses int, prerequisites [][]int) []int {
	ans := make([]int, 0, numCourses)
	mp := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		if mp[prerequisite[0]] == nil {
			mp[prerequisite[0]] = make([]int, 0, numCourses)
		}
		mp[prerequisite[0]] = append(mp[prerequisite[0]], prerequisite[1])
	}
	// 记录当前正在遍历的节点
	used := make([]int8, numCourses)
	var dfs func(int) int
	dfs = func(i int) int {
		// 检查过了
		if used[i] == -1 {
			return 0
		}
		// 已经在使用(循环依赖了)
		if used[i] == 1 {
			return -1
		}
		used[i] = 1
		for _, j := range mp[i] {
			if dfs(j) < 0 {
				return -1
			}
		}
		used[i] = -1
		ans = append(ans, i)
		return i
	}
	// 检查有依赖的节点
	for i := range mp {
		if dfs(i) < 0 {
			return nil
		}
	}
	return ans
}
