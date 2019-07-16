package one

func LengthOfLongestSubstring(s string) int {
	index := make(map[rune]int)
	maxSubLength := 0
	last := 0
	for i, c := range s {
		if e, ok := index[c]; ok {
			last = Max(last, e)
		}

		maxSubLength = Max(maxSubLength, i-last+1)
		index[c] = i + 1
	}

	return maxSubLength
}

func Max(f, s int) int {
	if f > s {
		return f
	}

	return s
}
