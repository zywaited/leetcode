package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func MinimumEffortPath(heights [][]int) int {
	r := len(heights)
	c := len(heights[0])
	ah := make([][]int, r)
	for i := range ah {
		ah[i] = make([]int, c)
		for j := range ah[i] {
			ah[i][j] = -1
		}
	}
	ah[0][0] = 0
	q := []int{0}
	for len(q) > 0 {
		i, j := q[0]/100, q[0]%100
		q = q[1:]
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || ni >= r || nj < 0 || nj >= c {
				continue
			}
			h := heights[i][j] - heights[ni][nj]
			if h < 0 {
				h = -h
			}
			if ah[i][j] > h {
				h = ah[i][j]
			}
			if ah[ni][nj] == -1 || h < ah[ni][nj] {
				ah[ni][nj] = h
				q = append(q, ni*100+nj)
			}
		}
	}
	return ah[r-1][c-1]
}
