package two

func SmallerNumbersThanCurrent(nums []int) []int {
	nn := make([]int, 101)
	for _, num := range nums {
		nn[num]++
	}
	dp := make([]int, 101)
	for index := 1; index < len(dp); index++ {
		dp[index] = dp[index-1] + nn[index-1]
	}
	ans := make([]int, len(nums))
	for index, num := range nums {
		ans[index] = dp[num]
	}
	return ans
}
