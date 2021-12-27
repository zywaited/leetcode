package one

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化
	mp := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		if mp[prerequisite[0]] == nil {
			mp[prerequisite[0]] = make([]int, 0, numCourses)
		}
		mp[prerequisite[0]] = append(mp[prerequisite[0]], prerequisite[1])
	}
	// 记录当前正在遍历的节点
	used := make([]int8, numCourses)
	var dfs func(int) bool
	dfs = func(i int) bool {
		// 没有节点或者检查过了
		if len(mp[i]) == 0 || used[i] == -1 {
			return true
		}
		// 已经在使用(循环依赖了)
		if used[i] == 1 {
			return false
		}
		used[i] = 1
		for _, j := range mp[i] {
			if !dfs(j) {
				return false
			}
		}
		used[i] = -1
		return true
	}
	// 检查有依赖的节点
	for i := range mp {
		if !dfs(i) {
			return false
		}
	}
	return true
}
