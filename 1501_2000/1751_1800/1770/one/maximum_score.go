package one

import "math"

func MaximumScore(nums []int, multipliers []int) int {
	l := make([]int, len(multipliers)+1)
	r := make([]int, len(multipliers)+1)
	tl := make([]int, len(multipliers)+1)
	tr := make([]int, len(multipliers)+1)
	for i := 1; i <= len(multipliers); i++ {
		l, tl = tl, l
		r, tr = tr, r
		for j := i; j > 0; j-- {
			l[j] = nums[j-1] * multipliers[i-1]
			r[j] = nums[len(nums)-j] * multipliers[i-1]
			if j-1 > 0 && i-j > 0 {
				l[j] += max(tl[j-1], tr[i-j])
				r[j] += max(tr[j-1], tl[i-j])
				continue
			}
			if j-1 > 0 {
				l[j] += tl[j-1]
				r[j] += tr[j-1]
				continue
			}
			l[j] += tr[i-j]
			r[j] += tl[i-j]
		}
	}
	m := math.MinInt32
	for i := 1; i <= len(multipliers); i++ {
		m = max(m, max(l[i], r[i]))
	}
	return m
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
