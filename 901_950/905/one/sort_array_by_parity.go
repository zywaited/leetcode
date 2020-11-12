package one

func SortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		for ; i < len(A) && A[i]&1 == 0; i++ {
		}
		for ; j >= 0 && A[j]&1 == 1; j-- {
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}
	return A
}
