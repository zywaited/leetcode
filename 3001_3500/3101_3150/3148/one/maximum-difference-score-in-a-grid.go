package one

func maxScore(grid [][]int) int {
	ans := int(-1e9)
	n := len(grid[0])
	cols := make([]int, n+1)
	for i := range cols {
		cols[i] = int(1e9)
	}
	for i := range grid {
		row := int(1e9)
		for j := range grid[i] {
			row = minNumber(row, cols[j])
			ans = maxNumber(ans, grid[i][j]-row)
			row = minNumber(row, grid[i][j])
			cols[j] = row
		}
	}
	return ans
}

func minNumber(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func maxNumber(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
