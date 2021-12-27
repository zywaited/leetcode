package one

func Reformat(s string) string {
	// 先算数量
	an := 0
	nn := 0
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			an++
			continue
		}
		nn++
	}
	if an-nn > 1 || an-nn < -1 {
		return ""
	}
	ai, ni := 0, 0
	ans := make([]byte, len(s))
	if an >= nn {
		// 字母在前
		ni = 1
	} else {
		ai = 1
	}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			ans[ai] = s[i]
			ai += 2
			continue
		}
		ans[ni] = s[i]
		ni += 2
	}
	return string(ans)
}
