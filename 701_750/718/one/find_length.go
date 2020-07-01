package one

func FindLength(A []int, B []int) int {
	ans := 0
	// dp[i][j] 代表A[i:]和B[j:]的公共长度
	dp := make([][]int, len(A)+1)
	dp[len(A)] = make([]int, len(B)+1)
	for i := len(A) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(B)+1)
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
				continue
			}
			dp[i][j] = 0
		}
	}
	return ans
}
