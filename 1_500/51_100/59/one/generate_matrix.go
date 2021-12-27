package one

func GenerateMatrix(n int) [][]int {
	minRow, minCol := 0, 0
	maxRow, maxCol := n, n
	row, col := 0, 0
	num := 1
	rs := make([][]int, n)
	for minRow < maxRow && minCol < maxCol {
		if len(rs[minRow]) == 0 {
			rs[minRow] = make([]int, n)
		}
		for col = minCol; col < maxCol; col++ {
			rs[minRow][col] = num
			num++
		}
		for row = minRow + 1; row < maxRow; row++ {
			if len(rs[row]) == 0 {
				rs[row] = make([]int, n)
			}
			rs[row][maxCol-1] = num
			num++
		}
		if minRow < maxRow-1 && minCol < maxCol-1 {
			for col = maxCol - 2; col >= minCol; col-- {
				rs[maxRow-1][col] = num
				num++
			}
			for row = maxRow - 2; row > minRow; row-- {
				if len(rs[row]) == 0 {
					rs[row] = make([]int, n)
				}
				rs[row][minCol] = num
				num++
			}
		}
		minRow++
		maxRow--
		minCol++
		maxCol--
	}
	return rs
}
