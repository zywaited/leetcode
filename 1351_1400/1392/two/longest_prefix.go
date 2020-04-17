package two

func LongestPrefix(s string) string {
	// next这里保留长度(就不会最后再处理一次)
	next := make([]int, len(s))
	ci := 0
	for i := 1; i < len(s); i++ {
		if s[ci] == s[i] {
			ci++
			next[i] = ci
			continue
		}
		if ci > 0 {
			// 可以还原继续匹配
			ci = next[ci-1]
			i--
			continue
		}
		// 只能重头匹配
		next[i] = 0
	}
	return s[:next[len(s)-1]]
}
