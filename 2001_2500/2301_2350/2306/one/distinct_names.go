package one

func DistinctNames(ideas []string) int64 {
	names := map[string]bool{}
	for _, idea := range ideas {
		names[idea] = true
	}
	nodes := [26][26]int64{}
	result := int64(0)
	for _, idea := range ideas {
		for chatIndex := byte(0); chatIndex < 26; chatIndex++ {
			if !names[string(chatIndex+'a')+idea[1:]] {
				result += nodes[chatIndex][idea[0]-'a'] * 2
				nodes[idea[0]-'a'][chatIndex]++
			}
		}
	}
	return result
}
