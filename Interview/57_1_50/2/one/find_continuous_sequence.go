package one

func FindContinuousSequence(target int) [][]int {
	nums := make([]int, 0, target)
	sum := 0
	ans := make([][]int, 0)
	for num := 1; num < target; num++ {
		sum += num
		nums = append(nums, num)
		if len(nums) == 1 {
			continue
		}
		for sum > target {
			sum -= nums[0]
			nums = nums[1:]
		}
		if sum == target {
			ans = append(ans, append([]int{}, nums...))
		}
	}
	return ans
}
