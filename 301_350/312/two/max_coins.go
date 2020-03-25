package two

func MaxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	coins := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if dp[i] == nil {
			dp[i] = make([]int, len(nums))
		}
		for j := i; j < len(nums); j++ {
			for k := i; k <= j; k++ {
				coins = nums[k]
				if i-1 >= 0 {
					coins *= nums[i-1]
				}
				if j+1 < len(nums) {
					coins *= nums[j+1]
				}
				if k-1 >= 0 {
					coins += dp[i][k-1]
				}
				if k+1 < len(nums) {
					coins += dp[k+1][j]
				}
				dp[i][j] = max(dp[i][j], coins)
			}
		}
	}
	return dp[0][len(nums)-1]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
