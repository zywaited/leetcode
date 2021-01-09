package one

func Rotate(nums []int, k int) {
	if len(nums) > 0 {
		i := 0
		tmp := nums[0]
		for (i+k)%len(nums) > 0 {
			i = (i + k) % len(nums)
			nums[i], tmp = tmp, nums[i]
		}
		nums[0] = tmp
	}
}
