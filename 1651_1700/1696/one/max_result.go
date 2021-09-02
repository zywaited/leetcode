package one

import "container/heap"

func MaxResult(nums []int, k int) int {
	cn := node{}
	h := newNodeHeap(len(nums))
	for index, num := range nums {
		for h.Len() > 0 && index-h.peek().(node).index > k {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			cn = h.peek().(node)
		}
		cn.sum += num
		if index+1 == len(nums) {
			return cn.sum
		}
		cn.index = index
		heap.Push(h, cn)
	}
	return 0
}

type node struct {
	index int
	sum   int
}

type nodeHeap []node

func newNodeHeap(n int) *nodeHeap {
	h := make(nodeHeap, 0, n)
	return &h
}

func (h *nodeHeap) Len() int {
	return len(*h)
}

func (h *nodeHeap) Less(i, j int) bool {
	return (*h)[i].sum > (*h)[j].sum
}

func (h *nodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *nodeHeap) Push(n interface{}) {
	*h = append(*h, n.(node))
}

func (h *nodeHeap) Pop() interface{} {
	n := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return n
}

func (h *nodeHeap) peek() interface{} {
	return (*h)[0]
}
