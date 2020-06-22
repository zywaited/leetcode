package one

func PatternMatching(pattern string, value string) bool {
	return isMatch(pattern, value, nil, nil)
}

// 如果string赋值过多，转为[]byte性能高一点
func isMatch(pattern, value string, as, bs []byte) bool {
	// 没有需要匹配的
	if pattern == "" && value == "" {
		return !eq(as, bs)
	}
	if pattern == "" && value != "" {
		return false
	}
	var s []byte
	if pattern[0] == 'a' {
		s = as
	} else {
		s = bs
	}
	if s != nil {
		if len(value) < len(s) || value[0:len(s)] != string(s) {
			return false
		}
		return isMatch(pattern[1:], value[len(s):], as, bs)
	}
	for i := 0; i <= len(value); i++ {
		if pattern[0] == 'a' {
			if !eq(bs, []byte(value[0:i])) && isMatch(pattern[1:], value[i:], []byte(value[0:i]), bs) {
				return true
			}
			continue
		}
		if !eq(as, []byte(value[0:i])) && isMatch(pattern[1:], value[i:], as, []byte(value[0:i])) {
			return true
		}
	}
	if len(value) > 0 {
		return false
	}
	for i := 1; i < len(pattern); i++ {
		if pattern[i] != pattern[0] {
			return false
		}
	}
	return string(as) != string(bs)
}

func eq(as, bs []byte) bool {
	if as != nil && bs != nil {
		return string(as) == string(bs)
	}
	return false
}
