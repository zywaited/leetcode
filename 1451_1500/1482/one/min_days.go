package one

func MinDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	l := 0
	r := 0
	for _, day := range bloomDay {
		r += day
	}
	for l < r {
		d := l + (r-l)>>1
		tk := 0
		tm := 0
		for _, day := range bloomDay {
			if day > d {
				tk = 0
				continue
			}
			tk++
			if tk == k {
				tk = 0
				tm++
			}
			if tm == m {
				break
			}
		}
		if tm == m {
			r = d
			continue
		}
		l = d + 1
	}
	return l
}
