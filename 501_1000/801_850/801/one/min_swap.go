package one

func MinSwap(nums1 []int, nums2 []int) int {
	s1 := 0
	s2 := 1
	for i := 1; i < len(nums1); i++ {
		c1 := -1
		c2 := -1
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			c1 = s1
			if 0 <= s2 {
				c2 = s2 + 1
			}
		}
		if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			c1 = minIfGtZero(c1, s2)
			if 0 <= s1 {
				c2 = minIfGtZero(c2, s1+1)
			}
		}
		s1 = c1
		s2 = c2
	}
	return minIfGtZero(s1, s2)
}

func minIfGtZero(f, s int) int {
	if f < 0 {
		return s
	}
	if s < 0 {
		return f
	}
	if f <= s {
		return f
	}
	return s
}
