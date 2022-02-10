package one

import (
	"container/heap"
	"math"
)

func MinimumDifference(nums []int) int64 {
	n := len(nums) / 3
	left := make([]int64, 0, n+1)
	sum := int64(0)
	bh := newHp(maxLess)
	for i := 0; i < n*2; i++ {
		sum += int64(nums[i])
		heap.Push(bh, nums[i])
		if n < bh.Len() {
			sum -= int64(heap.Pop(bh).(int))
		}
		if bh.Len() == n {
			left = append(left, sum)
		}
	}
	ans := int64(math.MaxInt64)
	sum = 0
	sh := newHp(minLess)
	for i := len(nums) - 1; i >= n; i-- {
		sum += int64(nums[i])
		heap.Push(sh, nums[i])
		if n < sh.Len() {
			sum -= int64(heap.Pop(sh).(int))
		}
		if sh.Len() == n {
			ans = min(ans, left[i-n]-sum)
		}
	}
	return ans
}

func min(f, s int64) int64 {
	if f <= s {
		return f
	}
	return s
}

type hp struct {
	nodes []int
	l     int
	less  func(nodes []int, i, j int) bool
}

func newHp(less func(nodes []int, i, j int) bool) *hp {
	return &hp{less: less}
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
	return n
}

func (h *hp) Push(n interface{}) {
	h.nodes = append(h.nodes, n.(int))
}

func minLess(nodes []int, i, j int) bool {
	return nodes[i] < nodes[j]
}

func maxLess(nodes []int, i, j int) bool {
	return nodes[i] > nodes[j]
}
