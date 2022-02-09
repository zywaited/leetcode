package one

func MinimumTime(s string) int {
	left := make([]int, len(s)+1)
	left[0] = len(s)
	for i := range s {
		if s[i] == '0' {
			left[i+1] = min(left[i], i)
			continue
		}
		left[i+1] = left[i] + 2
	}
	ans := min(left[len(s)], len(s))
	right := len(s)
	prev := right
	for i := len(s) - 1; i >= 0; i-- {
		right = prev + 2
		if s[i] == '0' {
			right = min(prev, len(s)-i-1)
		}
		ans = min(ans, min(right, len(s)-i)+min(left[i], i))
		prev = right
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
