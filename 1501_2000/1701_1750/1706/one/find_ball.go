package one

func FindBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
next:
	for i := 0; i < n; i++ {
		c := i
		for r := 0; r < m; r++ {
			pc := grid[r][c]
			c += grid[r][c]
			if c == -1 || c == n || grid[r][c] != pc {
				ans[i] = -1
				continue next
			}
			continue
		}
		ans[i] = c
	}
	return ans
}
