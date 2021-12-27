package two

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
	stack := make([]int, 0, len(heights)+1)
	stack = append(stack, -1)
	prevIndex := -1
	prevHeight := 0
	for index, height := range heights {
		// 一直往前面找
		for prevIndex > -1 && heights[prevIndex] >= height {
			prevHeight = heights[prevIndex]
			stack = stack[:len(stack)-1]
			prevIndex = stack[len(stack)-1]
			// 这里不用当前height高度是为了方便后续计算
			if index-prevIndex-1 >= prevHeight {
				maxArea = max(maxArea, prevHeight*prevHeight)
			}
		}
		stack = append(stack, index)
		prevIndex = index
	}
	// 剩下的就是后面比前面高的，对于中间数据来说就是最低的
	for prevIndex > -1 {
		prevHeight = heights[prevIndex]
		stack = stack[:len(stack)-1]
		prevIndex = stack[len(stack)-1]
		if len(heights)-prevIndex-1 >= prevHeight {
			maxArea = max(maxArea, prevHeight*prevHeight)
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
