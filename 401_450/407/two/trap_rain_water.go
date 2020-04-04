package one

// 该实现是对第一种解法的优化
// 堆操作的时候不用每次都不停的变化结构
// 优化是在pop和push方法上，减少重复交换数据

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type node struct {
	row    int
	col    int
	height int
}

type nodes struct {
	ns []*node
	ll int
	el int
}

func TrapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	m, n := len(heightMap), len(heightMap[0])
	ns := NewNodes(m * n)
	// 如果节点已经访问过了，那么重置数据为-1
	// 先把外围的数据加入数据中
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || row == m-1 || col == 0 || col == n-1 {
				ns.push(&node{row: row, col: col, height: heightMap[row][col]})
				heightMap[row][col] = -1
			}
		}
	}
	ans := 0
	for ns.len() > 0 {
		cn := ns.pop()
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
			ns.push(&node{row: nr, col: nc, height: max(heightMap[nr][nc], cn.height)})
			heightMap[nr][nc] = -1

		}
	}
	return ans
}

// 一下是自行实现的堆
func NewNodes(ll int) *nodes {
	return &nodes{
		ns: make([]*node, ll),
		ll: ll,
		el: 0,
	}
}

func (ns *nodes) len() int {
	return ns.el
}

// 小堆顶
func (ns *nodes) heapify(i int) bool {
	l := i<<1 + 1
	r := l + 1
	m := i
	if l < ns.el && ns.ns[l].height < ns.ns[m].height {
		m = l
	}
	if r < ns.el && ns.ns[r].height < ns.ns[m].height {
		m = r
	}
	if m != i {
		ns.ns[m], ns.ns[i] = ns.ns[i], ns.ns[m]
		ns.heapify(m)
		return true
	}
	return false
}

func (ns *nodes) pop() *node {
	if ns.el == 0 {
		return nil
	}
	cn := ns.ns[0]
	ns.el--
	if ns.el == 0 {
		return cn
	}
	// 为了不破坏结构
	// 把最后一个数放到第一位，然后构建
	ns.ns[0] = ns.ns[ns.el]
	ns.heapify(0)
	return cn
}

func (ns *nodes) push(n *node) {
	ns.ns[ns.el] = n
	ns.el++
	// 从底往上开始变化
	for i := (ns.el - 2) >> 1; i >= 0; i = (i - 1) >> 1 {
		// 不需要变化就跳过
		if !ns.heapify(i) {
			break
		}
	}
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
