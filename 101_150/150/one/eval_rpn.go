package one

import "strconv"

// 总是有效的，所以不必考虑数量和转换
func EvalRPN(tokens []string) int {
	// 采用栈
	// 计算完值后重新入栈
	stack := make([]int, 0, len(tokens))
	symbol := map[string]byte{ // 符号
		"+": 1,
		"-": 1,
		"*": 1,
		"/": 1,
	}
	for _, token := range tokens {
		// 如果不是符号就压入栈中
		if symbol[token] == 0 {
			// 转成数字
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			continue
		}
		// 取出前两个数
		f, s := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, f+s)
		case "-":
			stack = append(stack, f-s)
		case "*":
			stack = append(stack, f*s)
		case "/":
			stack = append(stack, f/s)
		}
	}
	// 最后stack只会剩下一个
	if len(stack) < 1 {
		return 0
	}
	return stack[0]
}
