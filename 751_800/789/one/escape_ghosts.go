package one

func EscapeGhosts(ghosts [][]int, target []int) bool {
	step := abs(target[0]) + abs(target[1])
	for _, ghost := range ghosts {
		if abs(target[0]-ghost[0])+abs(target[1]-ghost[1]) <= step {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
