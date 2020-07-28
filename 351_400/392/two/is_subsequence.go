package two

func IsSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}
	return si == len(s)
}
