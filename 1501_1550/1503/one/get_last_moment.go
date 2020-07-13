package one

func GetLastMoment(n int, left []int, right []int) int {
	t := 0
	for _, i := range left {
		t = max(t, i)
	}
	for _, i := range right {
		t = max(t, n-i)
	}
	return t
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
