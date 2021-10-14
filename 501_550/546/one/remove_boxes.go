package one

func RemoveBoxes(boxes []int) int {
	// dp[i][j][k] i-j区间连续移除k个j的最大分数s
	dp := [101][101][101]int{}
	for j := 1; j <= len(boxes); j++ {
		for i := j; i > 0; i-- {
			if i == j {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				continue
			}
			dp[i][j][1] = dp[i][j-1][0] + 1
			dp[i][j][0] = dp[i][j][1]
			for l := i; l < j; l++ {
				for k := l - i + 1; k > 0; k-- {
					if dp[i][l][k] == 0 {
						continue
					}
					if boxes[l-1] == boxes[j-1] {
						dp[i][j][k+1] = max(dp[i][j][k+1], dp[i][l][k]+dp[l+1][j-1][0]+k*2+1)
						dp[i][j][0] = max(dp[i][j][0], dp[i][j][k+1])
					}
				}
			}
		}
	}
	return dp[1][len(boxes)][0]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
