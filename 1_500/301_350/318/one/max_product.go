package one

func MaxProduct(words []string) int {
	bl := []int{}
	for _, word := range words {
		b := 0
		for i := range word {
			b |= 1 << uint(word[i]-'a')
		}
		bl = append(bl, b)
	}
	ans := 0
	for i := 0; i < len(bl); i++ {
		for j := i + 1; j < len(bl); j++ {
			if bl[i]&bl[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
