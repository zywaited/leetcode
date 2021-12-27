package one

import "strings"

func StringMatching(words []string) []string {
	ans := make([]string, 0, len(words))
	for i := range words {
		for j := range words {
			if i != j && strings.Index(words[j], words[i]) >= 0 {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}
