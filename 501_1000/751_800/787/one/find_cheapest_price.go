package one

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	points := make([][]int, n)
	for index, flight := range flights {
		points[flight[0]] = append(points[flight[0]], index)
	}
	type node struct {
		point int
		step  int
		cost  int
	}
	ans := -1
	visited := make([]int, n)
	for index := range visited {
		visited[index] = -1
	}
	visited[src] = 0
	queue := make([]node, 0, n)
	queue = append(queue, node{point: src})
	for len(queue) > 0 {
		cn := queue[0]
		queue = queue[1:]
		if cn.step > k {
			continue
		}
		cost := 0
		for _, index := range points[cn.point] {
			cost = cn.cost + flights[index][2]
			if (visited[flights[index][1]] == -1 || cost < visited[flights[index][1]]) && (ans == -1 || cost < ans) {
				if flights[index][1] == dst {
					ans = cn.cost + flights[index][2]
					continue
				}
				visited[flights[index][1]] = cost
				queue = append(queue, node{point: flights[index][1], step: cn.step + 1, cost: cost})
			}
		}
	}
	return ans
}
