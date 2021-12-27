package two

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] > nums[r] {
			l = m + 1
			continue
		}
		if nums[m] < nums[r] {
			r = m
			continue
		}
		r--
	}
	return nums[l]
}
