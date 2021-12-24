package one

import "container/heap"

func EatenApples(apples []int, days []int) int {
	ans := 0
	h := &th{}
	i := 0
	for ; i < len(apples); i++ {
		for h.Len() > 0 && (h.top().(*node).day <= i || h.top().(*node).apples <= 0) {
			heap.Pop(h)
		}
		if apples[i] > 0 {
			heap.Push(h, &node{day: i + days[i], apples: apples[i]})
		}
		if h.Len() > 0 {
			ans++
			n := h.top().(*node)
			n.apples--
		}
	}
	for h.Len() > 0 {
		n := heap.Pop(h).(*node)
		if n.day <= i || n.apples == 0 {
			continue
		}
		if n.apples <= n.day-i {
			ans += n.apples
			i += n.apples
			continue
		}
		ans += n.day - i
		i = n.day
	}
	return ans
}

type node struct {
	day    int
	apples int
}

type th []*node

func (h *th) Len() int {
	return len(*h)
}

func (h *th) Less(i, j int) bool {
	return (*h)[i].day < (*h)[j].day
}

func (h *th) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *th) Push(n interface{}) {
	*h = append(*h, n.(*node))
}

func (h *th) Pop() interface{} {
	n := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return n
}

func (h *th) top() interface{} {
	return (*h)[0]
}
