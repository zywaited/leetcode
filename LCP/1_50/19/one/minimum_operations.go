package one

func MinimumOperations(leaves string) int {
	dp := make([][3]int, len(leaves)+1)
	for index, leave := range leaves {
		dp[index+1][0] = dp[index][0]
		if leave != 'r' {
			dp[index+1][0]++
		}
		if index == 0 {
			dp[index+1][1] = len(leaves)
			dp[index+1][2] = len(leaves)
			continue
		}
		dp[index+1][1] = min(dp[index][1], dp[index][0])
		if leave != 'y' {
			dp[index+1][1]++
		}
		if index < 2 {
			dp[index+1][2] = len(leaves)
			continue
		}
		dp[index+1][2] = min(dp[index][2], dp[index][1])
		if leave != 'r' {
			dp[index+1][2]++
		}
	}
	return dp[len(leaves)][2]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
