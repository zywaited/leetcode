package two

import "container/list"

type LRUCache struct {
	// 索引
	index map[int]*list.Element
	// 数据
	data list.List
	// 最大值
	max int
}

type node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{index: make(map[int]*list.Element), max: capacity}
}

func (lru *LRUCache) Get(key int) int {
	if e, ok := lru.index[key]; ok {
		lru.data.MoveToFront(e)
		return e.Value.(node).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.max <= 0 {
		return
	}
	// 已经存在
	if e, ok := lru.index[key]; ok {
		n := e.Value.(node)
		n.value = value
		e.Value = n
		lru.data.MoveToFront(e)
		return
	}
	// 判断是否已经满了
	if len(lru.index) >= lru.max {
		e := lru.data.Back()
		delete(lru.index, e.Value.(node).key)
		lru.data.Remove(e)
	}

	lru.index[key] = lru.data.PushFront(node{key: key, value: value})
	return
}
