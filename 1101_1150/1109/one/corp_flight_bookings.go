package one

func CorpFlightBookings(bookings [][]int, n int) []int {
	t := newTree(n)
	for _, booking := range bookings {
		t.update(booking[0], booking[1], booking[2])
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = t.sum(i+1, i+1)
	}
	return ans
}

type node struct {
	li   int // 左索引区间
	ri   int // 右索引区间
	sum  int // LeetCoin数量
	lazy int // 待发放的数量
	l    *node
	r    *node
}

func (n *node) lr() {
	// 叶子节点
	if n.li == n.ri {
		return
	}
	// 延迟初始化左右节点
	if n.l == nil {
		n.l = &node{li: n.li, ri: (n.li + n.ri) >> 1}
	}
	if n.r == nil {
		n.r = &node{li: n.l.ri + 1, ri: n.ri}
	}
}

// 更新节点的懒加载数据
func (n *node) down() {
	n.lr()
	if n.lazy != 0 {
		n.l.sum += (n.l.ri - n.l.li + 1) * n.lazy
		n.l.lazy += n.lazy
		n.r.sum += (n.r.ri - n.r.li + 1) * n.lazy
		n.r.lazy += n.lazy
		n.lazy = 0
	}
}

type tree struct {
	is   []int // 每个节点对应的左索引
	os   []int // 每个节点对应的右索引
	root *node // 节点
}

func newTree(n int) *tree {
	te := &tree{
		is:   make([]int, n+1),
		os:   make([]int, n+1),
		root: &node{li: 1, ri: n},
	}
	return te
}

// 更新当前区间数据
func (t *tree) update(li, ri, s int) {
	if li > t.root.ri || ri < t.root.li {
		return
	}
	t.adjust(t.root, li, ri, s)
}

func (t *tree) adjust(n *node, li, ri, s int) {
	if n.li >= li && n.ri <= ri {
		// 包含
		n.sum += (n.ri - n.li + 1) * s
		n.lazy += s
		return
	}
	n.down()
	if n.l.ri >= li {
		t.adjust(n.l, li, ri, s)
	}
	if n.r.li <= ri {
		t.adjust(n.r, li, ri, s)
	}
	n.sum = n.l.sum + n.r.sum
}

// 计算和
func (t *tree) sum(li, ri int) int {
	if li > t.root.ri || ri < t.root.li {
		return 0
	}
	return t.search(t.root, li, ri)
}

func (t *tree) search(n *node, li, ri int) int {
	if n.li >= li && n.ri <= ri {
		return n.sum
	}
	n.down()
	s := 0
	if n.l.ri >= li {
		s += t.search(n.l, li, ri)
	}
	if n.r.li <= ri {
		s += t.search(n.r, li, ri)
	}
	return s
}
