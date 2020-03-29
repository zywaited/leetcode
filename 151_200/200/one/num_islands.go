package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func NumIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

// 这里直接暴力改原数据节约内存
func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	// 重置当前的值
	grid[i][j] = '0'
	for _, direct := range directs {
		dfs(grid, i+direct[0], j+direct[1])
	}
}
