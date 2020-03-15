package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func MaxAreaOfIsland(grid [][]int) int {
	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == 0 {
			return 0
		}
		// 这里改值为0代表访问过了
		// 如果不能修改原数据就用hash
		grid[row][col] = 0
		num := 1
		for _, direct := range directs {
			num += dfs(row+direct[0], col+direct[1])
		}
		return num
	}
	ans := 0
	for row := range grid {
		for col := range grid[row] {
			num := dfs(row, col)
			if num > ans {
				ans = num
			}
		}
	}
	return ans
}
