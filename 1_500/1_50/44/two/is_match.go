package two

func IsMatch(s string, p string) bool {
	// 循环版
	sl := len(s)
	pl := len(p)
	start, match := -1, 0
	i, j := 0, 0
	for i < sl {
		// 正常匹配或者是通配符匹配
		if j < pl && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
			continue
		}
		// *匹配
		if j < pl && p[j] == '*' {
			// 保留当前匹配位置
			start, match = j, i
			// 跳过*，下一位
			j++
			continue
		}
		// 匹配失败查看是否可以回到*保留的位置
		if start == -1 {
			return false
		}
		// *号位置继续匹配
		j = start + 1
		// *匹配所有，所以此处加1继续匹配下一个
		match++
		i = match
	}
	// s匹配完成后，如果p还有字符就只能是*
	if j >= pl {
		return true
	}
	for _, as := range p[j:] {
		if as == '*' {
			continue
		}
		return false
	}
	return true
}
