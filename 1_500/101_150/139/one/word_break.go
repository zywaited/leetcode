package one

func WordBreak(s string, wordDict []string) bool {
	if len(s) < 1 {
		return false
	}
	if len(wordDict) < 1 {
		return false
	}
	words := make(map[string]byte)
	for _, word := range wordDict {
		words[word] = 1
	}
	return dp(s, 0, words, make(map[int]bool))
}

func dp(s string, start int, words map[string]byte, mp map[int]bool) bool {
	// 到结尾了
	if start == len(s) {
		return true
	}
	// 计算过了
	if _, ok := mp[start]; ok {
		return mp[start]
	}
	for end := start + 1; end <= len(s); end++ {
		// 是否存在
		if words[s[start:end]] < 1 {
			continue
		}
		// 下一个单词是否存在
		if !dp(s, end, words, mp) {
			continue
		}
		mp[start] = true
		return true
	}
	mp[start] = false
	return false
}
