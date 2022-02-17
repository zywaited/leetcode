package two

var directs = [][]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func KnightProbability(n int, k int, row int, column int) float64 {
	curr := make([][]float64, n)
	next := make([][]float64, n)
	for i := range curr {
		curr[i] = make([]float64, n)
		next[i] = make([]float64, n)
	}
	curr[row][column] = 1
	ans := float64(1)
	for tk := 1; tk <= k; tk++ {
		ans = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				next[i][j] = 0
				for _, direct := range directs {
					pi, pj := i+direct[0], j+direct[1]
					if pi >= 0 && pi < n && pj >= 0 && pj < n {
						next[i][j] += curr[pi][pj] / 8.0
					}
				}
				if tk == k {
					ans += next[i][j]
				}
			}
		}
		curr, next = next, curr
	}
	return ans
}
