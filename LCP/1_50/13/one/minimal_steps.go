package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func MinimalSteps(maze []string) int {
	// 1: 计算S、T、O、M到各点之间最短距离
	// 2: 计算M->M的最短距离
	// 3: 计算最小步数(M->T)
	on := make([][2]int, 0)
	mn := make([][2]int, 0)
	sx, sy := 0, 0 // S
	tx, ty := 0, 0 // T
	for i := range maze {
		for j := range maze[i] {
			switch maze[i][j] {
			case 'S':
				sx, sy = i, j
			case 'T':
				tx, ty = i, j
			case 'O':
				on = append(on, [2]int{i, j})
			case 'M':
				mn = append(mn, [2]int{i, j})
			}
		}
	}

	// S到各个点的最短距离
	sds := bfs(sx, sy, maze)
	if len(mn) == 0 {
		return sds[tx][ty] // 没有机关直接看起点到终点
	}

	// M到各点的距离
	mds := make([][][]int, len(mn))
	for i := 0; i < len(mn); i++ {
		mds[i] = bfs(mn[i][0], mn[i][1], maze)
	}

	// M->O->S的最短距离
	smds := make([]int, len(mn))
	for i := 0; i < len(mn); i++ {
		smds[i] = -1
	}
	// M->O->M的最短距离
	mmds := make([][]int, len(mn))
	for i := 0; i < len(mn); i++ {
		mmds[i] = make([]int, len(mn))
		for j := 0; j < len(mn); j++ {
			mmds[i][j] = -1
		}
	}
	for i := 0; i < len(mmds); i++ {
		for j := range on {
			if sds[on[j][0]][on[j][1]] == -1 || mds[i][on[j][0]][on[j][1]] == -1 {
				continue
			}
			if smds[i] == -1 || sds[on[j][0]][on[j][1]]+mds[i][on[j][0]][on[j][1]] < smds[i] {
				smds[i] = sds[on[j][0]][on[j][1]] + mds[i][on[j][0]][on[j][1]]
			}
		}
		for j := i; j >= 0; j-- {
			min := -1
			for k := range on {
				if mds[j][on[k][0]][on[k][1]] == -1 || mds[i][on[k][0]][on[k][1]] == -1 {
					// 到不了
					continue
				}
				if min == -1 || mds[j][on[k][0]][on[k][1]]+mds[i][on[k][0]][on[k][1]] < min {
					min = mds[j][on[k][0]][on[k][1]] + mds[i][on[k][0]][on[k][1]]
				}
			}
			mmds[j][i] = min
			mmds[i][j] = min
		}
	}

	// 1: 判断M是否能到S
	// 2: 判断M是否能到T
	for i := 0; i < len(mn); i++ {
		if smds[i] == -1 || mds[i][tx][ty] == -1 {
			return -1
		}
	}

	// 以某个M结束的最小步数
	dp := make([][]int, 1<<uint(len(mn)))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(mn))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	// 以某个M为起点
	for i := 0; i < len(mn); i++ {
		dp[1<<uint(i)][i] = smds[i]
	}
	for m := 1; m < len(dp); m++ {
		for i := 0; i < len(mn); i++ {
			if m&(1<<uint(i)) == 0 {
				// 无关的M
				continue
			}
			for j := 0; j < len(mn); j++ {
				if m&(1<<uint(j)) > 0 {
					// 已经计算过当前M了
					continue
				}
				n := m | (1 << uint(j))
				if dp[n][j] == -1 || dp[m][i]+mmds[i][j] < dp[n][j] {
					dp[n][j] = dp[m][i] + mmds[i][j]
				}
			}
		}
	}

	// 最后遍历获取结果(需要加上S->M->T的距离)
	step := -1
	for i, s := range dp[len(dp)-1] {
		if step == -1 || s+mds[i][tx][ty] < step {
			step = s + mds[i][tx][ty]
		}
	}
	return step
}

// 返回x,y点到各个点的距离
func bfs(x, y int, maze []string) [][]int {
	m, n := len(maze), len(maze[0])
	distance := make([][]int, m)
	for i := range distance {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			distance[i][j] = -1
		}
	}
	distance[x][y] = 0
	queue := make([][2]int, 0, m*n)
	queue = append(queue, [2]int{x, y})
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n || distance[ni][nj] != -1 || maze[ni][nj] == '#' {
				continue
			}
			distance[ni][nj] = distance[i][j] + 1
			queue = append(queue, [2]int{ni, nj})
		}
	}
	return distance
}
