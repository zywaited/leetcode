package one

func FindMaxLength(nums []int) int {
	numIndex := map[int]int{}
	numIndex[0] = -1
	sum := 0
	ans := 0
	prevIndex := 0
	ok := false
	for index, num := range nums {
		switch num {
		case 0:
			sum--
		case 1:
			sum++
		}
		if prevIndex, ok = numIndex[sum]; ok {
			ans = max(ans, index-prevIndex)
			continue
		}
		numIndex[sum] = index
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
