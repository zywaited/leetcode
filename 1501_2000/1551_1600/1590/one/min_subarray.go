package one

func minSubarray(nums []int, p int) int {
	remainder := 0
	for _, num := range nums {
		remainder = (remainder + num) % p
	}
	if remainder == 0 {
		return 0
	}
	remainderIndexes := make(map[int]int, len(nums))
	remainderIndexes[0] = 0
	currentRemainder := 0
	length := len(nums)
	for index, num := range nums {
		currentRemainder = (currentRemainder + num) % p
		prevRemainder := (currentRemainder - remainder + p) % p
		if prevIndex, ok := remainderIndexes[prevRemainder]; ok {
			length = min(length, index+1-prevIndex)
		}
		remainderIndexes[currentRemainder] = index + 1
	}
	if length == len(nums) {
		return -1
	}
	return length
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
