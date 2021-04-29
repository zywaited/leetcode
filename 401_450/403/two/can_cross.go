package two

func CanCross(stones []int) bool {
	dp := make([][]bool, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]bool, len(stones))
		if i == 0 {
			dp[i][0] = true
		}
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				// 比长度还大，前置步数不可能到这
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if dp[i][k] && i == len(stones)-1 {
				return true
			}
		}
	}
	return false
}
