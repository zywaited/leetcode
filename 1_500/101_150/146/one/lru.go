package one

type LRUCache struct {
	// 索引
	index map[int]int
	// 数据
	data []*node
	// 最大值，主要是为了不直接初始化data
	max int
}

type node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{index: make(map[int]int), max: capacity}
}

// 这里没有双向链表好用
// 被访问的数据放到最前并且更新索引
func (lru *LRUCache) first(index int) {
	n := lru.data[index]
	index++
	for ; index < len(lru.data); index++ {
		lru.data[index-1] = lru.data[index]
		lru.index[lru.data[index].key]--
	}
	lru.data[len(lru.data)-1] = n
	lru.index[n.key] = len(lru.data) - 1
}

func (lru *LRUCache) Get(key int) int {
	if index, ok := lru.index[key]; ok {
		value := lru.data[index].value
		lru.first(index)
		return value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.max <= 0 {
		return
	}
	// 已经存在
	if index, ok := lru.index[key]; ok {
		lru.data[index].value = value
		lru.first(index)
		return
	}
	// 判断是否已经满了
	if len(lru.index) >= lru.max {
		lru.first(0)
		// 回收利用资源
		n := lru.data[len(lru.data)-1]
		delete(lru.index, n.key)
		n.key = key
		n.value = value
	} else {
		lru.data = append(lru.data, &node{key: key, value: value})
	}
	lru.index[key] = len(lru.data) - 1
	return
}
