package one

func WinnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	for num := 1; num <= n; num++ {
		for tn := 1; tn*tn <= num; tn++ {
			if !dp[num-tn*tn] {
				dp[num] = true
				break
			}
		}
	}
	return dp[n]
}
