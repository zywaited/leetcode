package one

func LongestValidParentheses(s string) int {
	res := 0
	if len(s) < 2 {
		return res
	}
	cn := 0
	li := 0
	// 初始化
	// stack始终不为空，是为了(())能够计算完整长度为4 = 3-(-1)
	stack := []int{-1}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) < 1 {
			stack = []int{i}
			continue
		}
		li = stack[len(stack)-1]
		cn = i - li
		if cn > res {
			res = cn
		}
	}
	return res
}
