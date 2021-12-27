package one

import "container/heap"

func TopKFrequent(words []string, k int) []string {
	wn := map[string]int{}
	for _, w := range words {
		wn[w]++
	}
	h := th{}
	for w, n := range wn {
		if len(h) < k {
			h = append(h, tn{w: w, n: n})
			if len(h) == k {
				heap.Init(&h)
			}
			continue
		}
		if h[0].n < n || (h[0].n == n && h[0].w > w) {
			h[0].w = w
			h[0].n = n
			heap.Fix(&h, 0)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(&h).(tn).w
	}
	return ans
}

type th []tn

type tn struct {
	w string
	n int
}

func (h *th) Len() int {
	return len(*h)
}

func (h *th) Less(i, j int) bool {
	if (*h)[i].n == (*h)[j].n {
		return (*h)[i].w > (*h)[j].w
	}
	return (*h)[i].n < (*h)[j].n
}

func (h *th) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *th) Push(n interface{}) {
	*h = append(*h, n.(tn))
}

func (h *th) Pop() interface{} {
	t := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return t
}
