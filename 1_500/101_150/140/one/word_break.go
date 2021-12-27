package one

func WordBreak(s string, wordDict []string) []string {
	words := make(map[string]byte)
	for _, word := range wordDict {
		words[word] = 1
	}
	// 这里用了string，当然用[]byte理论上会好点
	mp := make(map[int][]string)
	dp(s, 0, nil, words, mp)
	return mp[0]
}

func dp(s string, start int, cs []byte, words map[string]byte, mp map[int][]string) bool {
	sl := len(s)
	// 处理过
	if _, ok := mp[start]; ok {
		return len(mp[start]) > 0
	}
	word := ""
	for end := start + 1; end <= sl; end++ {
		word = s[start:end]
		if words[word] < 1 {
			continue
		}
		cs = append(cs, word...)
		// 下一波
		if end == sl || dp(s, end, nil, words, mp) {
			// 找到了
			// 结尾
			if end == sl {
				mp[start] = append(mp[start], string(cs))
				return true
			}
			// 中间
			cs = append(cs, ' ')
			for _, ms := range mp[end] {
				mp[start] = append(mp[start], string(cs)+ms)
			}
			cs = cs[:len(cs)-len(word)-1]
			continue
		}
		cs = cs[:len(cs)-len(word)]
	}
	if _, ok := mp[start]; !ok {
		mp[start] = nil
	}
	return len(mp[start]) > 0
}
