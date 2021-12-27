package one

func Calculate(s string) int {
	stack := []int{0}
	var (
		index       = 0
		symbol byte = '+'
		num         = 0
	)
	for ; index < len(s); index++ {
		if s[index] == ' ' {
			continue
		}
		if s[index] < '0' || s[index] > '9' {
			symbol = s[index]
			continue
		}
		num = 0
		for ; index < len(s) && s[index] >= '0' && s[index] <= '9'; index++ {
			num = num*10 + int(s[index]-'0')
		}
		index--
		switch symbol {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			stack[len(stack)-1] *= num
		case '/':
			stack[len(stack)-1] /= num
		}
	}
	index = len(stack) - 1
	for ; index > 0; index-- {
		stack[index-1] += stack[index]
	}
	return stack[0]
}
