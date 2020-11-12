package one

func SortArrayByParityII(A []int) []int {
	odd, even := make([]int, 0, len(A)>>1), make([]int, 0, len(A)>>1)
	for index, num := range A {
		if index&1 == num&1 {
			continue
		}
		if index&1 == 1 {
			if len(odd) == 0 {
				even = append(even, index)
				continue
			}
			A[index], A[odd[0]] = A[odd[0]], A[index]
			odd = odd[1:]
			continue
		}
		if len(even) == 0 {
			odd = append(odd, index)
			continue
		}
		A[index], A[even[0]] = A[even[0]], A[index]
		even = even[1:]
	}
	return A
}
