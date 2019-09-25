package one

import "math"

func LargestRectangleArea(heights []int) int {
	// 每次找到两个点之间的最小值，然后比较大小
	maxArea := 0
	nums := len(heights)
	for i := 0; i < nums; i++ {
		minHeight := math.MaxInt32
		for j := i; j < nums; j++ {
			minHeight = min(minHeight, heights[j])
			maxArea = max(maxArea, minHeight*(j-i+1))
		}
	}
	return maxArea
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
