package one

func Add(a int, b int) int {
	// a&b > 0 代表着对应的位都要进1
	// a^b 不相等的位变成1
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}
