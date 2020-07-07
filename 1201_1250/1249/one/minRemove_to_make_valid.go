package one

func MinRemoveToMakeValid(s string) string {
	// 需要删除的索引
	rmi := make(map[int]bool, len(s))
	// (的索引列
	li := make([]int, 0, len(s))
	for i := range s {
		switch s[i] {
		case '(':
			li = append(li, i)
		case ')':
			if len(li) == 0 {
				rmi[i] = true
				break
			}
			li = li[:len(li)-1] // 还原
		}
	}
	// 剩下的都是有问题的
	for _, i := range li {
		rmi[i] = true
	}
	ans := make([]byte, 0, len(s)-len(rmi))
	for i := range s {
		if !rmi[i] {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
