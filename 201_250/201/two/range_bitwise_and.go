package two

func RangeBitwiseAnd(m int, n int) int {
	// 简化第一种解法
	// 第一个解法后面补0
	// 这里直接就不位移直接保留0
	for n > m {
		n &= n - 1
	}
	return n
}
