package two

// 动态规划
func UpdateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] = -1
			if i-1 >= 0 && matrix[i-1][j] >= 0 {
				matrix[i][j] = matrix[i-1][j] + 1
			}
			if j-1 >= 0 && matrix[i][j-1] >= 0 && (matrix[i][j] == -1 || matrix[i][j-1]+1 < matrix[i][j]) {
				matrix[i][j] = matrix[i][j-1] + 1
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			if i+1 < m && matrix[i+1][j] >= 0 && (matrix[i][j] == -1 || matrix[i+1][j]+1 < matrix[i][j]) {
				matrix[i][j] = matrix[i+1][j] + 1
			}
			if j+1 < n && matrix[i][j+1] >= 0 && (matrix[i][j] == -1 || matrix[i][j+1]+1 < matrix[i][j]) {
				matrix[i][j] = matrix[i][j+1] + 1
			}
		}
	}
	return matrix
}
