package one

func SumSubarrayMins(arr []int) int {
	stack := make([][2]int, 0)
	ps := 0
	sum := 0
	cs := 0
	for index, num := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1][0]] >= num {
			ps -= stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
		}
		cs = num * (index + 1)
		if len(stack) > 0 {
			cs = num * (index - stack[len(stack)-1][0])
		}
		ps += cs
		stack = append(stack, [2]int{index, cs})
		sum = (sum + ps) % (1e9 + 7)
	}
	return sum
}
