package one

const mod = 1e9 + 7

func DieSimulator(n int, rollMax []int) int {
	dp := make([][6][15]int, n)
	// 为了不每次循环15次获取当前i数字的总和加个数组保存
	nums := [6]int{}
	i := 0
	for ; i < 6; i++ {
		// 第一次掷骰子连续i的总数
		dp[0][i][0] = 1
		nums[i] = 1
	}
	j, k, l := 0, 0, 0
	for i = 1; i < n; i++ {
		for j = 0; j < 6; j++ {
			for k = 0; k < 6; k++ {
				if j != k {
					// 不等就只有连续1个j
					dp[i][j][0] = (dp[i][j][0] + nums[k]) % mod
					continue
				}
				// 相等就判断满足条件的(至少连续2次)
				for l = 1; l < rollMax[j]; l++ {
					dp[i][j][l] = dp[i-1][j][l-1]
				}
			}
		}
		// 重新计算总和
		for j = 0; j < 6; j++ {
			nums[j] = 0
			for l = 0; l < rollMax[j]; l++ {
				nums[j] = (nums[j] + dp[i][j][l]) % mod
			}
		}
	}
	total := 0
	for i = range nums {
		total = (total + nums[i]) % mod
	}
	return total
}
