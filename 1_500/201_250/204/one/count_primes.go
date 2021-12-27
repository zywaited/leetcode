package one

func CountPrimes(n int) int {
	isNotPrimes := make([]bool, n+1) // 不是质数
	rs := 0
	for f := 2; f*f < n; f++ {
		if !isNotPrimes[f] {
			// 这里也要平方
			// 因为算3的时候会计算3*3,3*4
			// 算4的时候就直接算4*4,4*5，而不是重复计算4*3
			for s := f * f; s < n; s += f {
				// 每次取f的倍数
				isNotPrimes[s] = true
			}
		}
	}
	// 获取质数的个数
	for num := 2; num < n; num++ {
		if !isNotPrimes[num] {
			rs++
		}
	}
	return rs
}
