package one

// 题目没有明确说数值的范围
// 不然就可以标记删除
func SetZeroes(matrix [][]int) {
	row := make(map[int]byte)
	col := make(map[int]byte)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := range row {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	for j := range col {
		for i := len(matrix) - 1; i >= 0; i-- {
			matrix[i][j] = 0
		}
	}
}
