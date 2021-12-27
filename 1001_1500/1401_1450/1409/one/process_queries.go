package one

type bt struct {
	qn         int         // 查询长度
	nums       []int       // 整体数组(值为前缀和)
	numIndexes map[int]int // 每个数字对应的索引
}

func newTree(n, m int) *bt {
	return &bt{
		qn:         n,
		nums:       make([]int, n+m+1),
		numIndexes: make(map[int]int, n),
	}
}

func (t *bt) len() int {
	return len(t.nums)
}

func (t *bt) lowBit(index int) int {
	return index & (-index)
}

func (t *bt) init(num, index int) {
	t.numIndexes[num] = t.qn + index
	t.update(t.qn+index, 1)
}

func (t *bt) update(index, value int) {
	for index >= 0 && index < t.len() {
		t.nums[index] += value
		index += t.lowBit(index)
	}
}

func (t *bt) sum(num int) int {
	index := t.numIndexes[num]
	s := 0
	for index > 0 {
		s += t.nums[index]
		index -= t.lowBit(index)
	}
	return s
}

// 前置移动
func (t *bt) movePrev(num, prevIndex int) {
	index := t.numIndexes[num]
	// 当前减1
	t.update(index, -1)
	// 移动到首位加1
	t.update(t.qn-prevIndex, 1)
	// 更新数字索引
	t.numIndexes[num] = t.qn - prevIndex
}

func ProcessQueries(queries []int, m int) []int {
	// 树状数组
	// 计算每个索引前方数字的数量

	// 长度包含查询数组和m索引
	tree := newTree(len(queries), m)
	// 1-m节点更新
	for index := 1; index <= m; index++ {
		tree.init(index, index)
	}
	ans := make([]int, 0, len(queries))
	for index, query := range queries {
		ans = append(ans, tree.sum(query)-1)
		tree.movePrev(query, index)
	}
	return ans
}
