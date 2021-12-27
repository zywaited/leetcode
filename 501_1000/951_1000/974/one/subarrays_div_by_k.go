package one

func SubarraysDivByK(A []int, K int) int {
	ans := 0
	// 前缀和的余数
	mn := make(map[int]int, len(A))
	mn[0] = 1
	sum := 0
	for _, num := range A {
		sum += num
		// 余数一致
		m := (sum%K + K) % K
		// 找到当前余数的数量(当前和减去这些余数对应的和就能整除K)
		ans += mn[m]

		// 如果不按照上述统一转换为正余数，需要处理两种情况
		//m := sum % K
		//ans += mn[m]
		//if m > 0 {
		//	ans += mn[m-K]
		//}
		//if m < 0 {
		//	ans += mn[m+K]
		//}

		mn[m]++
	}
	return ans
}
