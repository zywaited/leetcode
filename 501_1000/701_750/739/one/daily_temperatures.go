package one

func DailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	// 用于保存索引
	stack := make([]int, 0, len(T))
	stack = append(stack, len(T)-1)
	for i := len(T) - 2; i >= 0; i-- {
		// 去除比当前小的温度值
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}
