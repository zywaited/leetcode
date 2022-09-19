package one

func SumPrefixScores(words []string) []int {
	wm := make(map[string]int)
	for _, word := range words {
		for i := range word {
			wm[word[:i+1]]++
		}
	}
	ans := make([]int, len(words))
	for i, word := range words {
		for j := range word {
			ans[i] += wm[word[:j+1]]
		}
	}
	return ans
}
