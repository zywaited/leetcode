package one

import "container/list"

func RemoveKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	// 转换成链表方便处理
	nl := list.New()
	for i := range num {
		nl.PushBack(num[i])
	}

	node := nl.Front()
	prev := node
	for k > 0 && node != nil {
		for ; node.Next() != nil && node.Value.(byte) <= node.Next().Value.(byte); node = node.Next() {
		}
		prev = node.Prev()
		nl.Remove(node)
		k--
		if prev != nil {
			node = prev
			continue
		}
		// 开头去除0
		for node = nl.Front(); node != nil && node.Value.(byte) == '0'; node = nl.Front() {
			nl.Remove(node)
		}
	}
	if nl.Len() == 0 {
		return "0"
	}
	ans := make([]byte, 0, nl.Len())
	for node = nl.Front(); node != nil; node = node.Next() {
		ans = append(ans, node.Value.(byte))
	}
	return string(ans)
}
