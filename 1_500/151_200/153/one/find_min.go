package one

func FindMin(nums []int) int {
	s, e := 0, len(nums)-1
	for s < e {
		m := s + (e-s)>>1
		if nums[m] > nums[e] {
			s = m + 1
			continue
		}
		e = m
	}
	return nums[s]
}
