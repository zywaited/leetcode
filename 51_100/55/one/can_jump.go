package one

func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	// 从右往左，判断前一个是否能够到达最后
	// 最后判断第一个
	step := 1
	can := false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < step {
			step++
			can = false
			continue
		}
		step = 1
		can = true
	}

	return can
}
