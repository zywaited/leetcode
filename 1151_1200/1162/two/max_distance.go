package two

// 上一个解法多源是用的陆地，因为用海洋不知道哪个是最近的陆地
// 这次用海洋，不过不是bfs，是动态规划
// dp[i][j]的值就是4个方向最近陆地的距离
// dp[i][j] = min(dp[i-1][j], dp[i][j+1], dp[i+1][j], dp[i][j-1])+1
func MaxDistance(grid [][]int) int {
	// 也可以用上一个方法那种score的方式改成map
	dp := make([][]int, len(grid))
	// 因为计算dp[i][j]时dp[i][j+1]和dp[i+1][j]是没有计算的
	// 所以需要先计算左边和上边的数据，再计算下边和右边
	i, j := 0, 0
	for i = range grid {
		for j = range grid[i] {
			if grid[i][j] == 1 {
				continue
			}
			if dp[i] == nil {
				dp[i] = make([]int, len(grid))
			}
			if j-1 >= 0 {
				if grid[i][j-1] == 1 {
					dp[i][j] = 1
					continue
				}
				if dp[i][j-1] > 0 && (dp[i][j] == 0 || dp[i][j-1]+1 < dp[i][j]) {
					dp[i][j] = dp[i][j-1] + 1
				}
			}
			if i-1 >= 0 {
				if grid[i-1][j] == 1 {
					dp[i][j] = 1
					continue
				}
				if dp[i-1][j] > 0 && (dp[i][j] == 0 || dp[i-1][j]+1 < dp[i][j]) {
					dp[i][j] = dp[i-1][j] + 1
				}
			}
		}
	}
	ans := -1
	for i = len(grid) - 1; i >= 0; i-- {
		for j = len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if j+1 < len(grid[i]) {
				if grid[i][j+1] == 1 {
					dp[i][j] = 1
					continue
				}
				if dp[i][j+1] > 0 && (dp[i][j] == 0 || dp[i][j+1]+1 < dp[i][j]) {
					dp[i][j] = dp[i][j+1] + 1
				}
			}
			if i+1 < len(grid) {
				if grid[i+1][j] == 1 {
					dp[i][j] = 1
					continue
				}
				if dp[i+1][j] > 0 && (dp[i][j] == 0 || dp[i+1][j]+1 < dp[i][j]) {
					dp[i][j] = dp[i+1][j] + 1
				}
			}
			if dp[i][j] > 0 && dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
