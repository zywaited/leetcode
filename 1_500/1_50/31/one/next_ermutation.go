package one

func NextPermutation(nums []int) {
	// 从后往前找到较小的数
	si := -1
	for si = len(nums) - 2; si >= 0; si-- {
		if nums[si] >= nums[si+1] {
			continue
		}
		// 为了性能可以用二分查找比当前大并且是较小的数
		r := -1
		e := len(nums) - 1
		s := si + 1
		for s <= e {
			m := s + (e-s)>>1
			if nums[m] <= nums[si] {
				// 后续不存在比i大的数字了
				if nums[s] <= nums[si] {
					break
				}
				e = m - 1
				continue
			}
			r = m
			s = m + 1
		}
		nums[si], nums[r] = nums[r], nums[si]
		break
	}
	// 对后面的数据排序
	//sort.Ints(nums[si+1:])
	// 后面的数据本来就是有序的，只是需要反转数组，所以不用整体排序(性能)
	s := si + 1
	e := len(nums) - 1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}
