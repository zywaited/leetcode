package one

const (
	unDivide = iota
	divideA
	divideB
)

func IsBipartite(graph [][]int) bool {
	ds := make([]int, len(graph))
	var dfs func(int, int) bool
	dfs = func(i int, d int) bool {
		ds[i] = d
		nd := divideA
		// 下一个相连点是相反的集合
		if d == divideA {
			nd = divideB
		}
		for _, ni := range graph[i] {
			if ds[ni] == unDivide {
				if !dfs(ni, nd) {
					return false
				}
				continue
			}
			if ds[ni] != nd {
				return false
			}
		}
		return true
	}
	for i := range graph {
		if ds[i] == unDivide && !dfs(i, divideA) {
			return false
		}
	}
	return true
}
