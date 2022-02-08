package one

var directs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func GridIllumination(n int, lamps [][]int, queries [][]int) []int {
	rows := make(map[int]int, len(lamps))
	cols := make(map[int]int, len(lamps))
	diagonal := make(map[int]int, len(lamps))
	reverseDiagonal := make(map[int]int, len(lamps))
	usedLamps := make(map[[2]int]bool, len(lamps))
	for _, lamp := range lamps {
		key := [2]int{lamp[0], lamp[1]}
		if !usedLamps[key] {
			rows[lamp[0]]++
			cols[lamp[1]]++
			diagonal[lamp[0]-lamp[1]]++
			reverseDiagonal[lamp[0]+lamp[1]]++
			usedLamps[key] = true
		}
	}
	ans := make([]int, len(queries))
	for index, query := range queries {
		if rows[query[0]] > 0 || cols[query[1]] > 0 || diagonal[query[0]-query[1]] > 0 || reverseDiagonal[query[0]+query[1]] > 0 {
			ans[index] = 1
		}
		for _, direct := range directs {
			r, c := query[0]+direct[0], query[1]+direct[1]
			key := [2]int{r, c}
			if r < 0 || r >= n || c < 0 || c >= n || !usedLamps[key] {
				continue
			}
			usedLamps[key] = false
			rows[r]--
			cols[c]--
			diagonal[r-c]--
			reverseDiagonal[r+c]--
		}
	}
	return ans
}
