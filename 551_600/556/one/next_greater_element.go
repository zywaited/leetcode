package one

import (
	"container/list"
	"math"
)

func NextGreaterElement(n int) int {
	// 递减链表(双向)
	origin := &list.List{}
	bit := 0
	find := false
	for n != 0 && !find {
		n, bit = n/10, n%10
		cn := origin.Front()
		mn := (*list.Element)(nil)
		for cn != nil && cn.Value.(int) > bit {
			mn = cn
			cn = cn.Next()
		}
		if mn != nil {
			find = true
			n = n*10 + mn.Value.(int)
			mn.Value = bit
			continue
		}
		origin.PushFront(bit)
	}
	rs := -1
	if find {
		rs = n
		prevMax, prevMaxBit := math.MaxInt32/10, math.MaxInt32%10
		cn := origin.Back()
		for cn != nil {
			// 是否超过32位
			if rs > prevMax || (rs == prevMax && cn.Value.(int) > prevMaxBit) {
				return -1
			}
			rs = rs*10 + cn.Value.(int)
			cn = cn.Prev()
		}
	}
	return rs
}
