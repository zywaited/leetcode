package one

// 循环
// f(n) = f(n-1) + f(n-2)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
