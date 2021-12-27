package one

// 4的幂，如4，16，64等本身是2的幂并且减1后是3的倍数
func IsPowerOfFour(num int) bool {
	return num&(num-1) == 0 && num%3 == 1
}
