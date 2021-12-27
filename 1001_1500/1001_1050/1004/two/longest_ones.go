package two

func LongestOnes(A []int, K int) int {
	nl := 0
	zn := 0
	l, r := 0, 0
	for r < len(A) {
		if A[r] == 0 {
			zn++
		}
		for zn > K {
			if A[l] == 0 {
				zn--
			}
			l++
		}
		nl = max(nl, r-l+1)
		r++
	}
	return nl
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
