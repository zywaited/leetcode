package one

func TotalSteps(nums []int) int {
	st := make([][]int, 0, len(nums))
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		step := 0
		for len(st) > 0 && st[len(st)-1][0] < nums[i] {
			step = max(step+1, st[len(st)-1][1])
			st = st[:len(st)-1]
		}
		ans = max(ans, step)
		st = append(st, []int{nums[i], step})
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
