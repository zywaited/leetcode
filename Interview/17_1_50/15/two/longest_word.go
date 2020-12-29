package two

func LongestWord(words []string) string {
	wm := make(map[string]int, len(words))
	for _, word := range words {
		wm[word]++
	}
	cm := make(map[string]bool) // 字符串是否满足条件
	var bc func(int, int) bool
	bc = func(i, s int) bool {
		if s == len(words[i]) {
			return true
		}
		if vm, ok := cm[words[i][s:]]; ok {
			return vm
		}
		for e := len(words[i]); e > s; e-- {
			if wm[words[i][s:e]] == 0 || !bc(i, e) {
				continue
			}
			if s > 0 || e < len(words[i]) || wm[words[i]] > 1 {
				cm[words[i][s:e]] = true
			}
			if !cm[words[i][s:e]] {
				continue
			}
			return true
		}
		return false
	}
	ans := ""
	for i, word := range words {
		if bc(i, 0) && (len(ans) == 0 || len(word) > len(ans) || (len(word) == len(ans) && word < ans)) {
			ans = word
		}
	}
	return ans
}
