package one

func Respace(dictionary []string, sentence string) int {
	wm := make(map[string]bool, len(dictionary))
	for _, word := range dictionary {
		wm[word] = true
	}
	dp := make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = len(sentence) // 默认最大
		for j := i; j > 0; j-- {
			if !wm[sentence[j-1:i]] {
				dp[i] = min(dp[i], dp[j-1]+i-j+1)
				continue
			}
			dp[i] = min(dp[i], dp[j-1])
		}
	}
	return dp[len(sentence)]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
