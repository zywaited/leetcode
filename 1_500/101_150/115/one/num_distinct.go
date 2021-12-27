package one

func NumDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	var (
		prev    = make([]int, len(t)+1)
		current = make([]int, len(t)+1)
	)
	for i := 0; i < len(s); i++ {
		prev[0] = 1
		for j := 1; j <= len(t); j++ {
			current[j] = prev[j]
			if prev[j-1] > 0 && s[i] == t[j-1] {
				current[j] += prev[j-1]
			}
		}
		prev, current = current, prev
	}
	return prev[len(t)]
}
