package one

import "math"

func PathInZigZagTree(label int) []int {
	level := 1
	for ; 1<<uint(level) <= label; level++ {
	}
	ans := make([]int, level)
	total := label
	current := int(math.Pow(float64(2), float64(level)))
	for index := level - 1; index >= 0; index-- {
		ans[index] = total
		current >>= 1
		total = current - (total-current)>>1 - 1
	}
	return ans
}
