package one

func StrStr(haystack string, needle string) int {
	i := 0
	j := 0
	if len(needle) > 0 {
		ns := next(needle)
		for i < len(haystack) && j < len(needle) {
			if j == -1 || haystack[i] == needle[j] {
				i++
				j++
				continue
			}
			j = ns[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func next(s string) []int {
	ns := make([]int, len(s))
	ns[0] = -1 // 去除第一个字符(后缀)
	k := -1
	i := 0
	for i < len(s)-1 { // 去除最后一个字符(前缀)
		if k == -1 || s[i] == s[k] {
			k++
			i++
			ns[i] = k
			continue
		}
		k = ns[k]
	}
	return ns
}
