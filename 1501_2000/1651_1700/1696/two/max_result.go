package two

func MaxResult(nums []int, k int) int {
	stack := make([][]int, 0, len(nums))
	sum := 0
	for index, num := range nums {
		for len(stack) > 0 && stack[0][0]+k < index {
			stack = stack[1:]
		}
		sum = 0
		if len(stack) > 0 {
			sum = stack[0][1]
		}
		sum += num
		if index+1 == len(nums) {
			return sum
		}
		for len(stack) > 0 && stack[len(stack)-1][1] < sum {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, []int{index, sum})
	}
	return 0
}
