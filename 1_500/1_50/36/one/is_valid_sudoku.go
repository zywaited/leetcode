package one

// 注意：只是判断给的数独是否有效，并不是一定要有解
func IsValidSudoku(board [][]byte) bool {
	var (
		rows [9][10]int
		cols [9][10]int
		// (row/3)*3 + col/3
		// 0 1 2
		// 3 4 5
		// 6 7 8
		boxes [9][10]int
	)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := int(board[row][col] - '0')
				idx := (row/3)*3 + col/3
				if rows[row][num] > 0 || cols[col][num] > 0 || boxes[idx][num] > 0 {
					return false
				}
				rows[row][num] = 1
				cols[col][num] = 1
				boxes[idx][num] = 1
			}
		}
	}
	return true
}
