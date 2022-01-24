package one

func SecondMinimum(n int, edges [][]int, time int, change int) int {
	points := make([][]int, n+1)
	for _, edge := range edges {
		points[edge[0]] = append(points[edge[0]], edge[1])
		points[edge[1]] = append(points[edge[1]], edge[0])
	}
	visited := make([][]int, n+1)
	for i := 2; i < len(visited); i++ {
		visited[i] = []int{-1, -1}
	}
	visited[1] = []int{0, -1}
	type pair struct {
		index int
		time  int
	}
	queue := []pair{{index: 1}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.time/change%2 > 0 {
			// 等红灯
			node.time += change - node.time%change
		}
		for _, nextIndex := range points[node.index] {
			next := node.time + time
			// 要求严格小于
			if visited[nextIndex][0] == next {
				continue
			}
			if visited[nextIndex][0] == -1 || next < visited[nextIndex][0] {
				visited[nextIndex][1] = visited[nextIndex][0]
				visited[nextIndex][0] = next
				queue = append(queue, pair{index: nextIndex, time: next})
				continue
			}
			if visited[nextIndex][1] == -1 || next < visited[nextIndex][1] {
				visited[nextIndex][1] = next
				queue = append(queue, pair{index: nextIndex, time: next})
			}
		}
	}
	return visited[n][1]
}
