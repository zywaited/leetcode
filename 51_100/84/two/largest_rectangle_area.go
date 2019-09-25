package two

// 采用栈
// 如果后一个比前一个高，则组合面积比前一个大，但不一定比后一个大
func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]int, 0, len(heights)+1)
	stack = append(stack, -1) // 方便计算
	prevIndex := -1
	prevHeight := 0
	for index, height := range heights {
		// 直到比前面的高
		for prevIndex != -1 && heights[prevIndex] >= height { // 这里不用等号也可以
			prevHeight = heights[prevIndex]
			stack = stack[:len(stack)-1]
			prevIndex = stack[len(stack)-1]
			maxArea = max(maxArea, prevHeight*(index-prevIndex-1)) // 用前前一个的索引，所以要减1
		}
		stack = append(stack, index)
		prevIndex = index
	}
	// 此时stack一定会有最后一个数据，所以还要继续计算面积
	for prevIndex != -1 {
		prevHeight = heights[prevIndex]
		stack = stack[:len(stack)-1]
		prevIndex = stack[len(stack)-1]
		maxArea = max(maxArea, prevHeight*(len(heights)-prevIndex-1))
	}
	return maxArea
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
