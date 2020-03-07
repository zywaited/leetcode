package one

import "container/list"

// 链表或者数组都可以
type MaxQueue struct {
	max   *list.List
	queue *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		max:   &list.List{},
		queue: &list.List{},
	}
}

func (mq *MaxQueue) MaxValue() int {
	if mq.max.Len() == 0 {
		return -1
	}
	return mq.max.Front().Value.(int)
}

// 均摊时间复杂度为0(1)
func (mq *MaxQueue) PushBack(value int) {
	// 删除比当前值小的数
	for mq.max.Len() > 0 && value > mq.max.Back().Value.(int) {
		mq.max.Remove(mq.max.Back())
	}
	mq.max.PushBack(value)
	mq.queue.PushBack(value)
}

func (mq *MaxQueue) PopFront() int {
	if mq.queue.Len() == 0 {
		return -1
	}
	e := mq.queue.Front()
	mq.queue.Remove(e)
	if e.Value.(int) == mq.max.Front().Value.(int) {
		mq.max.Remove(mq.max.Front())
	}
	return e.Value.(int)
}
