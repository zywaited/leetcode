package one

func MaxScore(cardPoints []int, k int) int {
	sum := 0
	n := len(cardPoints) - k
	i := 0
	for ; i < n; i++ {
		sum += cardPoints[i]
	}
	sw := sum
	mw := sum
	for ; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		sw += cardPoints[i] - cardPoints[i-n]
		mw = min(mw, sw)
	}
	return sum - mw
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
