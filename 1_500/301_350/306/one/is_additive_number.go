package one

func IsAdditiveNumber(num string) bool {
	var f []byte
	var s []byte
	for i := range num {
		f = append(f, num[i]-'0')
		if len(f) > 1 && f[0] == 0 {
			break
		}
		s = s[:0]
		for j := i + 1; j < len(num); j++ {
			s = append(s, num[j]-'0')
			if len(s) > 1 && s[0] == 0 {
				break
			}
			if dfs(f, s, num[j+1:]) {
				return true
			}
		}
	}
	return false
}

func dfs(f, s []byte, num string) bool {
	fs := sum(f, s)
	if len(num) < len(fs) {
		return false
	}
	for i := range fs {
		if fs[i] != num[i]-'0' {
			return false
		}
	}
	if len(num) == len(fs) {
		return true
	}
	return dfs(s, fs, num[len(fs):])
}

func sum(f, s []byte) []byte {
	fi := len(f)
	si := len(s)
	l := len(f)
	if l < len(s) {
		l = len(s)
	}
	fs := make([]byte, l+1)
	c := byte(0)
	for i := len(fs) - 1; i >= 0; i-- {
		fi--
		si--
		if fi >= 0 && si >= 0 {
			fs[i] = (f[fi] + s[si] + c) % 10
			c = (f[fi] + s[si] + c) / 10
			continue
		}
		if fi >= 0 {
			fs[i] = (f[fi] + c) % 10
			c = (f[fi] + c) / 10
			continue
		}
		if si >= 0 {
			fs[i] = (s[si] + c) % 10
			c = (s[si] + c) / 10
			continue
		}
		fs[i] = c
		c = 0
	}
	if fs[0] == 0 {
		fs = fs[1:]
	}
	return fs
}
