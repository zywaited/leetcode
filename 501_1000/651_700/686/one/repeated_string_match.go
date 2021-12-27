package one

import "strings"

func RepeatedStringMatch(a string, b string) int {
	i := strings.Index(b, a)
	if i < 0 {
		if strings.Contains(a, b) {
			return 1
		}
		if strings.Contains(a+a, b) {
			return 2
		}
		return -1
	}
	ans := 1
	j := i + len(a)
	for ; j+len(a) <= len(b); j += len(a) {
		if a != b[j:j+len(a)] {
			return -1
		}
		ans++
	}
	if i == 0 && j == len(b) {
		return ans
	}
	if (i > 0 && !strings.HasSuffix(a, b[:i])) || (j < len(b) && !strings.HasPrefix(a, b[j:])) {
		return -1
	}
	if i > 0 {
		ans++
	}
	if j < len(b) {
		ans++
	}
	return ans
}
