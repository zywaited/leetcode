package one

func LuckyNumbers(matrix [][]int) []int {
	ans := make([]int, 0)
	for i := range matrix {
		min := 0
		mj := -1
		for j := range matrix[i] {
			if mj == -1 || matrix[i][j] < min {
				mj = j
				min = matrix[i][j]
			}
		}
		// 当前列是否是最大的
		for j := range matrix {
			if matrix[j][mj] > min {
				min = 0
				break
			}
		}
		if min > 0 {
			ans = append(ans, min)
		}
	}
	return ans
}
