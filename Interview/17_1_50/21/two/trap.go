package two

func Trap(height []int) int {
	lm := 0
	rm := 0
	l, r := 0, len(height)-1
	sum := 0
	for l <= r {
		if lm <= rm {
			if height[l] >= lm {
				lm = height[l]
			} else {
				sum += lm - height[l]
			}
			l++
			continue
		}
		if height[r] >= rm {
			rm = height[r]
		} else {
			sum += rm - height[r]
		}
		r--
		continue
	}
	return sum
}
