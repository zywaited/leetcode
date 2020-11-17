package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func AllCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	queue := make([][]int, 0, R*C)
	visited := make(map[int]bool, R*C)

	queue = append(queue, []int{r0, c0})
	visited[r0*100+c0] = true
	queueIndex := 0
	for queueIndex < len(queue) {
		rc := queue[queueIndex]
		queueIndex++
		for _, direct := range directs {
			r1, c1 := rc[0]+direct[0], rc[1]+direct[1]
			if r1 >= 0 && r1 < R && c1 >= 0 && c1 < C && !visited[r1*100+c1] {
				queue = append(queue, []int{r1, c1})
				visited[r1*100+c1] = true
			}
		}
	}
	return queue
}
