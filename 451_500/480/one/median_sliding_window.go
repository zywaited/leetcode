package one

import "container/heap"

// 大小顶堆
type hp struct {
	nodes []int
	l     int
	rm    map[int]int
	less  func(nodes []int, i, j int) bool
}

func newHp(less func(nodes []int, i, j int) bool) *hp {
	return &hp{rm: make(map[int]int), less: less}
}

func (h *hp) Len() int {
	return len(h.nodes)
}

func (h *hp) Less(i, j int) bool {
	return h.less(h.nodes, i, j)
}

func (h *hp) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *hp) Pop() interface{} {
	n := h.nodes[h.Len()-1]
	h.nodes = h.nodes[:h.Len()-1]
	if h.rm[n] > 0 {
		h.rm[n]--
		return n
	}
	h.l--
	return n
}

func (h *hp) Push(n interface{}) {
	h.nodes = append(h.nodes, n.(int))
	h.l++
}

func (h *hp) adjust() {
	for h.Len() > 0 {
		num := h.nodes[0]
		if h.rm[num] == 0 {
			break
		}
		heap.Remove(h, 0)
	}
}

func (h *hp) remove(num int) {
	h.rm[num]++
	h.l--
	h.adjust()
}

func (h *hp) len() int {
	return h.l
}

func minLess(nodes []int, i, j int) bool {
	return nodes[i] < nodes[j]
}

func maxLess(nodes []int, i, j int) bool {
	return nodes[i] > nodes[j]
}

// 不考虑k为0的情况
func MedianSlidingWindow(nums []int, k int) []float64 {
	// 小顶堆存放大数据【保持sh.len >= bh.len】
	sh := newHp(minLess)
	// 大顶堆存放小数据
	bh := newHp(maxLess)

	// 这里独立个方法仅仅为了简化if-else
	add := func(num int) {
		if sh.len() == bh.len() {
			if sh.len() == 0 || sh.nodes[0] <= num || bh.nodes[0] <= num {
				heap.Push(sh, num)
				return
			}
			heap.Push(sh, heap.Pop(bh))
			heap.Push(bh, num)
			return
		}
		if sh.nodes[0] >= num {
			heap.Push(bh, num)
			return
		}
		heap.Push(bh, heap.Pop(sh))
		heap.Push(sh, num)
	}
	mn := func() float64 {
		if sh.len() == 0 {
			return 0
		}
		if sh.len() > bh.len() {
			return float64(sh.nodes[0])
		}
		return float64(bh.nodes[0]) + (float64(sh.nodes[0])-float64(bh.nodes[0]))/float64(2)
	}
	rn := func(num int) {
		if sh.len() == 0 {
			return
		}
		if sh.len() == bh.len() {
			if sh.nodes[0] > num {
				bh.remove(num)
				return
			}
			sh.remove(num)
			heap.Push(sh, heap.Pop(bh))
			bh.adjust()
			return
		}
		if sh.nodes[0] <= num {
			sh.remove(num)
			return
		}
		bh.remove(num)
		heap.Push(bh, heap.Pop(sh))
		sh.adjust()
	}

	ans := make([]float64, 0)
	for i := 0; i < len(nums); i++ {
		add(nums[i])
		if sh.len()+bh.len() < k {
			continue
		}
		sh.adjust()
		bh.adjust()
		ans = append(ans, mn())
		rn(nums[i+1-k])
	}
	return ans
}
