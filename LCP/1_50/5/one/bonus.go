package one

const mod int = 1e9 + 7

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
		n.l.sum = (n.l.sum + (n.l.ri-n.l.li+1)*n.lazy%mod) % mod
		n.l.lazy = (n.l.lazy + n.lazy) % mod
		n.r.sum = (n.r.sum + (n.r.ri-n.r.li+1)*n.lazy%mod) % mod
		n.r.lazy = (n.r.lazy + n.lazy) % mod
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

// 计算每个节点的区间索引
func (t *tree) index(subs [][]int) {
	t.dfs(subs, 1, 1)
}

func (t *tree) dfs(subs [][]int, i, r int) int {
	t.is[i] = r
	t.os[i] = r
	r++
	for _, j := range subs[i] {
		num := t.dfs(subs, j, r)
		t.os[i] += num
		r += num
	}
	return t.os[i] - t.is[i] + 1
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
		n.sum = (n.sum + (n.ri-n.li+1)*s) % mod
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
	n.sum = (n.l.sum + n.r.sum) % mod
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
		s = (s + t.search(n.l, li, ri)) % mod
	}
	if n.r.li <= ri {
		s = (s + t.search(n.r, li, ri)) % mod
	}
	return s
}

func Bonus(n int, leadership [][]int, operations [][]int) []int {
	// 题目上规定1代表负责人，所以用数组保证有序(也为了方便)
	subs := make([][]int, n+1)
	for _, lp := range leadership {
		subs[lp[0]] = append(subs[lp[0]], lp[1])
	}
	te := newTree(n)
	te.index(subs)
	ans := make([]int, 0)
	for _, operation := range operations {
		switch operation[0] {
		case 1:
			te.update(te.is[operation[1]], te.is[operation[1]], operation[2])
		case 2:
			te.update(te.is[operation[1]], te.os[operation[1]], operation[2])
		case 3:
			ans = append(ans, te.sum(te.is[operation[1]], te.os[operation[1]]))
		}
	}
	return ans
}
