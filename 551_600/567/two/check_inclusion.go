package two

// 通过计算得分，判断两边是否相等
// 每一位代表一个字符，这样所有的字符相加即为分数
// 目前只要小写字母
func CheckInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	// 计算分数
	h1, h2 := 0, 0
	for i := range s1 {
		h1 += 1 << (s1[i] - 'a')
		h2 += 1 << (s2[i] - 'a')
	}
	if h1 == h2 {
		return true
	}
	// 如果s2还有多的字符
	// 那么后面的替代前面
	i, j := n1, 0
	for i < n2 {
		h2 += (1 << (s2[i] - 'a')) - (1 << (s2[j] - 'a'))
		if h1 == h2 {
			return true
		}
		i++
		j++
	}
	return false
}
