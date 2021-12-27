package one

func AddBinary(a string, b string) string {
	ans := make([]byte, max(len(a), len(b))+1)
	ai := len(ans)
	i, j := len(a), len(b)
	ceil := byte(0)
	for i > 0 || j > 0 {
		// 这部分放在前面是为了后面分支看着简洁点
		ai--
		i--
		j--
		if i >= 0 && j >= 0 {
			if a[i] == '1' && b[j] == '1' {
				ans[ai] = ceil + '0'
				ceil = 1
				continue
			}
			if a[i] == '1' || b[j] == '1' {
				ans[ai] = (1 - ceil) + '0'
				continue
			}
			ans[ai] = ceil + '0'
			ceil = 0
			continue
		}
		c := byte('0')
		if i >= 0 {
			c = a[i]
		} else {
			c = b[j]
		}
		if c == '1' {
			ans[ai] = (1 - ceil) + '0'
			continue
		}
		ans[ai] = ceil + '0'
		ceil = 0
	}
	if ceil == 0 {
		ans = ans[1:]
	}
	if ceil == 1 {
		ai--
		ans[ai] = ceil + '0'
	}
	return string(ans)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
