package one

func IsNumber(s string) bool {
	num := false
	point := false
	i := 0
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = true
			continue
		}
		if s[i] == '-' || s[i] == '+' {
			if i > 0 {
				return false
			}
			continue
		}
		if s[i] == '.' {
			if point {
				return false
			}
			point = true
			continue
		}
		if s[i] == 'e' || s[i] == 'E' {
			break
		}
		return false
	}
	if !num {
		return false
	}
	if i >= len(s) {
		return true
	}
	i++
	// 判断是否是整数
	j := i
	num = false
	for ; j < len(s); j++ {
		if s[j] >= '0' && s[j] <= '9' {
			num = true
			continue
		}
		if s[j] == '-' || s[j] == '+' {
			if j > i {
				return false
			}
			continue
		}
		return false
	}
	return num
}
