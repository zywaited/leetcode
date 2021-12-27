package two

import "sort"

// 空间压缩
func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	maxIndex := -1
	maxNum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if maxIndex == -1 || dp[i] > maxNum {
			maxIndex = i
			maxNum = dp[i]
		}
	}
	ans := make([]int, 0, maxNum)
	i := maxIndex
	j := maxNum
	pn := nums[maxIndex]
	for ; i < len(nums) && j > 0; i++ {
		if nums[i]%pn == 0 && dp[i] == j {
			ans = append(ans, nums[i])
			pn = nums[i]
			j--
		}
	}
	return ans
}
