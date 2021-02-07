package one

func CheckPossibility(nums []int) bool {
	// 先找到分界点
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}
	if i == len(nums) {
		return true
	}
	// 变更前一个
	j := i + 1
	if i < 2 || nums[i] >= nums[i-2] {
		for ; j < len(nums); j++ {
			if nums[j] < nums[j-1] {
				break
			}
		}
		if j == len(nums) {
			return true
		}
	}
	// 变更当前
	prev := nums[i-1]
	for j = i + 1; j < len(nums); j++ {
		if nums[j] < prev {
			return false
		}
		prev = nums[j]
	}
	return true
}
