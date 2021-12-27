package one

import "math"

// 空间复杂度: O(N)(N为三角形行数)
func MinimumTotal(triangle [][]int) int {
	ans := 0                         // 提前声明为了少循环一次
	dp := make([]int, len(triangle)) // 保留上一层的最小值
	tmp := 0                         // 临时保留j的数据
	for i := range triangle {
		ans = math.MaxInt32 // 每次重置为最大
		for j := range triangle[i] {
			// 为了保持一致
			if j == 0 {
				tmp = dp[j]
			}
			if j == len(triangle[i])-1 {
				dp[j] = tmp
			}
			dp[j], tmp = min(tmp, dp[j])+triangle[i][j], dp[j]
			ans = min(ans, dp[j])
		}
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
