package one

import "sort"

// 排序 + 二分
func MaxAliveYear(birth []int, death []int) int {
	sort.Ints(birth)
	sort.Ints(death)
	var bs func(int) int
	bs = func(y int) int {
		s, e := 0, len(birth)-1
		n := 0
		m := 0
		for s <= e {
			m = s + (e-s)>>1
			if birth[m] <= y {
				s = m + 1
				n = m + 1
				continue
			}
			e = m - 1
		}
		return n
	}
	max := 0
	year := 0
	num := 0
	for i, y := range death {
		num = bs(y)
		if num-i > max {
			max = num - i
			year = birth[num-1]
		}
	}
	return year
}
