package one

func EventualSafeNodes(graph [][]int) []int {
	safe := make(map[int]bool, len(graph))
	ans := make([]int, 0, len(graph))
	var dfs func(int) bool
	dfs = func(i int) bool {
		if _, ok := safe[i]; ok {
			return safe[i]
		}
		// 重复访问肯定不是安全的
		safe[i] = false
		// 默认是安全的
		flag := true
		for _, ni := range graph[i] {
			if !dfs(ni) {
				flag = false
				break
			}
		}
		safe[i] = flag
		return flag
	}
	for i := range graph {
		if dfs(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
