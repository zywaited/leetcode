package one

func EqualSubstring(s string, t string, maxCost int) int {
	left := 0
	cost := 0
	ans := 0
	for index := range s {
		cost += abs(int(s[index]) - int(t[index]))
		if cost > maxCost {
			for left <= index && cost > maxCost {
				cost -= abs(int(s[left]) - int(t[left]))
				left++
				if cost <= maxCost {
					break
				}
			}
			if left > index {
				continue
			}
		}
		ans = max(ans, index-left+1)
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
