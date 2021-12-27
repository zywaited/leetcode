package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	ans := head
	prev := head
	tail := head
	if head != nil {
		tail = head.Next
	}
	for ; tail != nil; tail = tail.Next {
		if tail.Val == prev.Val {
			// 去除当前节点
			prev.Next = tail.Next
			continue
		}
		prev = tail
	}
	return ans
}
