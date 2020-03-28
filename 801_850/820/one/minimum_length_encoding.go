package one

import (
	"sort"
	"strings"
)

func MinimumLengthEncoding(words []string) int {
	// 只要一个单词是另一个单词的后缀就可以舍弃
	// 从后向前字典序排序
	sort.Slice(words, func(i, j int) bool {
		ki, kj := len(words[i])-1, len(words[j])-1
		for ki >= 0 && kj >= 0 {
			if words[i][ki] == words[j][kj] {
				ki--
				kj--
				continue
			}
			return words[i][ki] < words[j][kj]
		}
		return kj >= 0
	})
	ans := 0
	for i := range words {
		if i+1 == len(words) {
			ans += len(words[i]) + 1
			continue
		}
		// 当前是否会被下一个包含
		if strings.HasSuffix(words[i+1], words[i]) {
			continue
		}
		ans += len(words[i]) + 1
	}
	return ans
}
