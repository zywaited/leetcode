package one

import "sort"

func MinIncrementForUnique(A []int) int {
	sort.Ints(A)
	ans := 0
	for index := 1; index < len(A); index++ {
		if A[index] <= A[index-1] {
			ans += A[index-1] - A[index] + 1
			A[index] = A[index-1] + 1
		}
	}
	return ans
}
