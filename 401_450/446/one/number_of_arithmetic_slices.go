package one

func NumberOfArithmeticSlices(nums []int) int {
	ans := 0
	dp := make([]map[int]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = map[int]int{}
		for j := i - 1; j > 0; j-- {
			diff := nums[i-1] - nums[j-1]
			num := dp[j][diff]
			ans += num
			dp[i][diff] += num + 1
		}
	}
	return ans
}
