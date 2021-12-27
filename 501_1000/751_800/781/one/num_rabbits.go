package one

func NumRabbits(answers []int) int {
	sum := 0
	nums := make([]int, 1000)
	for _, answer := range answers {
		sum++
		if nums[answer] == answer {
			nums[answer] = 0
			continue
		}
		nums[answer]++
	}
	for i := range nums {
		if nums[i] > 0 {
			sum += i - nums[i] + 1
		}
	}
	return sum
}
