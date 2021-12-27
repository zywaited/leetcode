package one

// 广度优先搜索
var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

const max = 11

func OrangesRotting(grid [][]int) int {
	// 对列
	type node struct {
		m   int
		row int
		col int
	}
	queue := make([]node, 0)
	// 实际腐烂的橘子数
	cnt := 0
	// 总共的橘子数
	ant := 0
	// 访问过了
	visited := make(map[int]byte, 0)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				continue
			}
			ant++
			if grid[row][col] == 2 {
				cnt++
				queue = append(queue, node{m: 0, row: row, col: col})
				visited[row*max+col] = 1
			}
		}
	}
	// 不需要额外操作
	if cnt == ant {
		return 0
	}
	for len(queue) > 0 {
		cn := queue[0]
		queue = queue[1:]
		for _, direct := range directs {
			row, col := cn.row+direct[0], cn.col+direct[1]
			// 超出范围内
			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
				continue
			}
			// 不需要操作
			if grid[row][col] != 1 || visited[row*max+col] == 1 {
				continue
			}
			cnt++
			// 所有橘子都腐烂了
			if cnt == ant {
				return cn.m + 1
			}
			queue = append(queue, node{m: cn.m + 1, row: row, col: col})
			visited[row*max+col] = 1
		}
	}
	return -1
}
