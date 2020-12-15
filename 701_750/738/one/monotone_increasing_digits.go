package one

func MonotoneIncreasingDigits(N int) int {
	ans := 0
	if N == 0 {
		return ans
	}
	n := -1
	pn := -1
	v := 0
	ten := 1
	for ; N > 0; N /= 10 {
		v = N % 10
		n++
		if n > 0 {
			ten *= 10
		}
		if pn == -1 {
			ans = v
			pn = v
			continue
		}
		if v <= pn {
			ans += v * ten
			pn = v
			continue
		}
		ans = v*ten - 1
		pn = v - 1
	}
	if pn == 0 {
		ans = ten - 1
	}
	return ans
}
