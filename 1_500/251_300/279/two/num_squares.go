package two

import "math"

func NumSquares(n int) int {
	dp := make([]int, n+1)
	ap := int(math.Sqrt(float64(n)))
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j <= ap; j++ {
			if i < j*j {
				break
			}
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
