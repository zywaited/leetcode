package two

import (
	"container/heap"
	"sort"
)

func ScheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})
	curr := 0
	h := &th{}
	for _, course := range courses {
		if course[1] < course[0] {
			continue
		}
		if curr+course[0] <= course[1] {
			heap.Push(h, course[0])
			curr += course[0]
			continue
		}
		if course[0] < h.top().(int) {
			curr -= heap.Pop(h).(int)
			curr += course[0]
			heap.Push(h, course[0])
		}
	}
	return h.Len()
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

func (h *th) top() interface{} {
	return (*h)[0]
}
