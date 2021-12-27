package one

func NumOfMinutes(n int, headID int, manager []int, informTime []int) int {
	// 数据转换成map方便处理
	sm := make(map[int][]int, len(manager))
	for sr, mr := range manager {
		sm[mr] = append(sm[mr], sr)
	}
	var dfs func(int) int
	dfs = func(hd int) int {
		if len(sm[hd]) == 0 {
			return 0
		}
		cs := informTime[hd]
		for _, nhd := range sm[hd] {
			cs = max(cs, informTime[hd]+dfs(nhd))
		}
		return cs
	}
	return dfs(headID)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
