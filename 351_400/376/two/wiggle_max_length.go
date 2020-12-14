package two

func WiggleMaxLength(nums []int) int {
	up, down := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			up, down = 1, 1
			continue
		}
		if nums[i] > nums[i-1] {
			up = down + 1
			continue
		}
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}
