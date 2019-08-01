package two

// 与第一种解法是类似的
// 不过重复利用找到的单词
func FindSubstring(s string, words []string) []int {
	rs := make([]int, 0)
	// 所需的单词和数量
	nums := len(words)
	need := make(map[string]int, nums)
	for _, word := range words {
		need[word]++
	}
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
	find := make(map[string]int, nums)
	match := 0
	start, mid := 0, 0
	// 循环跟单词长度有关，因为每个单词长度是相等的
	// 所以可以包含所有的情况
	for i := 0; i < perWordLen; i++ {
		start, mid = i, i
		// 判断是否需要重置数据
		if match > 0 {
			find = make(map[string]int, nums)
			match = 0
		}
		// 是否长度足够
		for sl-start >= wl {
			word := s[mid : mid+perWordLen]
			// 如果不存在
			if _, ok := need[word]; !ok {
				// 跳过单词，更新起点
				start += perWordLen
				mid = start
				// 重置
				find = make(map[string]int, nums)
				match = 0
				continue
			}
			// 过界
			if find[word] > need[word]-1 {
				// 去除第一个
				find[s[start:start+perWordLen]]--
				match--
				// 更新起点不跳过单词
				start += perWordLen
				continue
			}
			find[word]++
			match++
			// 找到了
			if match == nums {
				rs = append(rs, start)
				// 去除第一个
				find[s[start:start+perWordLen]]--
				match--
				// 更新起点
				start += perWordLen
			}
			// 下一个单词
			mid += perWordLen
		}
	}
	return rs
}
