package one

var (
	// 返回结果
	res [][]string
	// 放置结果
	cq [][]byte
	// 皇后放置位置
	cds [][2]int
)

// 同一列、同一行、同一条对角线
func could(row, col int) bool {
	for _, cd := range cds {
		if row == cd[0] || col == cd[1] || row-cd[0] == col-cd[1] || row-cd[0]+col-cd[1] == 0 {
			return false
		}
	}
	return true
}

func bc(row, col, n int) {
	for j := col; j < n; j++ {
		if !could(row, j) {
			continue
		}
		cq[row][j] = 'Q'
		if row == n-1 {
			// 已经放置完成
			var r []string
			for _, cd := range cq {
				r = append(r, string(cd))
			}
			res = append(res, r)
			cq[row][j] = '.'
			continue
		}
		cds = append(cds, [2]int{row, j})
		bc(row+1, 0, n)
		cq[row][j] = '.'
		if len(cds) <= 1 {
			cds = nil
		} else {
			cds = cds[:len(cds)-1]
		}
	}
}

func SolveNQueens(n int) [][]string {
	res = nil
	cds = nil
	if n < 1 {
		return res
	}
	// 初始化为空
	cq = make([][]byte, n)
	for i := range cq {
		cq[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			cq[i][j] = '.'
		}
	}
	bc(0, 0, n)
	return res
}
