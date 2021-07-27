package one

import "container/heap"

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	if ladders >= len(heights) {
		return len(heights) - 1
	}
	diff := 0
	surplus := bricks
	h := &th{}
	for index := 1; index < len(heights); index++ {
		if heights[index] <= heights[index-1] {
			continue
		}
		diff = heights[index] - heights[index-1]
		heap.Push(h, diff)
		if h.Len() <= ladders {
			continue
		}
		diff = heap.Pop(h).(int)
		if surplus < diff {
			return index - 1
		}
		surplus -= diff
	}
	return len(heights) - 1
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
