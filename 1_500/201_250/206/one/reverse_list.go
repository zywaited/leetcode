package one

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func ReverseList(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		next := head.Next
		node := ReverseList(next)
		next.Next = head
		head.Next = nil
		return node
	}
	return head
}
