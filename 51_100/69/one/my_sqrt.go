package one

func MySqrt(x int) int {
	ans := 0
	s, e := 0, x
	for s <= e {
		m := s + (e-s)>>1
		if m*m <= x {
			ans = m
			s = m + 1
			continue
		}
		e = m - 1
	}
	return ans
}
