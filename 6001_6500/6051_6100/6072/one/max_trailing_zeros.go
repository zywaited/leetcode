package one

func MaxTrailingZeros(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	rows := make([][][2]int, m)
	cols := make([][][2]int, n)
	for i := range grid {
		rows[i] = make([][2]int, n+1)
		for j := range grid[i] {
			tn := twoNum(grid[i][j])
			fn := fiveNum(grid[i][j])
			rows[i][j+1][0] = rows[i][j][0] + tn
			rows[i][j+1][1] = rows[i][j][1] + fn
			if len(cols[j]) == 0 {
				cols[j] = make([][2]int, m+1)
			}
			cols[j][i+1][0] = cols[j][i][0] + tn
			cols[j][i+1][1] = cols[j][i][1] + fn
		}
	}
	ans := 0
	for i := range grid {
		ans = max(ans, min(rows[i][n][0], rows[i][n][1]))
		for j := range grid[i] {
			if i == 0 {
				ans = max(ans, min(cols[j][m][0], cols[j][m][1]))
			}
			ans = max(ans, min(rows[i][j+1][0]+cols[j][i][0], rows[i][j+1][1]+cols[j][i][1]))
			ans = max(ans, min(cols[j][i][0]+rows[i][n][0]-rows[i][j][0], cols[j][i][1]+rows[i][n][1]-rows[i][j][1]))
			ans = max(ans, min(rows[i][n][0]-rows[i][j][0]+cols[j][m][0]-cols[j][i+1][0], rows[i][n][1]-rows[i][j][1]+cols[j][m][1]-cols[j][i+1][1]))
			ans = max(ans, min(rows[i][j+1][0]+cols[j][m][0]-cols[j][i+1][0], rows[i][j+1][1]+cols[j][m][1]-cols[j][i+1][1]))
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

func twoNum(n int) int {
	num := 0
	for n > 0 && n%2 == 0 {
		num++
		n /= 2
	}
	return num
}

func fiveNum(n int) int {
	num := 0
	for n > 0 && n%5 == 0 {
		num++
		n /= 5
	}
	return num
}
