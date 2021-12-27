package one

type key [2]int
type dict map[key]bool

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
	asterisk := j < pl && p[j] == '*'
	match := false
	for {
		if !asterisk {
			// 不是*号
			match = (i < l && (s[i] == p[j] || p[j] == '?')) && dp(i+1, j+1, s, p, m)
			break
		}
		// 特殊优化
		if j == pl-1 {
			// 匹配完成
			match = true
			break
		}
		if i >= l-1 {
			// s结束
			match = dp(i, j+1, s, p, m)
			break
		}
		// *号匹配空或者多个
		match = dp(i, j+1, s, p, m) || dp(i+1, j, s, p, m)
		break
	}

	m[k] = match
	return match
}

func IsMatch(s string, p string) bool {
	return dp(0, 0, s, p, make(dict))
}
