package one

func IsPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	pn := 0
	pn1, pn2, pn3 := 0, 0, 0
	i := 0
	pi := 0
	n := 0
	n2, n3 := 0, 0
	for i < len(nums) {
		pi = i
		for i++; i < len(nums) && nums[i] == nums[pi]; i++ {
		}
		n = i - pi
		if pi == 0 || nums[pi]-1 != pn {
			if pn1 > 0 || pn2 > 0 {
				return false
			}
			pn1, pn2, pn3 = n, 0, 0
			pn = nums[pi]
			continue
		}
		if n-pn1 < pn2 {
			return false
		}
		n2, n3 = 0, 0
		if pn1 > 0 {
			n2 = pn1
			n -= pn1
		}
		if pn2 > 0 {
			n3 = pn2
			n -= pn2
		}
		if pn3 >= n {
			n3 += n
			n = 0
		} else {
			n3 += pn3
			n -= pn3
		}
		pn1, pn2, pn3 = n, n2, n3
		pn = nums[pi]
	}
	return pn1 == 0 && pn2 == 0
}
