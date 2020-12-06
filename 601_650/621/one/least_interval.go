package one

import "container/heap"

func LeastInterval(tasks []byte, n int) int {
	h := make(th, 26)
	for _, t := range tasks {
		h[t-'A']++
	}
	heap.Init(&h)
	ts := make([]int, n+1)
	pt := 0
	t := 0
	ct := 0
	for h.Len() > 0 {
		ts = ts[:0]
		for len(ts) <= n && h.Len() > 0 {
			t = heap.Pop(&h).(int)
			if t == 0 {
				break
			}
			ts = append(ts, t)
		}
		if len(ts) == 0 {
			break
		}
		if pt > 0 {
			ct += n + 1 - pt
		}
		pt = len(ts)
		ct += pt
		for _, t = range ts {
			if t > 1 {
				heap.Push(&h, t-1)
			}
		}
	}
	return ct
}

type th []int

func (h *th) Len() int {
	return len(*h)
}

func (h *th) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *th) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *th) Push(n interface{}) {
	*h = append(*h, n.(int))
}

func (h *th) Pop() interface{} {
	t := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return t
}
