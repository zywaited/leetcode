package two

func MinimumOperations(leaves string) int {
	var (
		prev [3]int
		next [3]int
	)
	for index, leave := range leaves {
		next[0] = prev[0]
		if leave != 'r' {
			next[0]++
		}
		if index == 0 {
			next[1] = len(leaves)
			next[2] = len(leaves)
			prev, next = next, prev
			continue
		}
		next[1] = min(prev[0], prev[1])
		if leave != 'y' {
			next[1]++
		}
		if index < 2 {
			next[2] = len(leaves)
			prev, next = next, prev
			continue
		}
		next[2] = min(prev[1], prev[2])
		if leave != 'r' {
			next[2]++
		}
		prev, next = next, prev
	}
	return prev[2]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
