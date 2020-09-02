package one

func IsNumber(s string) bool {
	lastIndex := 0
	endIndex := len(s) - 1
	for ; lastIndex < endIndex && s[lastIndex] == ' '; lastIndex++ {
	}
	for ; endIndex > lastIndex && s[endIndex] == ' '; endIndex-- {
	}
	// 以E为分割
	notN := false
	foundN := false
	index := lastIndex
	for ; index <= endIndex && (s[index] != 'e' && s[index] != 'E'); index++ {
		if s[index] >= '0' && s[index] <= '9' {
			foundN = true
			continue
		}
		if index == lastIndex {
			if s[index] == '+' || s[index] == '-' {
				continue
			}
			if s[index] != '.' {
				return false
			}
		}
		if s[index] != '.' || notN {
			return false
		}
		notN = true
	}
	if index == lastIndex {
		return false
	}
	if !foundN {
		return false
	}
	if index == endIndex+1 {
		return true
	}
	index++
	foundN = false
	lastIndex = index
	for ; index <= endIndex; index++ {
		if s[index] >= '0' && s[index] <= '9' {
			foundN = true
			continue
		}
		if index == lastIndex {
			if s[index] != '+' && s[index] != '-' {
				return false
			}
			continue
		}
		return false
	}
	return index > lastIndex && foundN
}
