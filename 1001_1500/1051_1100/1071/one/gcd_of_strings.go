package one

func GcdOfStrings(str1 string, str2 string) string {
	s1, s2 := len(str1), len(str2)
	t := str1[0:gcd(s1, s2)]
	if valid(str1, t) && valid(str2, t) {
		return t
	}
	return ""
}

func valid(s, t string) bool {
	for i := len(s) / len(t); i > 0; i-- {
		if t != s[len(t)*(i-1):len(t)*i] {
			return false
		}
	}
	return true
}

// 欧几里得算法找最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
