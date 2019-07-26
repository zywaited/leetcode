package ignore

func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	return dp(0, nums, make(map[int]bool))
}

func dp(start int, nums []int, mp map[int]bool) bool {
	// 已经到结尾了
	if start >= len(nums)-1 {
		return true
	}
	// 已经选择过了
	if _, ok := mp[start]; ok {
		return mp[start]
	}
	maxStep := nums[start]
	if maxStep > len(nums)-start {
		maxStep = len(nums) - start
	}
	for step := maxStep; step > 0; step-- {
		if !dp(start+step, nums, mp) {
			continue
		}
		mp[start] = true
		return true
	}
	mp[start] = false
	return false
}
