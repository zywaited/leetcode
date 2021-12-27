package one

// 题目说明了是有序的, 可以每列进行二分查找
func CountNegatives(grid [][]int) int {
	var bs func([]int) int
	bs = func(row []int) int {
		s, e := 0, len(row)-1
		n := 0
		for s <= e {
			m := s + (e-s)>>1
			if row[m] >= 0 {
				s = m + 1
				continue
			}
			e = m - 1
			n = len(row) - m
		}
		return n
	}
	ans := 0
	for _, row := range grid {
		ans += bs(row)
	}
	return ans
}
