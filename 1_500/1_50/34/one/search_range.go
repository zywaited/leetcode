package one

// 递归
// 也可以重新写个函数
// search(nums []int, index int, bool left)
func SearchRange(nums []int, target int) []int {
	rs := []int{-1, -1}
	if len(nums) <= 1 {
		if len(nums) == 1 && nums[0] == target {
			rs[0], rs[1] = 0, 0
		}
		return rs
	}
	left, right := 0, len(nums)-1
	// 需要加1，不然会死循环
	mid := left + (right-left)>>1 + 1
	if nums[mid] >= target {
		// search(nums, mid-1, true)
		rs = SearchRange(nums[:mid], target)
	}
	if nums[mid] <= target {
		// search(nums, mid, false)
		ri := SearchRange(nums[mid:], target)
		if ri[0] != -1 {
			// 被截取，实际长度需要加上mid
			// 如果按照上诉说的重新写个函数就没这些了
			ri[0] += mid
			ri[1] += mid
			// 合并区间
			if rs[0] == -1 || rs[0] > ri[0] {
				rs[0] = ri[0]
			}
			if rs[1] == -1 || rs[1] < ri[1] {
				rs[1] = ri[1]
			}
		}
	}
	return rs
}
