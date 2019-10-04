package one

func Calculator(s string) int {
	// 采用栈
	// 计算完后或者遇到括号需要压入栈(优先级问题)
	stack := make([]int, 0, len(s))
	num := 0     // 当前值
	nextNum := 0 // 下一个数字
	symbol := 1  // -1: 减法; 1: 加法
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			stack = append(stack, num)    // 把当前值推到栈中
			stack = append(stack, symbol) // 当前的符号
			num = 0
			symbol = 1
			continue
		}
		if s[i] == ')' {
			// 取出前面的数和符号
			num = stack[len(stack)-2] + stack[len(stack)-1]*num
			stack = stack[:len(stack)-2]
			continue
		}
		// 符号
		if s[i] == '+' {
			symbol = 1
			continue
		}
		if s[i] == '-' {
			symbol = -1
			continue
		}
		// 处理数字
		nextNum = int(s[i] - '0')
		for i = i + 1; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				nextNum = nextNum*10 + int(s[i]-'0')
				continue
			}
			break
		}
		num += symbol * nextNum
		// 最外层for循环会i++
		i--
	}
	return num
}
