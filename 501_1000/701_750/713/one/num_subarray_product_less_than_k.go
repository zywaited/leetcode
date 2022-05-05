package one

func NumSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	ms := 1
	pi := 0
	for i := 0; i < len(nums); i++ {
		ms *= nums[i]
		for ; k <= ms && pi <= i; pi++ {
			ms /= nums[pi]
		}
		ans += i - pi + 1
	}
	return ans
}
