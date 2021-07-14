package two

func FindDuplicates(nums []int) []int {
	index := 0
	num := 0
	for index < len(nums) {
		num = nums[index]
		if num > len(nums) {
			index++
			continue
		}
		if num > index && nums[num-1] <= len(nums) {
			nums[index] = nums[num-1]
		} else {
			index++
		}
		if nums[num-1] <= len(nums) {
			nums[num-1] = len(nums) + 1
			continue
		}
		nums[num-1]++
	}
	ans := make([]int, 0, len(nums))
	for index = range nums {
		if nums[index]-len(nums) > 1 {
			ans = append(ans, index+1)
		}
	}
	return ans
}
