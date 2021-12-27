package one

// 把每一层都想像成一个柱状图
// 这样解决思路就跟84题类似
func MaximalRectangle(matrix [][]byte) int {
	maxArea := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxArea
	}
	dp := make([]int, len(matrix[0])) // 每一列1的个数
	for _, cols := range matrix {
		for col, v := range cols {
			if v == '1' {
				dp[col]++
				continue
			}
			dp[col] = 0 // 注意这里一定要清零，因为柱状图不连续了
		}
		maxArea = max(maxArea, largestRectangleArea(dp))
	}
	return maxArea
}

// 采用栈
// 如果后一个比前一个高，则组合面积比前一个大，但不一定比后一个大
func largestRectangleArea(heights []int) int {
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
