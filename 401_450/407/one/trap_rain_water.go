package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type node struct {
	row    int
	col    int
	height int
}

func TrapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	m, n := len(heightMap), len(heightMap[0])
	nodes := make([]node, 0, m*n)
	// 如果节点已经访问过了，那么重置数据为-1
	// 先把外围的数据加入数据中
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || row == m-1 || col == 0 || col == n-1 {
				nodes = append(nodes, node{row: row, col: col, height: heightMap[row][col]})
				heightMap[row][col] = -1
			}
		}
	}
	heap(nodes)
	ans := 0
	for len(nodes) > 0 {
		cn := nodes[0]
		nodes = nodes[1:]
		// 4个方向
		for _, direct := range directs {
			nr, nc := cn.row+direct[0], cn.col+direct[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || heightMap[nr][nc] == -1 {
				continue
			}
			// 只要比外围矮就可以加水
			if heightMap[nr][nc] < cn.height {
				ans += cn.height - heightMap[nr][nc]
			}
			// 然后此时变更外围数据
			nodes = append(nodes, node{row: nr, col: nc, height: max(heightMap[nr][nc], cn.height)})
			heightMap[nr][nc] = -1
		}
		// 数据变更后要重新处理小顶堆
		heap(nodes)
	}
	return ans
}

func heap(nodes []node) {
	// 处理堆
	for i := len(nodes)>>1 - 1; i >= 0; i-- {
		heapify(nodes, i)
	}
}

// 小堆顶
func heapify(nodes []node, i int) {
	l := i<<1 + 1
	r := l + 1
	m := i
	if l < len(nodes) && nodes[l].height < nodes[m].height {
		m = l
	}
	if r < len(nodes) && nodes[r].height < nodes[m].height {
		m = r
	}
	if m != i {
		nodes[m], nodes[i] = nodes[i], nodes[m]
		heapify(nodes, m)
	}
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
