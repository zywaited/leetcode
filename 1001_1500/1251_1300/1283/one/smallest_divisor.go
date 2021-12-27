package one

import "math"

func SmallestDivisor(nums []int, threshold int) int {
	num := 0
	r := 0
	for _, num = range nums {
		if num > r {
			r = num
		}
	}
	l := 1
	m := 0
	sum := 0
	for l < r {
		m = l + (r-l)>>1
		sum = 0
		for _, num = range nums {
			sum += int(math.Ceil(float64(num) / float64(m)))
		}
		if sum <= threshold {
			r = m
			continue
		}
		l = m + 1
	}
	return l
}
