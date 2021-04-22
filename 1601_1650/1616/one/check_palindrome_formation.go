package one

func CheckPalindromeFormation(a string, b string) bool {
	return checkJoinPalindrome(a, b) || checkJoinPalindrome(b, a)
}

func checkJoinPalindrome(a, b string) bool {
	s, e := 0, len(a)-1
	for e >= s && a[s] == b[e] {
		s++
		e--
	}
	// 中间缺失的字符串是否是回文串
	return checkPalindrome(a, s, e) || checkPalindrome(b, s, e)
}

func checkPalindrome(p string, s, e int) bool {
	for s < e {
		if p[s] == p[e] {
			s++
			e--
			continue
		}
		return false
	}
	return true
}
