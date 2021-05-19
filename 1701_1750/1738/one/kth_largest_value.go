package one

import "container/heap"

func KthLargestValue(matrix [][]int, k int) int {
	rows := make([]int, len(matrix[0])+1)
	sum := make([]int, len(matrix[0]))
	h := th{}
	for i := range matrix {
		for j := range matrix[i] {
			rows[j+1] = rows[j] ^ matrix[i][j]
			sum[j] ^= rows[j+1]
			if len(h) < k {
				h = append(h, sum[j])
				if len(h) == k {
					heap.Init(&h)
				}
				continue
			}
			if h.top().(int) < sum[j] {
				heap.Pop(&h)
				heap.Push(&h, sum[j])
			}
		}
	}
	return h.top().(int)
}

type th []int

func (h *th) Len() int {
	return len(*h)
}

func (h *th) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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

func (h *th) top() interface{} {
	return (*h)[0]
}
