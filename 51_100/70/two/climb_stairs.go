package two

// 递归
// f(n) = f(n-1) + f(n-2)
func ClimbStairs(n int) int {
	// 数组和map都可以
	return steps(n, make(map[int]int))
}

func steps(n int, dp map[int]int) int {
	if n <= 2 {
		return n
	}
	if step, ok := dp[n]; ok {
		return step
	}
	step := steps(n-1, dp) + steps(n-2, dp)
	dp[n] = step
	return step
}
