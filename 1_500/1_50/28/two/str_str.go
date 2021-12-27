package two

func StrStr(haystack string, needle string) int {
	ns := uint32(0)
	for j := 0; j < len(needle); j++ {
		ns += uint32(needle[j]) * 31
	}
	i := 0
	hs := uint32(0)
	for ; i < len(needle) && i < len(haystack); i++ {
		hs += uint32(haystack[i]) * 31
	}
	if hs == ns && haystack[:i] == needle {
		return 0
	}
	for ; i < len(haystack); i++ {
		hs += uint32(haystack[i]) * 31
		hs -= uint32(haystack[i-len(needle)]) * 31
		if hs == ns && haystack[i-len(needle)+1:i+1] == needle {
			return i - len(needle) + 1
		}
	}
	return -1
}
