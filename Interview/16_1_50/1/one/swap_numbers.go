package one

func SwapNumbers(numbers []int) []int {
	numbers[0] += numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] -= numbers[1]
	return numbers
}
