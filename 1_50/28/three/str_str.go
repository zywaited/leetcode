package three

func StrStr(haystack string, needle string) int {
	co := make(map[byte]int, len(needle))
	for i := range needle {
		co[needle[i]] = len(needle) - i // 这里是是为了对其haystack中的字段
	}
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
			continue
		}
		// 没有多余的字符
		ni := i + len(needle) - j
		if ni >= len(haystack) {
			break
		}
		i -= j
		j = 0
		if _, ok := co[haystack[ni]]; ok {
			i += co[haystack[ni]]
			continue
		}
		i += len(needle) + 1
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}
