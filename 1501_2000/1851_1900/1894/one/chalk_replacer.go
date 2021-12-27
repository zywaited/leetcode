package one

func ChalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := range chalk {
		sum += chalk[i]
		if sum > k {
			return i
		}
	}
	remainder := k % sum
	sum = 0
	for i := range chalk {
		sum += chalk[i]
		if sum > remainder {
			return i
		}
	}
	return -1
}
