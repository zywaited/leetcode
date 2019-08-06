package one

func CheckInclusion(s1 string, s2 string) bool {
	// 需要的字符集合
	nums := len(s1)
	need := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	find := make(map[byte]int)
	match := 0
	start := 0
	for i := 0; i < len(s2); i++ {
		// 不存在
		if need[s2[i]] == 0 {
			if match > 0 {
				match = 0
				find = make(map[byte]int)
			}
			// 起点是下一位
			start = i + 1
			continue
		}
		// 多了
		for find[s2[i]] > need[s2[i]]-1 {
			find[s2[start]]--
			match--
			start++
		}
		find[s2[i]]++
		match++
		// 已经找到了
		if match == nums {
			return true
		}
	}
	return false
}
