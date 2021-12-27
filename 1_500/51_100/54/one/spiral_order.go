package one

func SpiralOrder(matrix [][]int) []int {
	minRow, minCol := 0, 0 // 正在访问的行列
	maxRow, maxCol := len(matrix), 0
	if maxRow > 0 {
		maxCol = len(matrix[0])
	}
	row, col := 0, 0
	rs := make([]int, 0, maxRow*maxCol)
	for minRow < maxRow && minCol < maxCol {
		// 上 行
		for col = minCol; col < maxCol; col++ {
			rs = append(rs, matrix[minRow][col])
		}
		// 右 列
		for row = minRow + 1; row < maxRow; row++ {
			rs = append(rs, matrix[row][maxCol-1])
		}
		if minRow < maxRow-1 && minCol < maxCol-1 {
			// 下 行
			for col = maxCol - 2; col >= minCol; col-- {
				rs = append(rs, matrix[maxRow-1][col])
			}
			// 左 列
			for row = maxRow - 2; row > minRow; row-- {
				rs = append(rs, matrix[row][minCol])
			}
		}
		// 一圈完成
		minRow++
		maxRow--
		minCol++
		maxCol--
	}
	return rs
}
