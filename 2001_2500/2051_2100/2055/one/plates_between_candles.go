package one

func PlatesBetweenCandles(s string, queries [][]int) []int {
	type candle struct {
		left  int
		right int
	}
	candles := make([]candle, len(s))
	sums := make([]int, len(s)+1)
	ci := 0
	for i := 1; i <= len(s); i++ {
		sums[i] = sums[i-1]
		switch s[i-1] {
		case '*':
			candles[i-1].left = ci
			sums[i]++
		case '|':
			candles[i-1].left = i
			for j := i - 1; j >= 0 && candles[j].right == 0; j-- {
				candles[j].right = i
			}
			ci = i
		}
	}
	ans := make([]int, len(queries))
	for index, query := range queries {
		if candles[query[0]].right == 0 || candles[query[1]].left == 0 || candles[query[1]].left <= candles[query[0]].right {
			continue
		}
		ans[index] = sums[candles[query[1]].left] - sums[candles[query[0]].right]
	}
	return ans
}
