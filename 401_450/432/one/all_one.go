package one

import "container/list"

type AllOne struct {
	ks *list.List               // 递增双向链表
	km map[string]*list.Element // 每个key对应的节点
	nm map[int]*lre             // 每个值对应的最左右边的节点
}

// 节点
type node struct {
	key string
	num int // 当前数量
}

type lre struct {
	left  *list.Element
	right *list.Element
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		ks: &list.List{},
		km: make(map[string]*list.Element),
		nm: make(map[int]*lre),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (an *AllOne) Inc(key string) {
	// key不存在
	if an.km[key] == nil {
		an.ks.PushFront(&node{key: key, num: 1})
		an.km[key] = an.ks.Front()
		if an.nm[1] == nil {
			an.nm[1] = &lre{right: an.ks.Front()}
		}
		an.nm[1].left = an.ks.Front() // 左边节点一定会更新
		return
	}
	ce := an.km[key]
	cn := ce.Value.(*node)
	// 数据发生变化
	if an.nm[cn.num] != nil {
		if an.nm[cn.num].left == ce && an.nm[cn.num].right == ce {
			an.nm[cn.num] = nil
		} else if an.nm[cn.num].left == ce {
			an.nm[cn.num].left = an.nm[cn.num].left.Next()
		} else if an.nm[cn.num].right == ce {
			an.nm[cn.num].right = an.nm[cn.num].right.Prev()
		}
	}
	cn.num++
	// 加入到新数据节点中
	if an.nm[cn.num] == nil {
		an.nm[cn.num] = &lre{right: ce}
	}
	an.nm[cn.num].left = ce
	// 判断是否需要移动
	if ce.Next() == nil {
		return
	}
	nn := ce.Next().Value.(*node)
	if cn.num <= nn.num {
		return
	}
	// 放到最右节点后
	an.ks.MoveAfter(ce, an.nm[nn.num].right)
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (an *AllOne) Dec(key string) {
	// 与Inc逻辑基本一致
	if an.km[key] == nil {
		return
	}
	ce := an.km[key]
	cn := ce.Value.(*node)
	if an.nm[cn.num] != nil {
		if an.nm[cn.num].left == ce && an.nm[cn.num].right == ce {
			an.nm[cn.num] = nil
		} else if an.nm[cn.num].left == ce {
			an.nm[cn.num].left = an.nm[cn.num].left.Next()
		} else if an.nm[cn.num].right == ce {
			an.nm[cn.num].right = an.nm[cn.num].right.Prev()
		}
	}
	cn.num--
	// 删除等于0
	if cn.num == 0 {
		an.ks.Remove(ce)
		return
	}
	if an.nm[cn.num] == nil {
		an.nm[cn.num] = &lre{left: ce}
	}
	an.nm[cn.num].right = ce
	if ce.Prev() == nil {
		return
	}
	pn := ce.Prev().Value.(*node)
	if cn.num >= pn.num {
		return
	}
	// 放到节点前
	an.ks.MoveBefore(ce, an.nm[pn.num].left)
}

/** Returns one of the keys with maximal value. */
func (an *AllOne) GetMaxKey() string {
	if an.ks.Len() == 0 {
		return ""
	}
	return an.ks.Back().Value.(*node).key
}

/** Returns one of the keys with Minimal value. */
func (an *AllOne) GetMinKey() string {
	if an.ks.Len() == 0 {
		return ""
	}
	return an.ks.Front().Value.(*node).key
}
