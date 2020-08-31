package one

func ReverseWords(s string) string {
	ans := make([]byte, 0, len(s))
	li := 0
	ti := 0
	i := 0
	for ; i < len(s); i++ {
		for i < len(s) && s[i] != ' ' {
			i++
		}
		for ti = i - 1; ti >= li; ti-- {
			ans = append(ans, s[ti])
		}
		if i < len(s) {
			ans = append(ans, s[i])
		}
		li = i + 1
	}
	return string(ans)
}
