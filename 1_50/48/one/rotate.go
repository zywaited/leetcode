package one

func Rotate(matrix [][]int) {
	s, n := 0, len(matrix)-1
	for s < n {
		i, j := s, n
		for i < n {
			tmp := matrix[s][i]
			matrix[s][i] = matrix[j][s]
			matrix[j][s] = matrix[n][j]
			matrix[n][j] = matrix[i][n]
			matrix[i][n] = tmp
			i++
			j--
		}
		s++
		n--
	}
}
