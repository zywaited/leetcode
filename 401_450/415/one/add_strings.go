package one

func AddStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	m := n1
	if n2 > m {
		m = n2
	}
	rs := make([]byte, m+1) // 预留1个进位（这样最后也不用反转字符串）
	var (
		c int // 进位
		n int // 当前位和
	)
	i, j, k := n1-1, n2-1, m // 两个字符串索引, 从尾往前加
	for i >= 0 || j >= 0 {
		n = c
		if i >= 0 {
			n += int(num1[i] - '0')
		}
		if j >= 0 {
			n += int(num2[j] - '0')
		}
		c, rs[k] = n/10, byte(n%10)+'0'
		i--
		j--
		k--
	}
	if c > 0 {
		rs[0] = byte(c) + '0'
		return string(rs)
	}
	return string(rs[1:])
}
