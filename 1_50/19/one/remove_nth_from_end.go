package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 为什么需要新增一节点
	// 这样是为了删除头节点是也能保证s在倒数N+1节点上
	tmp := &ListNode{Next: head}
	f := tmp
	s := tmp
	// 先处理到N,这样f到链尾时，s刚还在倒数N+1个节点
	for i := 0; i <= n; i++ {
		f = f.Next
	}
	for f != nil {
		f = f.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return tmp.Next
}
