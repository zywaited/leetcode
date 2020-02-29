package one

import (
	"strconv"
	"strings"
)

func CompressString(S string) string {
	// 最大长度也就52
	bs := strings.Builder{}
	l := -1
	n := 0
	for i := range S {
		if n == 0 {
			bs.WriteByte(S[i])
		}
		// 数量
		if l == -1 || S[i] == S[l] {
			l = i
			n++
			continue
		}
		bs.WriteString(strconv.Itoa(n))
		l = i
		n = 1
		bs.WriteByte(S[i])
	}
	if n > 0 {
		bs.WriteString(strconv.Itoa(n))
	}
	if len(S) <= bs.Len() {
		return S
	}
	return bs.String()
}
