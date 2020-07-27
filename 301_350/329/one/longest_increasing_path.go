package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func LongestIncreasingPath(matrix [][]int) int {
	ans := 0
	nums := make([][][3]int, len(matrix))
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if nums[i] == nil {
			nums[i] = make([][3]int, len(matrix[i]))
		}
		if nums[i][j][0] > 0 {
			return
		}
		nums[i][j][0] = 1
		nums[i][j][1]++
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[ni]) {
				continue
			}
			if matrix[i][j] <= matrix[ni][nj] {
				continue
			}
			dfs(ni, nj)
			nums[i][j][1] = max(nums[i][j][1], nums[ni][nj][1]+1)
		}
		ans = max(ans, nums[i][j][1])
	}
	for i := range matrix {
		if nums[i] == nil {
			nums[i] = make([][3]int, len(matrix[i]))
		}
		for j := range matrix[i] {
			dfs(i, j)
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
