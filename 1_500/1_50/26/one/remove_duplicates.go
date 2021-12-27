package one

func RemoveDuplicates(nums []int) int {
	nl := len(nums)
	if nl < 2 {
		return nl
	}
	rs := 1
	for i := 1; i < nl; i++ {
		if nums[i]^nums[i-1] != 0 {
			nums[rs] = nums[i]
			rs++
		}
	}
	return rs
}
