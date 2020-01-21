package one

const max int = 9

var directs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func Domino(n int, m int, broken [][]int) int {
	// 不能访问的节点
	unable := make(map[int]byte, len(broken))
	// 相连节点
	matched := make(map[int]int, len(broken))
	for _, broke := range broken {
		score := broke[0]*max + broke[1]
		unable[score] = 1
		matched[score] = score
	}
	// 访问过的节点
	visited := make(map[int]int, len(broken))
	base := 0 // 为了不每次DFS初始化visited
	var dfs func(int) bool
	dfs = func(os int) bool {
		i, j := os/max, os%max
		for _, direct := range directs {
			x := i + direct[0]
			y := j + direct[1]
			if x >= 0 && x < n && y >= 0 && y < m {
				cs := x*max + y
				if visited[cs] == base || unable[cs] == 1 {
					continue
				}
				visited[cs] = base
				ds, ok := matched[cs]
				if !ok || dfs(ds) {
					matched[os] = cs
					matched[cs] = os
					return true
				}
			}
		}
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			score := i*max + j
			base = score
			if _, ok := matched[score]; ok || unable[score] == 1 {
				continue
			}
			if dfs(score) {
				sum++
			}
		}
	}
	return sum
}
