package two

func SortArrayByParityII(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		for ; i < len(A) && i&1 == A[i]&1; i++ {
		}
		for ; j >= 0 && j&1 == A[j]&1; j-- {
		}
		if i >= j {
			break
		}
		if i&1 != j&1 {
			A[i], A[j] = A[j], A[i]
			continue
		}
		k := i
		for ; k < len(A) && A[k]&1 != i&1; k++ {
		}
		A[i], A[k] = A[k], A[i]
		i++
	}
	return A
}
