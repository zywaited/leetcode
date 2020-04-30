package one

// 大顶堆+小顶堆
// 数据量各一半
type MedianFinder struct {
	maxHalf *nodes // 左边一半的数据
	minHalf *nodes // 右边一半的数据
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHalf: newNodes(),
		minHalf: newNodes(),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	// 这里也可以用<=，不过FindMedian需要多处理右边比左边多1的情况
	if mf.maxHalf.len() == mf.minHalf.len() {
		// 都先加到左边
		if mf.maxHalf.len() == 0 || mf.minHalf.top() >= num {
			// 左边没数据或者不比右边大
			mf.maxHalf.pushMax(num)
			return
		}
		// 把右边的最小值放到左边
		mf.maxHalf.pushMax(mf.minHalf.top())
		// 加入右边
		mf.minHalf.replaceMin(num)
		return
	}
	// 左边比右边多1的时候
	if mf.maxHalf.top() <= num {
		mf.minHalf.pushMin(num)
		return
	}
	mf.minHalf.pushMin(mf.maxHalf.top())
	mf.maxHalf.replaceMax(num)
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHalf.len() == 0 {
		return 0
	}
	if mf.maxHalf.len() > mf.minHalf.len() {
		return float64(mf.maxHalf.top())
	}
	return float64(mf.maxHalf.top()) + float64(mf.minHalf.top()-mf.maxHalf.top())/float64(2)
}

// 自行实现的堆
type nodes []int

func newNodes() *nodes {
	return &nodes{}
}

func (ns *nodes) len() int {
	return len(*ns)
}

func (ns *nodes) swap(i, j int) {
	(*ns)[i], (*ns)[j] = (*ns)[j], (*ns)[i]
}

func (ns *nodes) min(i, j int) bool {
	return (*ns)[i] < (*ns)[j]
}

func (ns *nodes) max(i, j int) bool {
	return (*ns)[i] > (*ns)[j]
}

func (ns *nodes) heapify(i, ll int, compare func(int, int) bool) bool {
	l := i<<1 + 1
	r := l + 1
	m := i
	if l < ll && compare(l, m) {
		m = l
	}
	if r < ll && compare(r, m) {
		m = r
	}
	if m != i {
		ns.swap(i, m)
		ns.heapify(m, ll, compare)
		return true
	}
	return false
}

func (ns *nodes) push(num int, compare func(int, int) bool) {
	// 从底往上开始变化
	*ns = append(*ns, num)
	for i := (len(*ns) - 2) >> 1; i >= 0; i = (i - 1) >> 1 {
		// 不需要变化就跳过
		if !ns.heapify(i, ns.len(), compare) {
			break
		}
	}
}

func (ns *nodes) pushMax(num int) {
	ns.push(num, ns.max)
}

func (ns *nodes) pushMin(num int) {
	ns.push(num, ns.min)
}

func (ns *nodes) replace(num int) int {
	pn := ns.top()
	(*ns)[0] = num
	return pn
}

func (ns *nodes) replaceMax(num int) int {
	pn := ns.replace(num)
	ns.heapify(0, ns.len(), ns.max)
	return pn
}

func (ns *nodes) replaceMin(num int) int {
	pn := ns.replace(num)
	ns.heapify(0, ns.len(), ns.min)
	return pn
}

func (ns *nodes) top() int {
	return (*ns)[0]
}
