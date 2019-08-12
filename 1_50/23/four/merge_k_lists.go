package four

type ListNode struct {
	Val  int
	Next *ListNode
}

// 维护一个递增的链表(类似优先队列)
func MergeKLists(lists []*ListNode) *ListNode {
	num := len(lists)
	if num == 0 {
		return nil
	}
	if num == 1 {
		return lists[0]
	}
	if num == 2 {
		return mergeList(lists[0], lists[1])
	}
	middle := num >> 1
	return mergeList(MergeKLists(lists[:middle]), MergeKLists(lists[middle:]))
}

// 合并两个链表
func mergeList(f, s *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	for f != nil && s != nil {
		if f.Val <= s.Val {
			prev.Next = f
			prev, f = prev.Next, f.Next
		} else {
			prev.Next = s
			prev, s = prev.Next, s.Next
		}
	}
	if f == nil {
		f = s
	}
	for f != nil {
		prev.Next = f
		prev, f = prev.Next, f.Next
	}
	return head.Next
}
