package one

func IsRectangleOverlap(rec1 []int, rec2 []int) bool {
	return min(rec1[2], rec2[2]) > max(rec1[0], rec2[0]) && min(rec1[3], rec2[3]) > max(rec1[1], rec2[1])
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
