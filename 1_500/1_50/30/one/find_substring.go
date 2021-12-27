package one

func FindSubstring(s string, words []string) []int {
	rs := make([]int, 0)
	// 所需的单词和数量
	need := make(map[string]int)
	for _, word := range words {
		need[word]++
	}
	nums := len(words)
	// 特别注意长度相等
	// 所以每次移动一个单词的距离即可，不需要每个字符都进行判断
	perWordLen := 0
	if nums > 0 {
		perWordLen = len(words[0])
	}
	// 最小长度判断
	sl := len(s)
	wl := perWordLen * nums
	if sl < wl {
		return rs
	}
	// 已经匹配的单词数
	find := make(map[string]int)
	match := 0
	left := 0
	start, end := 0, 0
	for ; left < sl; left++ {
		start = left
		// 最小长度判断
		for end = start + perWordLen; end <= sl && sl-start >= wl-perWordLen*match; end += perWordLen {
			word := s[start:end]
			// 未找到了或者已经过界
			// 长度固定，所以直接退出
			if _, ok := need[word]; !ok || find[word] > need[word]-1 {
				break
			}
			find[word]++
			match++
			if match == nums {
				// 已经找到
				rs = append(rs, left)
				break
			}
			start = end
		}
		// 重置查找
		if match > 0 {
			find = make(map[string]int)
			match = 0
		}
	}
	return rs
}
