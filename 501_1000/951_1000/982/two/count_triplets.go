package two

// 两个数找到对应的1个数
func CountTriplets(A []int) int {
	// 确定一个数
	vs := make(map[int]int, 1<<16)
	for _, num := range A {
		// 两个数异或就是第三个数所有可为1的位
		// 因为这两个数某位都为1，另一个数一定为0才可以
		mk := (1<<16 - 1) ^ num
		m := mk
		for m > 0 {
			vs[m]++
			// 该位可为1，那么也可为0，但也要保证其他位1的情况，因为-1会导致后面位可能变成1，所以要&mk
			m = (m - 1) & mk
		}
		// 0肯定可以的
		vs[0]++
	}
	ans := 0
	for i := range A {
		for j := range A {
			ans += vs[A[i]&A[j]]
		}
	}
	return ans
}
