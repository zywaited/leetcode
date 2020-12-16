package one

func WordPattern(pattern string, s string) bool {
	ps := [26]string{}
	sm := make(map[string]bool, 26)
	n := 0
	si, ei := 0, 0
	ts := ""
	for ; ei < len(s); ei++ {
		si = ei
		for ; ei < len(s) && s[ei] != ' '; ei++ {
		}
		if si == ei {
			// 跳过空格【不晓得案例上是否存在】
			continue
		}
		n++
		if n > len(pattern) {
			// 长度不一致
			return false
		}
		ts = s[si:ei]
		if ps[pattern[n-1]-'a'] == "" {
			if sm[ts] {
				// 字符串被使用过
				return false
			}
			// 新结果
			ps[pattern[n-1]-'a'] = ts
			sm[ts] = true
			continue
		}
		if ps[pattern[n-1]-'a'] != ts {
			return false
		}
		sm[ts] = true
	}
	// 最后判断长度是否一致
	return n == len(pattern)
}
