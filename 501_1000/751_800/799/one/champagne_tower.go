package one

func ChampagneTower(poured int, query_row int, query_glass int) float64 {
	curr := []float64{float64(poured)}
	next := []float64{}
	c := float64(0)
	for row := 1; row <= query_row; row++ {
		next = next[:0]
		for glass := 0; glass <= row; glass++ {
			c = 0
			if glass > 0 && 1 < curr[glass-1] {
				c += (curr[glass-1] - 1) * 0.5
			}
			if glass < row && 1 < curr[glass] {
				c += (curr[glass] - 1) * 0.5
			}
			next = append(next, c)
		}
		curr, next = next, curr
	}
	if 1 <= curr[query_glass] {
		return 1.0
	}
	return curr[query_glass]
}
