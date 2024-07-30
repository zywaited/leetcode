package one

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = math.MinInt
		for j, x := range nums[:i] {
			if -target <= nums[i]-x && nums[i]-x <= target {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	if f[n-1] < 0 {
		return -1
	}
	return f[n-1]
}
