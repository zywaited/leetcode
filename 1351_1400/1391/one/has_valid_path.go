package one

import "container/list"

// bfs
func HasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	lsc := (m-1)*300 + n - 1
	queue := list.New()
	visited := make(map[int]byte, m*n)
	// 关联左边
	left := func(sc int) {
		i, j := sc/300, sc%300
		if j-1 >= 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) && visited[sc-1] == 0 {
			visited[sc-1] = 1
			queue.PushBack(sc - 1)
		}
	}
	// 关联右边
	right := func(sc int) {
		i, j := sc/300, sc%300
		if j+1 < n && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) && visited[sc+1] == 0 {
			visited[sc+1] = 1
			queue.PushBack(sc + 1)
		}
	}
	// 关联上边
	up := func(sc int) {
		i, j := sc/300, sc%300
		if i-1 >= 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) && visited[sc-300] == 0 {
			visited[sc-300] = 1
			queue.PushBack(sc - 300)
		}
	}
	// 关联下边
	down := func(sc int) {
		i, j := sc/300, sc%300
		if i+1 < m && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) && visited[sc+300] == 0 {
			visited[sc+300] = 1
			queue.PushBack(sc + 300)
		}
	}
	queue.PushBack(0)
	visited[0] = 1
	for queue.Len() > 0 {
		sc := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		switch grid[sc/300][sc%300] {
		case 1:
			left(sc)
			right(sc)
		case 2:
			up(sc)
			down(sc)
		case 3:
			left(sc)
			down(sc)
		case 4:
			right(sc)
			down(sc)
		case 5:
			left(sc)
			up(sc)
		case 6:
			right(sc)
			up(sc)
		}
		// 这里做个判断是否是到了结尾
		if visited[lsc] == 1 {
			return true
		}
	}
	return false
}
