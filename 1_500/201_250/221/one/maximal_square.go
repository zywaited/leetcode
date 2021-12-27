package one

import "math"

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	heights := make([]int, len(matrix[0]))
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '0' {
				heights[col] = 0
				continue
			}
			heights[col]++
		}
		maxArea = max(maxArea, maximal(heights))
	}
	return maxArea
}

func maximal(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minHeight := math.MaxInt32
		for j := i; j < len(heights); j++ {
			minHeight = min(minHeight, heights[j])
			if j-i+1 >= minHeight {
				maxArea = max(maxArea, minHeight*minHeight)
			}
			if j-i+1 >= heights[i] && minHeight == heights[i] {
				maxArea = max(maxArea, heights[i]*heights[i])
			}
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
