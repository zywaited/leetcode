package one

// 1个数和另外两个数
func CountTriplets(A []int) int {
	// 先求另外两个数的值
	vs := make(map[int]int, len(A))
	for i := range A {
		for j := range A {
			vs[A[i]&A[j]]++
		}
	}
	ans := 0
	for i := range A {
		for v := range vs {
			if A[i]&v == 0 {
				ans += vs[v]
			}
		}
	}
	return ans
}
