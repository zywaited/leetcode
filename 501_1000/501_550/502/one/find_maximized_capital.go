package one

import (
	"container/heap"
	"sort"
)

type node struct {
	profit int
	cost   int
}

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	nodes := make([]*node, len(profits))
	for i := range profits {
		nodes[i] = &node{profit: profits[i], cost: capital[i]}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].cost < nodes[j].cost
	})
	tk := 0
	ans := w
	pq := newPriorityQueue(len(nodes))
	i := 0
	for tk < k && (i <= len(nodes) || pq.Len() > 0) {
		if i < len(nodes) && nodes[i].cost <= ans {
			heap.Push(pq, nodes[i].profit)
			i++
			continue
		}
		if pq.Len() == 0 {
			break
		}
		ans += heap.Pop(pq).(int)
		tk++
	}
	return ans
}

type priorityQueue []int

func newPriorityQueue(cap int) *priorityQueue {
	pq := make(priorityQueue, 0, cap)
	return &pq
}

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i] > (*pq)[j]
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Push(v interface{}) {
	*pq = append(*pq, v.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	n := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return n
}
