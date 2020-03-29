package one

// 大部分题写的解法都是bfs，这里就用一下dfs吧
var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func MovingCount(m int, n int, k int) int {
	return dfs(m, n, 0, 0, k, make(map[int]byte))
}

// 由题目可以x,y的范围是0-99
func dfs(m, n, i, j, k int, visited map[int]byte) int {
	if visited[i*100+j] == 1 {
		return 0
	}
	visited[i*100+j] = 1
	num := 1
	for _, direct := range directs {
		xi, xj := i+direct[0], j+direct[1]
		// 下面的条件就不合并了，方便看
		if xi < 0 || xi >= m || xj < 0 || xj >= n {
			continue
		}
		if visited[xi*100+xj] == 1 {
			continue
		}
		if xi/10+xi%10+xj/10+xj%10 > k {
			continue
		}
		num += dfs(m, n, xi, xj, k, visited)
	}
	return num
}
