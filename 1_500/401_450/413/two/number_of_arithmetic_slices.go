package two

func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	ans := 0
	left := 0
	diff := nums[1] - nums[0]
	dp := 0
	for right := 2; right < len(nums); right++ {
		// 第一个方法利用等差求和
		// 这里就直接正向处理，每次叠加
		if nums[right]-nums[right-1] == diff {
			dp++
			ans += dp
		} else {
			dp = 0
		}
		left = right - 1
		diff = nums[right] - nums[left]
	}
	return ans
}
