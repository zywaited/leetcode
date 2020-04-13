package one

// 就是找到所有的路径算概率和
// 深度优先搜索
func FrogPosition(n int, edges [][]int, t int, target int) float64 {
	// 当前点能跳到的定点
	cps := make(map[int][]int, n)
	for _, edge := range edges {
		cps[edge[0]] = append(cps[edge[0]], edge[1])
		cps[edge[1]] = append(cps[edge[1]], edge[0])
	}
	// 顶点是否使用过
	used := make(map[int]byte, n)
	var dfs func(int, int) float64
	dfs = func(ci, ct int) float64 {
		// 刚好找到了
		if ci == target && ct == 0 {
			return 1
		}
		// 没时间了
		if ct == 0 {
			return 0
		}
		used[ci] = 1
		// 可以到达下一个顶点数量
		n := 0
		rake := float64(0)
		for _, ni := range cps[ci] {
			if used[ni] == 1 {
				continue
			}
			// 不能停留在当前
			if ci == target {
				return 0
			}
			n++
			rake += dfs(ni, ct-1)
		}
		used[ci] = 0
		// 只能原地停留
		if ci == target {
			return 1
		}
		if n == 0 || rake == 0 {
			return 0
		}
		return rake / float64(n)
	}
	return dfs(1, t)
}
