package one

func Trap(height []int) int {
	lm := make([]int, len(height)+1)
	for hi := range height {
		lm[hi+1] = max(lm[hi], height[hi])
	}
	rm := make([]int, len(height)+1)
	for hi := len(height) - 1; hi >= 0; hi-- {
		rm[hi] = max(rm[hi+1], height[hi])
	}
	sum := 0
	for hi := 1; hi < len(height); hi++ {
		sum += min(lm[hi+1], rm[hi]) - height[hi]
	}
	return sum
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
