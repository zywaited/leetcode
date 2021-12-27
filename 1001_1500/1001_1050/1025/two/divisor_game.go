package two

import "math"

func DivisorGame(N int) bool {
	// 动态规划
	dp := make([]bool, N+1)
	dp[1] = false
	for n := 2; n <= N; n++ {
		// 选择因数
		for f := int(math.Sqrt(float64(n))); f > 0; f-- {
			if n%f == 0 && f < n && !dp[n-f] {
				dp[n] = true
				break
			}
		}
	}
	return dp[N]
}
