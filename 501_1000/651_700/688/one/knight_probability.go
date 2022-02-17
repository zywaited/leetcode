package one

var directs = [][]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func KnightProbability(n int, k int, row int, column int) float64 {
	curr := make([][]float64, n)
	for i := range curr {
		curr[i] = make([]float64, n)
		for j := range curr[i] {
			curr[i][j] = 1
		}
	}
	next := make([][]float64, n)
	for i := range next {
		next[i] = make([]float64, n)
		copy(next[i], curr[i])
	}
	for tk := 1; tk <= k; tk++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				next[i][j] = 0
				for _, direct := range directs {
					pi, pj := i+direct[0], j+direct[1]
					if pi >= 0 && pi < n && pj >= 0 && pj < n {
						next[i][j] += curr[pi][pj] / 8.0
					}
				}
			}
		}
		curr, next = next, curr
	}
	return curr[row][column]
}
