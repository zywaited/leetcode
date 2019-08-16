package two

// 循环
func SearchRange(nums []int, target int) []int {
	rs := []int{-1, -1}
	if len(nums) < 1 {
		return rs
	}
	// 先找左边
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= target {
			r = m
			continue
		}
		l = m + 1 // 加1防止l=0,r=1时,l=0，如果m+1,则r=m-1
	}
	// 是否有目标值
	if nums[l] != target {
		return rs
	}
	// 最后找右边，这里让r等于数组的长度，因为l可能就是最后一个数字
	rs[0] = l
	r = len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] <= target {
			l = m + 1
			continue
		}
		r = m
	}
	// 因为r初始为数组长度
	rs[1] = r - 1
	return rs
}
