package one

func MaximumGain(s string, x int, y int) int {
	score := 0
	ca := 0
	cb := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			ca++
			if y >= x && cb > 0 {
				ca--
				cb--
				score += y
			}
		case 'b':
			cb++
			if x >= y && ca > 0 {
				ca--
				cb--
				score += x
			}
		default:
			score += min(x, y) * min(ca, cb)
			ca = 0
			cb = 0
		}
	}
	return score + min(x, y)*min(ca, cb)
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
