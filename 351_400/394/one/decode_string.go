package one

func DecodeString(s string) string {
	ans := make([]byte, 0)
	// 存储数量(这里就用循环处理，不想递归了)
	stack := make([][2]int, 0)
	// 数字
	num := 0
	for i := range s {
		if s[i] == '[' {
			stack = append(stack, [2]int{num, len(ans)})
			num = 0 // 重置数字
			continue
		}
		if number(s[i]) {
			// 计算数字
			num = num*10 + int(s[i]-'0')
			continue
		}
		if s[i] != ']' {
			// 正常字符
			ans = append(ans, s[i])
			continue
		}
		// 闭环
		pn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		bl := len(ans)
		// bs已经存了一份了，所以这里要大于1
		for ; pn[0] > 1; pn[0]-- {
			ans = append(ans, ans[pn[1]:bl]...)
		}
	}
	return string(ans)
}

func number(n byte) bool {
	return n >= '0' && n <= '9'
}
