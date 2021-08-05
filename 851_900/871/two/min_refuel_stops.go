package two

import "container/heap"

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	fuels := &th{}
	ans := 0
	pm := startFuel
	tm := 0
	for index := 0; index <= len(stations); index++ {
		tm = target
		if index < len(stations) {
			tm = stations[index][0]
		}
		for pm < tm && fuels.Len() > 0 {
			pm += heap.Pop(fuels).(int)
			ans++
		}
		if pm < tm {
			return -1
		}
		if index < len(stations) {
			heap.Push(fuels, stations[index][1])
		}
	}
	return ans
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
