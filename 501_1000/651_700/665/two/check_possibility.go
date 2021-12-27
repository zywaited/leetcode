package two

// 优化第一种解法
func CheckPossibility(nums []int) bool {
	pi := 0
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[pi] {
			pi = i
			continue
		}
		if cnt == 0 {
			return false
		}
		cnt = 0
		// 变更当前为前一个值
		pi = i - 1
		// 如果可以，变更前一个为当前值
		if (i < 2 || nums[i] >= nums[i-2]) && nums[i] < nums[pi] {
			pi = i
		}
	}
	return true
}
