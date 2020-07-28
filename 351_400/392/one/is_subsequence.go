package one

func IsSubsequence(s string, t string) bool {
	ts := make([][]int, 26)
	for i := range t {
		ts[t[i]-'a'] = append(ts[t[i]-'a'], i)
	}
	li := -1
	for i := range s {
		if len(ts[s[i]-'a']) == 0 {
			return false
		}
		l, r := 0, len(ts[s[i]-'a'])-1
		for l <= r {
			m := l + (r-l)>>1
			if ts[s[i]-'a'][m] <= li {
				l = m + 1
				continue
			}
			r = m - 1
		}
		if l >= len(ts[s[i]-'a']) || ts[s[i]-'a'][l] <= li {
			return false
		}
		li = ts[s[i]-'a'][l]
	}
	return true
}
