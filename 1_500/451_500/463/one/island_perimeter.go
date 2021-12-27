package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func IslandPerimeter(grid [][]int) int {
	ans := 0
	s, e := -1, -1
	// 找到起点
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				s, e = i, j
				break
			}
		}
		if s > -1 && e > -1 {
			break
		}
	}
	if s < 0 || e < 0 {
		return ans
	}
	var dfs func(int, int)
	dfs = func(i int, j int) {
		grid[i][j] = 2
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) || grid[ni][nj] == 0 {
				ans++
				continue
			}
			if grid[ni][nj] == 1 {
				dfs(ni, nj)
			}
		}
	}
	dfs(s, e)
	return ans
}
