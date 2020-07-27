package one

func LongestArithSeqLength(A []int) int {
	ans := 0
	// dp[i] 代表第i个数与前所有数的差值&数量
	dp := make([]map[int]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make(map[int]int, len(A))
		if i == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			dp[i][A[i]-A[j]] = dp[j][A[i]-A[j]] + 1
			if dp[i][A[i]-A[j]] > ans {
				ans = dp[i][A[i]-A[j]]
			}
		}
	}
	return ans + 1 // 自身也在里面
}
