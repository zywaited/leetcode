package one

type NumMatrix struct {
	sums [][]int
}

// 以下暂没有考虑边界值

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	rows, cols := len(matrix), len(matrix[0])
	sums := make([][]int, rows+1)
	for i := range matrix {
		sums[i+1] = make([]int, cols+1)
		for j := range matrix[i] {
			sums[i+1][j+1] = sums[i+1][j] + matrix[i][j]
		}
	}
	return NumMatrix{sums: sums}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	if len(nm.sums) == 0 {
		return sum
	}
	for row := row1; row <= row2; row++ {
		sum += nm.sums[row+1][col2+1] - nm.sums[row+1][col1]
	}
	return sum
}
