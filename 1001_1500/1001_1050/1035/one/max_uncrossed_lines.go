package one

// 动态规则
// dp[i][j] A中前i个数字与B中前j个数字所绘画的最大数
// 1: 如果A[i] == B[j] 则可以多一条数据 dp[i][j] = dp[i-1][j-1] + 1
// 2: 如果A[i] != B[j] dp[i][j] = max(dp[i-1][y], dp[i][y-1])
func MaxUncrossedLines(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := range A {
		for j := range B {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
		}
	}
	return dp[len(A)][len(B)]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
