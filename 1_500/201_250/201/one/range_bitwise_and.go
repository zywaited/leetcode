package one

func RangeBitwiseAnd(m int, n int) int {
	// 找公共前缀, 然后补上0
	num := uint(0)
	for n > m {
		n >>= 1
		m >>= 1
		num++
	}
	return m << num
}
