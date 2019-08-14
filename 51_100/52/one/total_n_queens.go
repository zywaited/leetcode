package one

func TotalNQueens(n int) int {
	rs := 0                     // 结果
	cq := make([][]int, n)      // 放置结果
	cds := make([][2]int, 0, n) // 皇后放置点
	// 初始化
	for row := 0; row < n; row++ {
		cq[row] = make([]int, n)
		for col := 0; col < n; col++ {
			cq[row][col] = '.'
		}
	}
	bc(cq, cds, 0, 0, n, &rs)
	return rs
}

// 回溯，每一个点都试一下
func bc(cq [][]int, cds [][2]int, row, index, n int, rs *int) {
	for col := index; col < n; col++ {
		if !could(row, col, cds) {
			continue
		}
		// 是否已经放置完成
		if row == n-1 {
			*rs++
			continue
		}
		cq[row][col] = 'Q'
		// 放置下一行
		bc(cq, append(cds, [2]int{row, col}), row+1, 0, n, rs)
		cq[row][col] = '.'
	}
}

// 同一列、同一行、对角线
func could(row, col int, cds [][2]int) bool {
	for _, cd := range cds {
		if row == cd[0] || col == cd[1] || row-cd[0] == col-cd[1] || row-cd[0]+col-cd[1] == 0 {
			return false
		}
	}
	return true
}
