package two

// 动态规划
func MinimumMoves(grid [][]int) int {
	// 不沿用解法一个节点计算规则
	// 不然计算太麻烦了，每次都要计算
	// 多搞一个为存储: 0 水平 1 竖直
	n := len(grid[0])
	dp := make(map[[3]int]int, n*n)
	key := [3]int{0, 0, 0}
	ms := 0
	dp[key] = 1
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				continue
			}
			ms = 0
			// 以蛇尾当做参考点
			if c+1 < n && grid[r][c+1] == 0 {
				// 可水平
				// 右移
				key[0], key[1], key[2] = r, c-1, 0
				if dp[key] != 0 {
					ms = dp[key]
				}
				// 下移
				key[0], key[1], key[2] = r-1, c, 0
				if dp[key] != 0 {
					ms = min(ms, dp[key])
				}
				if ms != 0 {
					key[0], key[1], key[2] = r, c, 0
					ms++
					dp[key] = ms
				}
			}
			if r+1 < n && grid[r+1][c] == 0 {
				ms = 0
				if c+1 < n && grid[r+1][c+1] == 0 {
					// 从水平旋转
					key[0], key[1], key[2] = r, c, 0
					ms = dp[key]
				}
				// 可竖直
				// 下移
				key[0], key[1], key[2] = r-1, c, 1
				if dp[key] != 0 {
					ms = min(ms, dp[key])
				}
				// 右移
				key[0], key[1], key[2] = r, c-1, 1
				if dp[key] != 0 {
					ms = min(ms, dp[key])
				}
				if ms != 0 {
					key[0], key[1], key[2] = r, c, 1
					ms++
					dp[key] = ms
					// 重新更新下水平
					if c+1 < n && grid[r+1][c+1] == 0 && grid[r][c+1] == 0 {
						key[2] = 0
						dp[key] = min(dp[key], ms+1)
					}
				}
			}
		}
	}
	key[0], key[1], key[2] = n-1, n-2, 0
	return dp[key] - 1
}

func min(f, s int) int {
	// 调用该方法一定不会全为0
	if f == 0 {
		return s
	}
	if s == 0 {
		return f
	}
	if f <= s {
		return f
	}
	return s
}
