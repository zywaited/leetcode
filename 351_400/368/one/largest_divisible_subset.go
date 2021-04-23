package one

import "sort"

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	maxIndex := -1
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = append(dp[i], nums[i])
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = dp[i][:1]
					dp[i] = append(dp[i], dp[j]...)
				}
			}
		}
		if maxIndex == -1 || len(dp[i]) > len(dp[maxIndex]) {
			maxIndex = i
		}
	}
	return dp[maxIndex]
}
