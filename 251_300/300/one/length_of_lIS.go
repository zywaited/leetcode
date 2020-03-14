package one

// 动态规划
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	ans := 0
	for i := 1; i <= len(nums); i++ {
		max := 0
		for j := i - 1; j > 0; j-- {
			if nums[i-1] > nums[j-1] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
