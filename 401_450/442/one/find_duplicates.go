package one

func FindDuplicates(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			ans = append(ans, num)
		}
		nums[num-1] = -nums[num-1]
	}
	return ans
}
