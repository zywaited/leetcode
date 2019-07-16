package one

// 思路:
// 尽量匹配最高最长的区间
func MaxArea(height []int) int {
	ma := 0
	i, j := 0, len(height)-1
	for i != j {
		w := j - i
		h := height[i]
		if height[j] < h {
			h = height[j]
		}

		ca := w * h
		if ca > ma {
			ma = ca
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return ma
}
