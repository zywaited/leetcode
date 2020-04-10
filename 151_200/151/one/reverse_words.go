package one

import "strings"

func ReverseWords(s string) string {
	ans := strings.Builder{}
	ei := len(s)
	bi := ei - 1
	for ; bi >= 0; bi-- {
		if s[bi] != ' ' {
			continue
		}
		if s[bi] == ' ' && ei-bi < 2 {
			ei = bi
			continue
		}
		ans.WriteString(s[bi+1 : ei])
		ans.WriteByte(' ')
		ei = bi
	}
	if ei-bi >= 2 {
		ans.WriteString(s[bi+1 : ei])
		return ans.String()
	}
	if ans.Len() == 0 {
		return ans.String()
	}
	return ans.String()[:ans.Len()-1]
}
