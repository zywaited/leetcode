package two

import "math"

// 动态规划
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	ans := 0
	for i, num := range nums {
		dp[i+1] = math.MaxInt32
		for j := i; j >= 0; j-- {
			if (j == 0 || num > dp[j]) && num < dp[j+1] {
				dp[j+1] = num
				if j+1 > ans {
					ans = j + 1
				}
			}
		}
	}
	return ans
}
