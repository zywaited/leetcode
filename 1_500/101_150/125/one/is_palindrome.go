package one

func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1
	// 验证是否是数字或者字母
	valid := func(index int) bool {
		return (s[index] >= '0' && s[index] <= '9') ||
			(s[index] >= 'a' && s[index] <= 'z') ||
			(s[index] >= 'A' && s[index] <= 'Z')
	}
	for l < r {
		// 跳过其他的字符
		if !valid(l) {
			l++
			continue
		}
		if !valid(r) {
			r--
			continue
		}
		tl := s[l]
		tr := s[r]
		// 忽略大小写
		if s[l] >= 'A' && s[l] <= 'Z' {
			tl += 32
		}
		if s[r] >= 'A' && s[r] <= 'Z' {
			tr += 32
		}
		if tl != tr {
			return false
		}
		l++
		r--
	}
	return true
}
