package two

import (
	"container/heap"
)

func MaxProfit(inventory []int, orders int) int {
	mod := int(1e9 + 7)
	ans := 0
	h := th(inventory)
	heap.Init(&h)
	times := 0
	num := 0
	buy := orders
	remainder := 0
	multi := 0
	for buy > 0 {
		multi++
		num = heap.Pop(&h).(int)
		times = num
		if (&h).Len() > 0 {
			times = num - (&h).top().(int)
		}
		remainder = 0
		if buy < times*multi {
			times = buy / multi
			remainder = buy % multi
		}
		buy -= times * multi
		num -= times - 1
		ans = (ans + num*times*multi) % mod
		ans = (ans + (times-1)*times/2*multi) % mod
		if remainder > 0 {
			ans = (ans + (num-1)*remainder) % mod
			break
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

func (h *th) top() interface{} {
	return (*h)[0]
}
