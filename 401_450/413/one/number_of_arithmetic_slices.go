package one

func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	ans := 0
	left := 0
	diff := nums[1] - nums[0]
	for right := 2; right <= len(nums); right++ {
		if right < len(nums) && nums[right]-nums[right-1] == diff {
			continue
		}
		if right-left >= 3 {
			ans += ((right - left - 2) * (right - left - 1)) >> 1
		}
		if right < len(nums) {
			left = right - 1
			diff = nums[right] - nums[left]
		}
	}
	return ans
}
