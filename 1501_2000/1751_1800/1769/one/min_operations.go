package one

func MinOperations(boxes string) []int {
	ans := make([]int, len(boxes))
	prevStep := 0
	prevNum := 0
	for index := 0; index < len(boxes); index++ {
		ans[index] = prevStep + prevNum
		prevStep = ans[index]
		prevNum += int(boxes[index] - '0')
	}
	lastStep := 0
	lastNum := 0
	for index := len(boxes) - 1; index >= 0; index-- {
		lastStep += lastNum
		ans[index] += lastStep
		lastNum += int(boxes[index] - '0')
	}
	return ans
}
