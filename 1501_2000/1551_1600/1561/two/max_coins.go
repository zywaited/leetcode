package two

import "container/heap"

func MaxCoins(piles []int) int {
	ans := 0
	ps := ph(piles)
	heap.Init(&ps)
	for num := ps.Len() / 3; num > 0; num-- {
		heap.Pop(&ps)
		ans += heap.Pop(&ps).(int)
	}
	return ans
}

type ph []int

func (a ph) Len() int {
	return len(a)
}

func (a ph) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a ph) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *ph) Push(p interface{}) {
	*a = append(*a, p.(int))
}

func (a *ph) Pop() interface{} {
	p := (*a)[a.Len()-1]
	*a = (*a)[:a.Len()-1]
	return p
}
