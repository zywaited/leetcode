package one

import "container/list"

type node struct {
	num   int
	value int
	key   int
}

type LFUCache struct {
	min  int // 最小变更索引(num)
	nums map[int]*list.List
	keys map[int]*list.Element
	cap  int
	ll   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nums: make(map[int]*list.List, 0),
		keys: make(map[int]*list.Element, 0),
		cap:  capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.ll == 0 || lfu.keys[key] == nil {
		return -1
	}
	e := lfu.keys[key]
	lfu.incr(e)
	return e.Value.(*node).value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.cap == 0 {
		return
	}
	// 如果本来就存在
	if lfu.keys[key] != nil {
		e := lfu.keys[key]
		n := e.Value.(*node)
		lfu.incr(e)
		n.value = value
		return
	}
	// 如果容量满了
	if lfu.cap == lfu.ll {
		// 去除最小的一个
		e := lfu.nums[lfu.min].Front()
		n := e.Value.(*node)
		lfu.nums[lfu.min].Remove(e)
		lfu.keys[n.key] = nil
		// 新加入的索引是1
		lfu.min = 1
		// 为了重复利用内存，这里就不新创建node
		n.key = key
		n.value = value
		n.num = 1
		if lfu.nums[n.num] == nil {
			lfu.nums[n.num] = &list.List{}
		}
		lfu.nums[n.num].PushBack(n)
		lfu.keys[key] = lfu.nums[n.num].Back()
		return
	}
	n := &node{key: key, value: value, num: 1}
	if lfu.nums[n.num] == nil {
		lfu.nums[n.num] = &list.List{}
	}
	lfu.nums[n.num].PushBack(n)
	lfu.keys[key] = lfu.nums[n.num].Back()
	if lfu.min == 0 || n.num < lfu.min {
		lfu.min = n.num
	}
	lfu.ll++
}

func (lfu *LFUCache) incr(e *list.Element) {
	n := e.Value.(*node)
	// 先去除就数据
	lfu.nums[n.num].Remove(e)
	if lfu.min == n.num && lfu.nums[n.num].Len() == 0 {
		// 如果去除后没有就更新最小索引
		lfu.min++
	}
	// 加入新队列
	n.num++
	if lfu.nums[n.num] == nil {
		lfu.nums[n.num] = &list.List{}
	}
	lfu.nums[n.num].PushBack(n)
	lfu.keys[n.key] = lfu.nums[n.num].Back()
}
