package one

func GridGame(grid [][]int) int64 {
	n := len(grid[0])
	first := make([]int64, n+1)
	second := make([]int64, n+1)
	for index := 0; index < n; index++ {
		first[index+1] = first[index] + int64(grid[0][index])
		second[index+1] = second[index] + int64(grid[1][index])
	}
	ans := min(first[n]-first[1], second[n-1])
	for index := 2; index <= n; index++ {
		ans = min(ans, max(first[n]-first[index], second[index-1]))
	}
	return ans
}

func min(f, s int64) int64 {
	if f <= s {
		return f
	}
	return s
}

func max(f, s int64) int64 {
	if f >= s {
		return f
	}
	return s
}
