package one

func MatrixScore(A [][]int) int {
	col := make([]int, len(A[0]))
	for i := range A {
		for j := range A[i] {
			if j == 0 {
				if A[i][j] == 0 {
					// 横向移动
					for k := 0; k < len(A[i]); k++ {
						A[i][k] ^= 1
					}
				}
				continue
			}
			col[j] += A[i][j]
			if i < len(A)-1 || col[j] >= ((len(A)+1)>>1) {
				continue
			}
			for k := 0; k < len(A); k++ {
				A[k][j] ^= 1
			}
		}
	}
	s := 0
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1 {
				s += 1 << uint(len(A[i])-j-1)
			}
		}
	}
	return s
}
