package one

func MaxDepth(s string) int {
	num := 0
	depth := 0
	for i := range s {
		switch s[i] {
		case '(':
			num++
			depth = max(depth, num)
		case ')':
			num--
		}
	}
	return depth
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
