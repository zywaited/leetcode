package three

func Trap(height []int) int {
	sm := 0
	stack := make([]int, 0)
	sum := 0
	for hi := range height {
		if height[sm] > height[hi] {
			stack = append(stack, sm)
			sm = hi
			continue
		}
		for len(stack) > 0 && height[sm] < height[hi] {
			sum += (min(height[stack[len(stack)-1]], height[hi]) - height[sm]) * (hi - stack[len(stack)-1] - 1)
			if height[stack[len(stack)-1]] >= height[hi] {
				sm = hi
			}
			if height[stack[len(stack)-1]] < height[hi] {
				sm = stack[len(stack)-1]
			}
			if height[stack[len(stack)-1]] <= height[hi] {
				stack = stack[:len(stack)-1]
			}
		}
		if len(stack) == 0 {
			sm = hi
		}
		if height[sm] == height[hi] {
			sm = hi
		}
	}
	return sum
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
