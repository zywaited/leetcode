package one

func NextGreaterElements(nums []int) []int {
	// 拼接数组
	// [1,2,1] => [1,2,1,1,2]
	stack := make([]int, 0, len(nums)*2)
	rs := make([]int, len(nums))
	realIndex := 0
	for index := len(nums)*2 - 2; index >= 0; index-- {
		realIndex = index
		if index >= len(nums) {
			realIndex = index - len(nums)
		}
		for len(stack) > 0 && stack[len(stack)-1] <= nums[realIndex] {
			stack = stack[:len(stack)-1]
		}
		if index < len(nums) {
			// 真实数据
			if len(stack) > 0 {
				rs[realIndex] = stack[len(stack)-1]
			} else {
				rs[realIndex] = -1
			}
		}
		stack = append(stack, nums[realIndex])
	}
	return rs
}
