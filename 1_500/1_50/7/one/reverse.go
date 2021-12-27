package one

func Reverse(x int) int {
	rv := 0
	max := 2<<30 - 1
	min := -2 << 30
	pop := 0
	maxP := max % 10
	minP := min % 10
	for x != 0 {
		pop = x % 10
		x /= 10
		// 边界值
		if rv > max/10 || (rv == max/10 && pop > maxP) {
			return 0
		}

		if rv < min/10 || (rv == min/10 && pop < minP) {
			return 0
		}

		rv = rv*10 + pop
	}
	return rv
}
