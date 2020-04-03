package one

// 拆分成两种情况(这样不用考虑循环数组)
// 1: 选择第一个，那么最后一个就选不了
// 2: 不选择第一个，那么最后一个就可选
func MaxSizeSlices(slices []int) int {
	return max(maxSizeSlice(slices[:len(slices)-1]), maxSizeSlice(slices[1:]))
}

func maxSizeSlice(slices []int) int {
	// dp[i][j] 代表前i块披萨取j次的最大值
	dp := make([][]int, len(slices)+1)
	mj := (len(slices) + 1) / 3 // j的最大值
	for i := 0; i <= len(slices); i++ {
		dp[i] = make([]int, mj+1)
		if i == 0 { // 仅仅为了初始化
			continue
		}
		for j := 1; j <= mj; j++ {
			// 当前不取
			dp[i][j] = dp[i-1][j]
			// 隔着取
			if i-2 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-2][j-1]+slices[i-1])
				continue
			}
			// 只能取当前一个数字
			dp[i][j] = max(dp[i][j], slices[i-1])
		}
	}
	return dp[len(slices)][mj]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
