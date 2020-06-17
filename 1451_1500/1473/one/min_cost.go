package one

func MinCost(houses []int, cost [][]int, m int, n int, target int) int {
	// 原始状态
	// dp[i][j][k][l] 代表区间i-j, 第j个涂成k颜色分成l个街区的最小总花费
	// 后面发现以下状态可以替换上诉状态，(0-(i-1))已经包含了各种状态，因为i对于i-1只有相等或者不相等
	// dp[i][j][k] 代表区间0-i，第i个房子涂成j颜色分成k个街区的最小总花费
	dp := make([][][]int, m)
	// 初始化数组
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	// 对第一个处理
	if houses[0] == 0 {
		// 每种颜色都有花费
		for j := 0; j < n; j++ {
			dp[0][j][0] = cost[0][j]
		}
	} else {
		// 颜色固定则不需要花费
		dp[0][houses[0]-1][0] = 0
	}
	for i := 1; i < len(houses); i++ {
		if houses[i] == 0 {
			for j := 0; j < n; j++ {
				for k := 0; k < target; k++ {
					// 这里还要对颜色进行遍历
					// 因为dp[i-1]也有不同颜色对应的k个街区
					for jj := 0; jj < n; jj++ {
						// 不可能完成
						if (jj == j && dp[i-1][jj][k] == -1) || (jj != j && (k == 0 || dp[i-1][jj][k-1] == -1)) {
							continue
						}
						if jj == j {
							// 相同颜色药达到K个街区
							if dp[i][j][k] == -1 || dp[i-1][jj][k]+cost[i][j] < dp[i][j][k] {
								dp[i][j][k] = dp[i-1][jj][k] + cost[i][j]
							}
							continue
						}
						if dp[i][j][k] == -1 || dp[i-1][jj][k-1]+cost[i][j] < dp[i][j][k] {
							dp[i][j][k] = dp[i-1][jj][k-1] + cost[i][j]
						}
					}
				}
			}
			continue
		}
		// 与上述逻辑一致, 只是固定了当前颜色
		j := houses[i] - 1
		for k := 0; k < target; k++ {
			for jj := 0; jj < n; jj++ {
				if (jj == j && dp[i-1][jj][k] == -1) || (jj != j && (k == 0 || dp[i-1][jj][k-1] == -1)) {
					continue
				}
				if jj == j {
					if dp[i][j][k] == -1 || dp[i-1][jj][k] < dp[i][j][k] {
						dp[i][j][k] = dp[i-1][jj][k]
					}
					continue
				}
				if dp[i][j][k] == -1 || dp[i-1][jj][k-1] < dp[i][j][k] {
					dp[i][j][k] = dp[i-1][jj][k-1]
				}
			}
		}
	}
	ans := -1
	for j := 0; j < n; j++ {
		if dp[m-1][j][target-1] > -1 && (ans == -1 || dp[m-1][j][target-1] < ans) {
			ans = dp[m-1][j][target-1]
		}
	}
	return ans
}
