package one

func MinPathCost(grid [][]int, moveCost [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cost := make([][]int, m)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = -1
		}
	}
	ans := -1
	// 队列只传输索引，不管当前的代价，代价用最新的即可
	var queue []int
	for j := 0; j < n; j++ {
		queue = append(queue, j)
		cost[0][j] = grid[0][j]
	}
	for len(queue) > 0 {
		r := queue[0] / 100
		c := queue[0] % 100
		queue = queue[1:]
		i := r + 1
		for j := 0; j < n; j++ {
			rc := moveCost[grid[r][c]][j] + grid[i][j]
			if cost[r][c] > 0 {
				rc += cost[r][c]
			}
			if cost[i][j] > 0 && cost[i][j] <= rc {
				continue
			}
			cost[i][j] = rc
			if i < m-1 {
				queue = append(queue, i*100+j)
				continue
			}
			if ans == -1 || cost[i][j] < ans {
				ans = cost[i][j]
			}
		}
	}
	return ans
}
