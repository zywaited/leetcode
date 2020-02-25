package one

import "math"

func ClosestDivisors(num int) []int {
	ans := []int{0, math.MaxInt32}
	// 因数以平方根为基础
	handle := func(n int) {
		// 为了以防数据有误+1
		sq := int(math.Sqrt(float64(n))) + 1
		// 反着查一个数据相隔最近
		for ; sq > 0; sq-- {
			if n%sq != 0 {
				continue
			}
			if abs(sq-n/sq) < abs(ans[1]-ans[0]) {
				ans[0] = n / sq
				ans[1] = sq
			}
			return
		}
	}
	handle(num + 1)
	handle(num + 2)
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
