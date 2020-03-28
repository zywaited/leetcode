package three

func MinimumLengthEncoding(words []string) int {
	// 保存所有的后缀
	dm := make(map[string]byte)
	for _, word := range words {
		dm[word] = 1
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			// 删除已经存在的所有后缀
			delete(dm, word[i:])
		}
	}
	ans := 0
	for word := range dm {
		ans += len(word) + 1
	}
	return ans
}
