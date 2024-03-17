package two

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	pts := make([]int, m)
	cts := make([]int, m)
	mt := 0
	for j := n - 2; j >= 0; j-- {
		for i := 0; i < m; i++ {
			cts[i] = 0
			for _, direct := range []int{-1, 0, 1} {
				ni := i + direct
				if ni < 0 || m <= ni || grid[ni][j+1] <= grid[i][j] {
					continue
				}
				if cts[i] < pts[ni]+1 {
					cts[i] = pts[ni] + 1
				}
			}
			if j == 0 && mt < cts[i] {
				mt = cts[i]
			}
		}
		pts, cts = cts, pts
	}
	return mt
}
