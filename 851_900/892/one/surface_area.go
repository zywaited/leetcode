package one

func SurfaceArea(grid [][]int) int {
	ans := 0
	for i := range grid {
		for j, num := range grid[i] {
			if num == 0 {
				continue
			}
			ans += num*4 + 2
			// 如果前和左有相连的要去除接触的面积
			if i > 0 && grid[i-1][j] > 0 {
				ans -= min(grid[i-1][j], num) * 2
			}
			if j > 0 && grid[i][j-1] > 0 {
				ans -= min(grid[i][j-1], num) * 2
			}
		}
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
