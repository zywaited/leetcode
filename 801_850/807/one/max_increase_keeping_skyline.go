package one

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	rows := make([]int, len(grid))
	cols := make([]int, len(grid))
	for i := range grid {
		for j := range grid[i] {
			rows[i] = max(rows[i], grid[i][j])
			cols[j] = max(cols[j], grid[i][j])
		}
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			ans += min(rows[i], cols[j]) - grid[i][j]
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
