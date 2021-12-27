package one

// 计算每一列是否需要删除
// 就是尽可能的保留
func MinDeletionSize(A []string) int {
	rs := 0
	if len(A) > 1 {
		for c := range A[0] {
			for r := 0; r < len(A)-1; r++ {
				if A[r][c] > A[r+1][c] {
					rs++
					break
				}
			}
		}
	}
	return rs
}
