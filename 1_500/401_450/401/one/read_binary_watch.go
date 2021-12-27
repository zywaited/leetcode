package one

import "strconv"

func ReadBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	var bc func(n, hi, mi int)
	h := 0
	m := 0
	bc = func(n, hi, mi int) {
		if n == 0 {
			ans = append(ans, strconv.Itoa(h)+":"+strconv.Itoa(m/10)+strconv.Itoa(m%10))
			return
		}
		for ; mi < 6 && (m|1<<uint(mi)) < 60; mi++ {
			m |= 1 << uint(mi)
			bc(n-1, hi, mi+1)
			m ^= 1 << uint(mi)
		}
		for ; hi < 4 && (h|1<<uint(hi)) < 12; hi++ {
			h |= 1 << uint(hi)
			bc(n-1, hi+1, mi)
			h ^= 1 << uint(hi)
		}
	}
	bc(turnedOn, 0, 0)
	return ans
}
