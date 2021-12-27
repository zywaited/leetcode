package one

func HIndex(citations []int) int {
	l, r := 0, len(citations)-1
	m := 0
	for l < r {
		m = l + (r-l)>>1
		if citations[m] > 0 && citations[m] >= len(citations)-m {
			r = m
			continue
		}
		l = m + 1
	}
	if citations[l] == 0 {
		return 0
	}
	return len(citations) - l
}
