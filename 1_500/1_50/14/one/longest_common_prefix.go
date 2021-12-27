package one

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	ans := find(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		if ans == "" {
			break
		}
		ans = find(ans, strs[i])
	}
	return ans
}

func find(f, s string) string {
	prefix := ""
	for i := len(f); i > 0; i-- {
		if i <= len(s) && s[0:i] == f[0:i] {
			prefix = f[0:i]
			break
		}
	}
	return prefix
}
