package one

func Merge(A []int, m int, B []int, n int) {
	// 先填充尾部
	i := m - 1
	j := n - 1
	t := len(A) - 1
	for ; j >= 0; t-- {
		if i >= 0 {
			// 谁大先放
			if A[i] <= B[j] {
				A[t] = B[j]
				j--
			} else {
				A[t] = A[i]
				i--
			}
			continue
		}
		A[t] = B[j]
		j--
	}
}
