package three

type ListNode struct {
	Val  int
	Next *ListNode
}

// 维护一个递增的链表(类似优先队列)
func MergeKLists(lists []*ListNode) *ListNode {
	type List struct {
		Val  int
		Next *List
		Node *ListNode
	}
	var l *List
	// 添加
	add := func(cn *ListNode) {
		if l == nil {
			l = &List{Val: cn.Val, Node: cn}
			return
		}
		tmp := l
		prev := (*List)(nil)
		for tmp != nil {
			if cn.Val <= tmp.Val {
				if prev == nil {
					l = &List{Val: cn.Val, Next: l, Node: cn}
					return
				}
				prev.Next = &List{Val: cn.Val, Next: prev.Next, Node: cn}
				return
			}
			prev = tmp
			tmp = tmp.Next
		}
		prev.Next = &List{Val: cn.Val, Node: cn}
	}
	// 取出
	pop := func() *List {
		node := l
		if node != nil {
			l = l.Next
		}
		return node
	}
	// 加入首位
	for _, head := range lists {
		if head != nil {
			add(head)
		}
	}
	var (
		rs   *ListNode
		tail *ListNode
	)
	// 依次取出数据
	for {
		head := pop()
		if head == nil {
			break
		}
		if rs == nil {
			rs = head.Node
		}
		if tail != nil {
			tail.Next = head.Node
		}
		tail = head.Node
		if head.Node.Next != nil {
			add(head.Node.Next)
		}
	}
	return rs
}
