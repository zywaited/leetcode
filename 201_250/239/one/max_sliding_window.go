package one

func MaxSlidingWindow(nums []int, k int) []int {
	var r []int
	var dq []int
	nl := len(nums)
	dl := 0
	ds := 0
	for i := 0; i < nl; i++ {
		// 先保持递减
		ds = 0
		dl = len(dq)
		for dl > 0 && nums[dq[dl-1]] < nums[i] {
			dl--
			ds++
		}
		if ds > 0 {
			dq = dq[:dl]
		}
		dq = append(dq, i)
		dl++
		// 首位是否有效, 窗口最大值为k
		if dq[0] <= i-k {
			dq = dq[1:]
		}
		if i < k-1 {
			continue
		}
		// 首位即是当前窗口最大值
		r = append(r, nums[dq[0]])
	}
	return r
}
