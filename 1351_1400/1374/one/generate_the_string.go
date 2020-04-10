package one

import "strings"

func GenerateTheString(n int) string {
	ans := strings.Builder{}
	if n&1 == 0 {
		ans.WriteByte('a')
		n--
	}
	for ; n > 0; n-- {
		ans.WriteByte('b')
	}
	return ans.String()
}
