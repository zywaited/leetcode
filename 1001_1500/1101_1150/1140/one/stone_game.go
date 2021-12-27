package one

// 动态规划
// dp[i][m] 代表剩下i堆石子(m为起始数)的最大值
// 下一次起始是max(m, x)
// 从1-2m中选出最大值
// dp[i][m] = max(sum[i]-dp[i-x][max(m, x)]) (1 <= x <= 2m)
// 这里如果不用sum用石子游戏那种先后手也是一样的, 也就是后手为sum-先手(最大值)
func StoneGame(piles []int) int {
	dp := make([][]int, len(piles)+1)
	for i := range dp {
		// 初始化
		dp[i] = make([]int, len(piles)+1)
	}
	// 剩余的石子总数
	sum := 0
	// num 剩余的石子堆数
	for num := 1; num <= len(piles); num++ {
		sum += piles[len(piles)-num]
		for m := 1; m <= len(piles); m++ {
			if m*2 >= num {
				// 这种情况全选
				dp[num][m] = sum
				continue
			}
			// 每种情况取最大值 x代表当前能一次性选择的石子堆数
			for x := 1; x <= m*2; x++ {
				dp[num][m] = max(dp[num][m], sum-dp[num-x][max(m, x)])
			}
		}
	}
	return dp[len(piles)][1]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
