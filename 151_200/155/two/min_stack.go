package two

import "container/list"

// 链表
type MinStack struct {
	data *list.List
	min  *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: list.New(), min: list.New()}
}

func (ms *MinStack) Push(x int) {
	ms.data.PushFront(x)
	if ms.min.Len() == 0 || x <= ms.min.Front().Value.(int) {
		ms.min.PushFront(x)
	}
}

func (ms *MinStack) Pop() {
	v := ms.data.Front()
	if v != nil {
		ms.data.Remove(v)
		mv := ms.min.Front()
		if v.Value.(int) == mv.Value.(int) {
			ms.min.Remove(mv)
		}
	}
}

func (ms *MinStack) Top() int {
	v := ms.data.Front()
	if v != nil {
		return v.Value.(int)
	}
	return 0
}

func (ms *MinStack) GetMin() int {
	mv := ms.min.Front()
	if mv != nil {
		return mv.Value.(int)
	}
	return 0
}
