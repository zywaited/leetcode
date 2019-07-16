package one

func RemoveDuplicates(nums []int) int {
	res := 0
	cn := 1
	nl := len(nums)
	if nl < 3 {
		return nl
	}
	l := 0
	r := l + 1
	for ; r < nl; r++ {
		// 找到当前数字的个数
		// 因为是有序的，所以相同数字都是挨着的
		if nums[r] == nums[l] {
			if nums[l+cn] != nums[r] {
				nums[l+cn] = nums[r]
			}
			cn++
			continue
		}
		// 如果个数小于等于2，满足要求
		if cn <= 2 {
			res += cn
			l += cn
			// 特别注意如果有一次相同数字个数大于2，那么只保留了2个，所以得索引得用左指针
			if nums[l] != nums[r] {
				nums[l] = nums[r]
			}
		} else {
			// 大于2个只保留2ge
			l += 2
			nums[l] = nums[r]
			res += 2
		}
		cn = 1
	}
	if cn != 1 {
		// 说明最后一次数字多于2个
		res += 2
	} else {
		res++
	}
	return res
}
