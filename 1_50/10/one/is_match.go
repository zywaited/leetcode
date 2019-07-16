package one

type key [2]int
type dict map[key]bool

func IsMatch(s string, p string) bool {
	m := make(dict)
	return dp(0, 0, s, p, m)
}

func dp(i, j int, s, p string, m dict) bool {
	k := key{i, j}
	if _, ok := m[k]; ok {
		return m[k]
	}

	l := len(s)
	pl := len(p)
	if j == pl {
		return i == l
	}

	first := i < l && (s[i] == p[j] || p[j] == '.')
	match := false
	if j < pl-1 && p[j+1] == '*' {
		match = dp(i, j+2, s, p, m) || first && dp(i+1, j, s, p, m)
	} else {
		match = first && dp(i+1, j+1, s, p, m)
	}

	m[k] = match
	return match
}
