package one

import "container/heap"

func NetworkDelayTime(times [][]int, n int, k int) int {
	points := make([][]int, n+1)
	for index := 0; index < len(times); index++ {
		points[times[index][0]] = append(points[times[index][0]], index)
	}
	costs := make([]int, n+1)
	for index := range costs {
		costs[index] = 1 << 30
	}
	queue := &th{{index: k}}
	for queue.Len() > 0 {
		qn := heap.Pop(queue).(tn)
		if qn.cost >= costs[qn.index] {
			continue
		}
		costs[qn.index] = qn.cost
		for _, index := range points[qn.index] {
			if qn.cost+times[index][2] < costs[times[index][1]] {
				heap.Push(queue, tn{index: times[index][1], cost: qn.cost + times[index][2]})
			}
		}
	}
	ans := 0
	for index := 1; index <= n; index++ {
		if costs[index] == 1<<30 {
			return -1
		}
		if costs[index] > ans {
			ans = costs[index]
		}
	}
	return ans
}

type tn struct {
	index int
	cost  int
}

type th []tn

func (h *th) Len() int {
	return len(*h)
}

func (h *th) Less(i, j int) bool {
	return (*h)[i].cost < (*h)[j].cost
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
