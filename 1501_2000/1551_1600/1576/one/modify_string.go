package one

import "strings"

func ModifyString(s string) string {
	ans := strings.Builder{}
	prev := byte('z')
	for i := range s {
		if s[i] != '?' {
			ans.WriteByte(s[i])
			prev = s[i]
			continue
		}
		prev = ((prev-'a')+1)%26 + 'a'
		if i < len(s)-1 && prev == s[i+1] {
			prev = ((prev-'a')+1)%26 + 'a'
		}
		ans.WriteByte(prev)
	}
	return ans.String()
}
