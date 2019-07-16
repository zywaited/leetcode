package one

// 就跟快速减法一致，n = n - 3, 如果可以，下次完全可以 n = n - 6，n = n - 12，尽量对半操作
func MyPow(x float64, n int) float64 {
	sum := float64(1)
	// 小于0需要转换
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := x
	for n != 0 {
		if n&1 == 1 {
			// 奇数, 则需要在乘以本身一次
			// a^4 = (a^2)^2
			// a^5 = (a^2)^2 * a
			sum *= tmp
		}
		// 对半操作
		tmp *= tmp
		n >>= 1
	}
	return sum
}
