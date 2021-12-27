package two

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ev := make(map[string]map[string]float64)
	for index, equation := range equations {
		if ev[equation[0]] == nil {
			ev[equation[0]] = make(map[string]float64)
		}
		if ev[equation[1]] == nil {
			ev[equation[1]] = make(map[string]float64)
		}
		ev[equation[0]][equation[1]] = values[index]
		ev[equation[1]][equation[0]] = 1.0 / values[index]
		ev[equation[0]][equation[0]] = 1.0
		ev[equation[1]][equation[1]] = 1.0
	}
	used := make(map[string]int)
	var dfs func(int, string, string) float64
	dfs = func(index int, f string, s string) float64 {
		if used[f] == index || len(ev[f]) == 0 {
			return -1.0
		}
		if _, ok := ev[f][s]; ok {
			return ev[f][s]
		}
		used[f] = index
		for nf, v := range ev[f] {
			nv := dfs(index, nf, s)
			if nv > -1.0 {
				// 优化点：写入结果中 [合并]
				ev[f][s] = v * nv
				ev[s][f] = 1.0 / ev[f][s]
				return ev[f][s]
			}
		}
		return -1.0
	}
	ans := make([]float64, len(queries))
	for index, query := range queries {
		ans[index] = dfs(index+1, query[0], query[1])
	}
	return ans
}
