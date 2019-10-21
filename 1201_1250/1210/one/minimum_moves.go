package one

// BFS
func MinimumMoves(grid [][]int) int {
	n := len(grid[0])
	// 是否已经访问过了
	visited := make(map[[2]int]byte, n*n)
	// 节点列
	// 节点计算 (n*r1+c1, n*r2+c2)
	stack := make([][3]int, 0, n*n)
	stack = append(stack, [3]int{0, 1, 0})
	key := [2]int{}
	for len(stack) > 0 {
		pn := stack[0]
		stack = stack[1:]
		r1, c1 := pn[0]/n, pn[0]%n
		r2, c2 := pn[1]/n, pn[1]%n
		if r1 == r2 {
			// 水平
			// 向右
			key[0], key[1] = pn[1], n*r2+c2+1
			if c2+1 < n && grid[r2][c2+1] == 0 && visited[key] == 0 {
				if r2 == n-1 && c2+1 == n-1 {
					// 到了
					return pn[2] + 1
				}
				stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
				visited[key] = 1
			}
			if r2+1 < n && grid[r1+1][c1] == 0 && grid[r2+1][c2] == 0 {
				// 向下
				key[0], key[1] = n*(r1+1)+c1, n*(r2+1)+c2
				if visited[key] == 0 {
					if r2+1 == n-1 && c2 == n-1 {
						return pn[2] + 1
					}
					stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
					visited[key] = 1
				}
				// 旋转
				key[0], key[1] = pn[0], n*(r1+1)+c1
				if visited[key] == 0 {
					stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
					visited[key] = 1
				}
			}
			continue
		}
		// 向下
		key[0], key[1] = pn[1], n*(r2+1)+c2
		if r2+1 < n && grid[r2+1][c2] == 0 && visited[key] == 0 {
			stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
			visited[key] = 1
		}
		if c2+1 < n && grid[r1][c1+1] == 0 && grid[r2][c2+1] == 0 {
			// 向右
			key[0], key[1] = n*r1+c1+1, n*r2+c2+1
			if visited[key] == 0 {
				stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
				visited[key] = 1
			}
			// 旋转
			key[0], key[1] = pn[0], n*r1+c1+1
			if visited[key] == 0 {
				stack = append(stack, [3]int{key[0], key[1], pn[2] + 1})
				visited[key] = 1
			}
		}
	}
	return -1
}
